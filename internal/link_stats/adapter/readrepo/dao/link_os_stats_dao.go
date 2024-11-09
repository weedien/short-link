package dao

import (
	"context"
	"gorm.io/gorm"
	"shortlink/internal/link_stats/adapter/po"
)

type LinkOsStatDao struct {
	db *gorm.DB
}

func NewLinkOsStatDao(db *gorm.DB) LinkOsStatDao {
	return LinkOsStatDao{db: db}
}

//// LinkOsState 记录操作系统访问监控数据
//func (m *LinkOsStatDao) LinkOsState(LinkOsStat po.LinkOsStat) error {
//	rawSql := `
//INSERT INTO t_link_os_stats (full_short_url, date, cnt, os, create_time, update_time, del_flag)
//VALUES (?, ?, ?, ?, NOW(), NOW(), 0)
//ON DUPLICATE KEY UPDATE cnt = cnt + ?;
//`
//	return m.db.Exec(rawSql, LinkOsStat.ShortUri, LinkOsStat.Date, LinkOsStat.Cnt, LinkOsStat.Os, LinkOsStat.Cnt).Error
//}

// ListOsStatByLink 根据短链接获取指定日期内操作系统监控数据
func (d *LinkOsStatDao) ListOsStatByLink(ctx context.Context, param LinkQueryParam) ([]po.LinkOsStat, error) {
	rawSql := `
SELECT
    tlos.os,
    SUM(tlos.cnt) AS count
FROM
    t_link tl INNER JOIN
    t_link_os_stats tlos ON tl.full_short_url = tlos.full_short_url
WHERE
    tlos.full_short_url = ?
    AND tl.gid = ?
    AND tl.del_flag = '0'
    AND tl.enable_status = ?
    AND tlos.date BETWEEN ? and ?
GROUP BY
    tlos.full_short_url, tl.gid, tlos.os;
`
	var result []po.LinkOsStat
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.FullShortUrl, param.Gid, param.Status, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}

// ListOsStatByGroup 根据分组获取指定日期内操作系统监控数据
func (d *LinkOsStatDao) ListOsStatByGroup(ctx context.Context, param LinkGroupQueryParam) ([]po.LinkOsStat, error) {
	rawSql := `
SELECT
    tlos.os,
    SUM(tlos.cnt) AS count
FROM
    t_link tl INNER JOIN
    t_link_os_stats tlos ON tl.full_short_url = tlos.full_short_url
WHERE
    tl.gid = ?
    AND tl.del_flag = '0'
    AND tl.enable_status = ?
    AND tlos.date BETWEEN ? and ?
GROUP BY
    tl.gid, tlos.os;
`
	var result []po.LinkOsStat
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.Gid, param.Status, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}
