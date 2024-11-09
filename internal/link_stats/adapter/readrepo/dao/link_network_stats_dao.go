package dao

import (
	"context"
	"gorm.io/gorm"
	"shortlink/internal/link_stats/adapter/po"
)

type LinkNetworkStatDao struct {
	db *gorm.DB
}

func NewLinkNetworkStatDao(db *gorm.DB) LinkNetworkStatDao {
	return LinkNetworkStatDao{db: db}
}

//// LinkNetworkState 记录访问网络监控数据
//func (m *LinkNetworkStatDao) LinkNetworkState(LinkNetworkStat po.LinkNetworkStat) error {
//	rawSql := `
//INSERT INTO t_link_network_stats (full_short_url, date, cnt, network, create_time, update_time, del_flag)
//VALUES (?, ?, ?, ?, NOW(), NOW(), 0)
//ON DUPLICATE KEY UPDATE cnt = cnt + ?;
//`
//	return m.db.Exec(rawSql, LinkNetworkStat.ShortUri, LinkNetworkStat.Date, LinkNetworkStat.Cnt, LinkNetworkStat.Network, LinkNetworkStat.Cnt).Error
//}

// ListNetworkStatByLink 根据短链接获取指定日期内访问网络监控数据
func (d *LinkNetworkStatDao) ListNetworkStatByLink(ctx context.Context, param LinkQueryParam) ([]po.LinkNetworkStat, error) {
	rawSql := `
SELECT
    tlns.network,
    SUM(tlns.cnt) AS cnt
FROM
    t_link tl INNER JOIN
    t_link_network_stats tlns ON tl.full_short_url = tlns.full_short_url
WHERE
    tlns.full_short_url = ?
    AND tl.gid = ?
    AND tl.del_flag = '0'
    AND tl.enable_status = ?
    AND tlns.date BETWEEN ? and ?
GROUP BY
    tlns.full_short_url, tl.gid, tlns.network;
`
	var result []po.LinkNetworkStat
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.FullShortUrl, param.Gid, param.Status, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}

// ListNetworkStatByGroup 根据分组获取指定日期内访问网络监控数据
func (d *LinkNetworkStatDao) ListNetworkStatByGroup(ctx context.Context, param LinkGroupQueryParam) ([]po.LinkNetworkStat, error) {
	rawSql := `
SELECT
    tlns.network,
    SUM(tlns.cnt) AS cnt
FROM
    t_link tl INNER JOIN
    t_link_network_stats tlns ON tl.full_short_url = tlns.full_short_url
WHERE
    tl.gid = ?
    AND tl.del_flag = false
    AND tl.enable_status = '0'
    AND tlns.date BETWEEN ? and ?
GROUP BY
    tl.gid, tlns.network;
`
	var result []po.LinkNetworkStat
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.Gid, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}
