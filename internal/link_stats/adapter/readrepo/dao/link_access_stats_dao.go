package dao

import (
	"context"
	"gorm.io/gorm"
	"shortlink/internal/link_stats/adapter/po"
)

type LinkAccessStatDao struct {
	db *gorm.DB
}

func NewLinkAccessStatDao(db *gorm.DB) LinkAccessStatDao {
	return LinkAccessStatDao{db: db}
}

//// LinkStats 记录基础访问监控数据
//func (m *LinkAccessStatDao) LinkStats(LinkAccessStat po.LinkAccessStat) error {
//	rawSql := `
//INSERT INTO t_link_access_stats (full_short_url, date, pv, uv, uip, hour, weekday, create_time, update_time, del_flag)
//VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW(), 0)
//ON DUPLICATE KEY UPDATE pv = pv + ?, uv = uv + ?, uip = uip + ?;
//`
//	return m.db.Exec(rawSql, LinkAccessStat.ShortUri, LinkAccessStat.Date, LinkAccessStat.Pv, LinkAccessStat.Uv, LinkAccessStat.Uip, LinkAccessStat.Hour, LinkAccessStat.Weekday, LinkAccessStat.Pv, LinkAccessStat.Uv, LinkAccessStat.Uip).Error
//}

// ListStatByLink 根据短链接获取指定日期内基础监控数据
func (d *LinkAccessStatDao) ListStatByLink(ctx context.Context, param LinkQueryParam) ([]po.LinkAccessStat, error) {
	rawSql := `
SELECT
    tlas.date,
    SUM(tlas.pv) AS pv,
    SUM(tlas.uv) AS uv,
    SUM(tlas.uip) AS uip
FROM
    t_link tl INNER JOIN
    t_link_access_stats tlas ON tl.full_short_url = tlas.full_short_url
WHERE
    tlas.full_short_url = ?
    AND tl.gid = ?
    AND tl.del_flag = '0'
    AND tl.enable_status = ?
    AND tlas.date BETWEEN ? and ?
GROUP BY
    tlas.full_short_url, tl.gid, tlas.date;
`
	var result []po.LinkAccessStat
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.FullShortUrl, param.Gid, param.Status, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}

// ListStatByGroup 根据分组获取指定日期内基础监控数据
func (d *LinkAccessStatDao) ListStatByGroup(ctx context.Context, param LinkGroupQueryParam) ([]po.LinkAccessStat, error) {
	rawSql := `
SELECT
    tlas.date,
    SUM(tlas.pv) AS pv,
    SUM(tlas.uv) AS uv,
    SUM(tlas.uip) AS uip
FROM
    t_link tl INNER JOIN
    t_link_access_stats tlas ON tl.full_short_url = tlas.full_short_url
WHERE
    tl.gid = ?
    AND tl.del_flag = '0'
    AND tl.enable_status = '0'
    AND tlas.date BETWEEN ? and ?
GROUP BY
    tl.gid, tlas.date;
`
	var result []po.LinkAccessStat
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.Gid, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}

// ListHourStatByLink 根据短链接获取指定日期内小时基础监控数据
func (d *LinkAccessStatDao) ListHourStatByLink(ctx context.Context, param LinkQueryParam) ([]po.LinkAccessStat, error) {
	rawSql := `
SELECT
    tlas.hour,
    SUM(tlas.pv) AS pv
FROM
    t_link tl INNER JOIN
    t_link_access_stats tlas ON tl.full_short_url = tlas.full_short_url
WHERE
    tlas.full_short_url = ?
    AND tl.gid = ?
    AND tl.del_flag = '0'
    AND tl.enable_status = ?
    AND tlas.date BETWEEN ? and ?
GROUP BY
    tlas.full_short_url, tl.gid, tlas.hour;
`
	var result []po.LinkAccessStat
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.FullShortUrl, param.Gid, param.Status, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}

// ListHourStatByGroup 根据分组获取指定日期内小时基础监控数据
func (d *LinkAccessStatDao) ListHourStatByGroup(ctx context.Context, param LinkGroupQueryParam) ([]po.LinkAccessStat, error) {
	rawSql := `
SELECT
    tlas.hour,
    SUM(tlas.pv) AS pv
FROM
    t_link tl INNER JOIN
    t_link_access_stats tlas ON tl.full_short_url = tlas.full_short_url
WHERE
    tl.gid = ?
    AND tl.del_flag = '0'
    AND tl.enable_status = '0'
    AND tlas.date BETWEEN ? and ?
GROUP BY
    tl.gid, tlas.hour;
`
	var result []po.LinkAccessStat
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.Gid, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}

// ListWeekdayStatByLink 根据短链接获取指定日期内小时基础监控数据
func (d *LinkAccessStatDao) ListWeekdayStatByLink(ctx context.Context, param LinkQueryParam) ([]po.LinkAccessStat, error) {
	rawSql := `
SELECT
    tlas.weekday,
    SUM(tlas.pv) AS pv
FROM
    t_link tl INNER JOIN
    t_link_access_stats tlas ON tl.full_short_url = tlas.full_short_url
WHERE
    tlas.full_short_url = ?
    AND tl.gid = ?
    AND tl.del_flag = '0'
    AND tl.enable_status = ?
    AND tlas.date BETWEEN ? and ?
GROUP BY
    tlas.full_short_url, tl.gid, tlas.weekday;
`
	var result []po.LinkAccessStat
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.FullShortUrl, param.Gid, param.Status, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}

// ListWeekdayStatByGroup 根据分组获取指定日期内小时基础监控数据
func (d *LinkAccessStatDao) ListWeekdayStatByGroup(ctx context.Context, param LinkGroupQueryParam) ([]po.LinkAccessStat, error) {
	rawSql := `
SELECT
    tlas.weekday,
    SUM(tlas.pv) AS pv
FROM
    t_link tl INNER JOIN
    t_link_access_stats tlas ON tl.full_short_url = tlas.full_short_url
WHERE
    tl.gid = ?
    AND tl.del_flag = '0'
    AND tl.enable_status = '0'
    AND tlas.date BETWEEN ? and ?
GROUP BY
    tl.gid, tlas.weekday;
`
	var result []po.LinkAccessStat
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.Gid, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}
