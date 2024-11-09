package dao

import (
	"context"
	"gorm.io/gorm"
	"shortlink/internal/base/types"
	"shortlink/internal/link_stats/adapter/po"
)

type LinkAccessLogsDao struct {
	db *gorm.DB
}

func NewLinkAccessLogsDao(db *gorm.DB) LinkAccessLogsDao {
	return LinkAccessLogsDao{db: db}
}

type IpCntDTO struct {
	Ip  string `json:"ip"`
	Cnt int    `json:"cnt"`
}

// ListTopIpByLink 根据短链接获取指定日期内高频访问IP数据
func (d LinkAccessLogsDao) ListTopIpByLink(ctx context.Context, param LinkQueryParam) ([]IpCntDTO, error) {
	rawSql := `
SELECT 
    tlal.ip, 
    COUNT(tlal.ip) AS cnt
FROM
    t_link tl INNER JOIN 
    t_link_access_logs tlal ON tl.full_short_url = tlal.full_short_url 
WHERE 
    tlal.full_short_url = ? 
    AND tl.gid = ? 
    AND tl.del_flag = false 
    AND tl.enable_status = ? 
    AND tlal.create_time BETWEEN ? and ? 
GROUP BY 
    tlal.full_short_url, tl.gid, tlal.ip 
ORDER BY 
    cnt DESC 
LIMIT 5;
`
	var result []IpCntDTO
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.FullShortUrl, param.Gid, param.Status, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}

// ListTopIpByGroup 根据分组获取指定日期内高频访问IP数据
func (d LinkAccessLogsDao) ListTopIpByGroup(ctx context.Context, param LinkGroupQueryParam) ([]IpCntDTO, error) {
	rawSql := `
SELECT 
    tlal.ip, 
    COUNT(tlal.ip) AS count 
FROM 
    t_link tl INNER JOIN 
    t_link_access_logs tlal ON tl.full_short_url = tlal.full_short_url 
WHERE 
    tl.gid = ? 
    AND tl.enable_status = ?
    AND tlal.create_time BETWEEN ? and ? 
GROUP BY 
    tl.gid, tlal.ip 
ORDER BY 
    count DESC 
LIMIT 5;
`
	var result []IpCntDTO
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.Gid, param.Status, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}

type UvTypeCntDTO struct {
	OldUserCnt int `json:"oldUserCnt"`
	NewUserCnt int `json:"newUserCnt"`
}

// FindUvTypeCntByLink 根据短链接获取指定日期内新旧访客数据
func (d LinkAccessLogsDao) FindUvTypeCntByLink(ctx context.Context, param LinkQueryParam) (UvTypeCntDTO, error) {
	rawSql := `
SELECT 
    SUM(old_user) AS oldUserCnt, 
    SUM(new_user) AS newUserCnt 
FROM ( 
    SELECT 
        CASE WHEN COUNT(DISTINCT DATE_TRUNC('day', tlal.create_time)) > 1 THEN 1 ELSE 0 END AS old_user, 
        CASE WHEN COUNT(DISTINCT DATE_TRUNC('day', tlal.create_time)) = 1 AND MAX(tlal.create_time) >= ? AND MAX(tlal.create_time) <= ? THEN 1 ELSE 0 END AS new_user 
    FROM 
        t_link tl INNER JOIN 
        t_link_access_logs tlal ON tl.full_short_url = tlal.full_short_url 
    WHERE 
        tlal.full_short_url = ? 
        AND tl.gid = ? 
        AND tl.enable_status = ? 
        AND tl.del_flag = '0' 
    GROUP BY 
        tlal.user 
) AS user_counts;
`
	var result UvTypeCntDTO
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.StartDate, param.EndDate, param.FullShortUrl, param.Gid, param.Status).Scan(&result).Error
	return result, err
}

// SelectUvTypeByUsers 获取用户信息是否新老访客
func (d LinkAccessLogsDao) SelectUvTypeByUsers(
	ctx context.Context,
	param LinkQueryParam,
	users []string,
) ([]UserType, error) {
	rawSql := `
SELECT 
    tlal.user, 
    CASE 
        WHEN MIN(tlal.create_time) BETWEEN ? AND ? THEN '新访客' 
        ELSE '老访客' 
    END AS uvType 
FROM 
    t_link tl INNER JOIN 
    t_link_access_logs tlal ON tl.full_short_url = tlal.full_short_url 
WHERE 
    tlal.full_short_url = ? 
    AND tl.gid = ? 
    AND tl.del_flag = '0' 
    AND tl.enable_status = ? 
    AND tlal.user IN (?) 
GROUP BY 
    tlal.user;
`
	var result []UserType
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.StartDate, param.EndDate, param.FullShortUrl, param.Gid, param.Status, users).Scan(&result).Error
	return result, err
}

type UserType struct {
	User   string
	UvType string
}

