package dao

import (
	"context"
	"gorm.io/gorm"
	"shortlink/internal/link_stats/adapter/po"
)

type LinkBrowserStatDao struct {
	db *gorm.DB
}

func NewLinkBrowserStatDao(db *gorm.DB) LinkBrowserStatDao {
	return LinkBrowserStatDao{db: db}
}

//// LinkBrowserState 记录浏览器访问监控数据
//func (m *LinkBrowserStatDao) LinkBrowserState(LinkBrowserStat po.LinkBrowserStat) error {
//	rawSql := `
//INSERT INTO t_link_browser_stats (full_short_url, date, cnt, browser, create_time, update_time, del_flag)
//VALUES (?, ?, ?, ?, NOW(), NOW(), 0)
//ON DUPLICATE KEY UPDATE cnt = cnt + ?;
//`
//	return m.db.Exec(rawSql, LinkBrowserStat.ShortUri, LinkBrowserStat.Date, LinkBrowserStat.Cnt, LinkBrowserStat.Browser, LinkBrowserStat.Cnt).Error
//}

// ListBrowserStatByLink 根据短链接获取指定日期内浏览器监控数据
func (d *LinkBrowserStatDao) ListBrowserStatByLink(ctx context.Context, param LinkQueryParam) ([]po.LinkBrowserStat, error) {
	rawSql := `
SELECT
    tlbs.browser,
    SUM(tlbs.cnt) AS cnt
FROM
    t_link tl INNER JOIN
    t_link_browser_stats tlbs ON tl.full_short_url = tlbs.full_short_url
WHERE
    tlbs.full_short_url = ?
    AND tl.gid = ?
    AND tl.del_flag = '0'
    AND tl.enable_status = ?
    AND tlbs.date BETWEEN ? and ?
GROUP BY
    tlbs.full_short_url, tl.gid, tlbs.browser;
`
	var result []po.LinkBrowserStat
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.FullShortUrl, param.Gid, param.Status, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}

// ListBrowserStatByGroup 根据分组获取指定日期内浏览器监控数据
func (d *LinkBrowserStatDao) ListBrowserStatByGroup(ctx context.Context, param LinkGroupQueryParam) ([]po.LinkBrowserStat, error) {
	rawSql := `
SELECT
    tlbs.browser,
    SUM(tlbs.cnt) AS count
FROM
    t_link tl INNER JOIN
    t_link_browser_stats tlbs ON tl.full_short_url = tlbs.full_short_url
WHERE
    tl.gid = ?
    AND tl.del_flag = '0'
    AND tl.enable_status = '0'
    AND tlbs.date BETWEEN ? and ?
GROUP BY
    tl.gid, tlbs.browser;
`
	var result []po.LinkBrowserStat
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.Gid, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}
