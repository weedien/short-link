package dao

import (
	"context"
	"gorm.io/gorm"
	"shortlink/internal/link_stats/adapter/po"
)

type LinkLocaleStatDao struct {
	db *gorm.DB
}

func NewLinkLocaleStatDao(db *gorm.DB) LinkLocaleStatDao {
	return LinkLocaleStatDao{db: db}
}

//// LinkLocaleState 记录地区访问监控数据
//func (m *LinkLocaleStatDao) LinkLocaleState(LinkLocaleStat po.LinkLocaleStat) error {
//	rawSql := `
//INSERT INTO t_link_locale_stats (full_short_url, date, cnt, country, province, city, adcode, create_time, update_time, del_flag)
//VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW(), 0)
//ON DUPLICATE KEY UPDATE cnt = cnt + ?;
//`
//	return m.db.Exec(rawSql, LinkLocaleStat.ShortUri, LinkLocaleStat.Date, LinkLocaleStat.Cnt, LinkLocaleStat.Country, LinkLocaleStat.Province, LinkLocaleStat.City, LinkLocaleStat.Adcode, LinkLocaleStat.Cnt).Error
//}

// ListLocaleByLink 根据短链接获取指定日期内地区监控数据
func (d *LinkLocaleStatDao) ListLocaleByLink(ctx context.Context, param LinkQueryParam) ([]po.LinkLocaleStat, error) {
	rawSql := `
SELECT
    tlls.province,
    SUM(tlls.cnt) AS cnt
FROM
    t_link tl INNER JOIN
    t_link_locale_stats tlls ON tl.full_short_url = tlls.full_short_url
WHERE
    tlls.full_short_url = ?
    AND tl.gid = ?
    AND tl.del_flag = '0'
    AND tl.enable_status = ?
    AND tlls.date BETWEEN ? and ?
GROUP BY
    tlls.full_short_url, tl.gid, tlls.province;
`
	var result []po.LinkLocaleStat
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.FullShortUrl, param.Gid, param.Status, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}

// ListLocaleByGroup 根据分组获取指定日期内地区监控数据
func (d *LinkLocaleStatDao) ListLocaleByGroup(ctx context.Context, param LinkGroupQueryParam) ([]po.LinkLocaleStat, error) {
	rawSql := `
SELECT
    tlls.province,
    SUM(tlls.cnt) AS cnt
FROM
    t_link tl INNER JOIN
    t_link_locale_stats tlls ON tl.full_short_url = tlls.full_short_url
WHERE
    tl.gid = ?
    AND tl.del_flag = '0'
    AND tl.enable_status = '0'
    AND tlls.date BETWEEN ? and ?
GROUP BY
    tl.gid, tlls.province;
`
	var result []po.LinkLocaleStat
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.Gid, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}