// SelectGroupUvTypeByUsers 获取分组用户信息是否新老访客
func (d LinkAccessLogsDao) SelectGroupUvTypeByUsers(
	ctx context.Context,
	param LinkGroupQueryParam,
	users []string,
) ([]UserType, error) {
	rawSql := `
SELECT 
    tlal.user, 
    CASE 
        WHEN MIN(tlal.create_time) BETWEEN ? AND ? THEN '新访客' 
        ELSE '老访客' 
    END AS uvType 
FROM 
    t_link tl INNER JOIN 
    t_link_access_logs tlal ON tl.full_short_url = tlal.full_short_url 
WHERE 
    tl.gid = ? 
    AND tl.enable_status = ?
    AND tlal.user IN (?) 
GROUP BY 
    tlal.user;
`
	var result []UserType
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.StartDate, param.EndDate, param.Gid, param.Status, users).Scan(&result).Error
	return result, err
}

type PvUvUidStatDTO struct {
	Pv  int `json:"pv"`
	Uv  int `json:"uv"`
	Uip int `json:"uip"`
}

// FindPvUvUidStatByLink 根据短链接获取指定日期内PV、UV、UIP数据
func (d LinkAccessLogsDao) FindPvUvUidStatByLink(ctx context.Context, param LinkQueryParam) (PvUvUidStatDTO, error) {
	rawSql := `
SELECT 
    COUNT(tlal.user) AS pv, 
    COUNT(DISTINCT tlal.user) AS uv, 
    COUNT(DISTINCT tlal.ip) AS uip 
FROM 
    t_link tl INNER JOIN 
    t_link_access_logs tlal ON tl.full_short_url = tlal.full_short_url 
WHERE 
    tlal.full_short_url = ? 
    AND tl.gid = ? 
    AND tl.del_flag = '0' 
    AND tl.enable_status = ? 
    AND tlal.create_time BETWEEN ? and ? 
GROUP BY 
    tlal.full_short_url, tl.gid;
`
	var result PvUvUidStatDTO
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.FullShortUrl, param.Gid, param.Status, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}

// Page 根据短链接分页获取访问日志
func (d LinkAccessLogsDao) Page(
	ctx context.Context,
	param LinkQueryParam,
	page, size int,
) (*types.PageResp[po.LinkAccessLog], error) {
	rawSql := `
SELECT
	tlal.*
FROM
	t_link tl INNER JOIN
	t_link_access_logs tlal ON tl.full_short_url = tlal.full_short_url
WHERE
	tlal.full_short_url = ?
	AND tl.gid = ?
	AND tl.del_flag = '0'
	AND tl.enable_status = ?
	AND tlal.create_time BETWEEN ? and ?
ORDER BY
	tlal.create_time DESC
LIMIT ? OFFSET ?;
`
	var total int64
	var result []po.LinkAccessLog
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.FullShortUrl, param.Gid, param.Status, param.StartDate, param.EndDate, page, size).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	err = d.db.Model(&po.LinkAccessLog{}).Count(&total).Error
	if err != nil {
		return nil, err
	}
	return &types.PageResp[po.LinkAccessLog]{
		Total:   total,
		Records: result,
	}, nil
}

// FindPvUvUidStatByGroup 根据分组获取指定日期内PV、UV、UIP数据
func (d LinkAccessLogsDao) FindPvUvUidStatByGroup(ctx context.Context, param LinkGroupQueryParam) (PvUvUidStatDTO, error) {
	rawSql := `
SELECT 
    COUNT(tlal.user) AS pv, 
    COUNT(DISTINCT tlal.user) AS uv, 
    COUNT(DISTINCT tlal.ip) AS uip 
FROM 
    t_link tl INNER JOIN 
    t_link_access_logs tlal ON tl.full_short_url = tlal.full_short_url 
WHERE 
    tl.gid = ? 
    AND tl.del_flag = '0' 
    AND tl.enable_status = ?
    AND tlal.create_time BETWEEN ? and ? 
GROUP BY 
    tl.gid;
`
	var result PvUvUidStatDTO
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.Gid, param.Status, param.StartDate, param.EndDate).Scan(&result).Error
	return result, err
}

// PageGroup 分页获取分组访问日志
func (d LinkAccessLogsDao) PageGroup(
	ctx context.Context,
	param LinkGroupQueryParam,
	page, size int,
) (*types.PageResp[po.LinkAccessLog], error) {
	rawSql := `
SELECT 
    tlal.* 
FROM 
    t_link tl 
    INNER JOIN t_link_access_logs tlal ON tl.full_short_url = tlal.full_short_url 
WHERE 
    tl.gid = ? 
    AND tl.enable_status = ? 
    AND tlal.create_time BETWEEN ? and ? 
ORDER BY 
    tlal.create_time DESC
LIMIT ? OFFSET ?;
`
	var total int64
	var result []po.LinkAccessLog
	err := d.db.WithContext(ctx).
		Raw(rawSql, param.Gid, param.Status, param.StartDate, param.EndDate, page, size).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	err = d.db.Model(&po.LinkAccessLog{}).Count(&total).Error
	if err != nil {
		return nil, err
	}
	return &types.PageResp[po.LinkAccessLog]{
		Total:   total,
		Records: result,
	}, nil
}
