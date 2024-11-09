package read

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"shortlink/internal/base/cache"
	"shortlink/internal/base/errno"
	"shortlink/internal/base/types"
	"shortlink/internal/link/adapter/po"
	"shortlink/internal/link/app/query"
	"shortlink/internal/link/common/constant"
	"shortlink/internal/link/domain/link"
	"strconv"
	"time"
)

type LinkQuery struct {
	linkFactory      *link.Factory
	db               *gorm.DB
	distributedCache cache.DistributedCache
}

func NewLinkQuery(db *gorm.DB, factory *link.Factory, distributedCache cache.DistributedCache) *LinkQuery {
	return &LinkQuery{
		linkFactory:      factory,
		db:               db,
		distributedCache: distributedCache,
	}
}

func (q LinkQuery) GetLink(ctx context.Context, shortUri string) (*link.Link, error) {

	var linkPo po.Link
	if err := q.db.WithContext(ctx).Where("short_uri = ?", shortUri).First(&linkPo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errno.LinkNotExists
		}
		return nil, err
	}

	// 将持久化对象转换为领域模型
	lk := &link.Link{}

	start, end := new(time.Time), new(time.Time)
	if linkPo.StartDate.Valid {
		*start = linkPo.StartDate.Time
	}
	if linkPo.EndDate.Valid {
		*end = linkPo.EndDate.Time
	}

	if validDate, err := link.NewValidDate(link.ValidType(linkPo.ValidType), start, end); err != nil {
		return nil, err
	} else {
		if lk, err = q.linkFactory.NewLinkFromDB(linkPo.ID, linkPo.Gid, linkPo.ShortUri, linkPo.OriginalUrl, link.Status(linkPo.Status),
			link.CreateType(linkPo.CreateType), linkPo.Favicon, linkPo.Desc, validDate); err != nil {
			return nil, err
		}
	}

	return lk, nil
}

func (q LinkQuery) PageLink(ctx context.Context, param query.PageLink) (*types.PageResp[query.Link], error) {
	var records []query.Link
	var total int64

	// 1. 创建基本查询构建器
	baseQuery := q.db.WithContext(ctx).
		Table("t_link l").
		Where("l.recycle_time IS NULL and l.delete_time IS NULL and l.tenant_id = ?", ctx.Value("username"))

	// 2. 根据 Gid 过滤条件
	if param.Gid != nil {
		baseQuery = baseQuery.Where("l.gid = ?", *param.Gid)
	}

	// 3. 查询记录
	queryB := baseQuery.
		Select("l.*, COALESCE(st.today_pv, 0) AS todayPv, COALESCE(st.today_uv, 0) AS todayUv, COALESCE(st.today_uip, 0) AS todayUip").
		Joins("LEFT JOIN t_link_stats_today st ON l.short_uri = st.short_uri AND st.date = current_date").
		Joins("LEFT JOIN t_link_stats ls ON l.short_uri = ls.short_uri").
		Order(getOrderClause(param.OrderTag)).
		Limit(param.Limit()).
		Offset(param.Offset())

	if err := queryB.Find(&records).Error; err != nil {
		return nil, fmt.Errorf("failed to scan records: %w", err)
	}

	// 4. 查询总数
	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count total: %w", err)
	}

	// 5. 返回结果
	return &types.PageResp[query.Link]{
		Current: param.Current,
		Size:    param.Size,
		Total:   total,
		Records: records,
	}, nil
}

// getOrderClause 根据 OrderTag 生成排序条件
func getOrderClause(orderTag *string) string {
	if orderTag == nil {
		return "l.create_time DESC"
	}
	order := *orderTag
	switch order {
	case "todayPv":
		return "t.todayPv DESC"
	case "todayUv":
		return "t.todayUv DESC"
	case "todayUip":
		return "t.todayUip DESC"
	case "totalPv":
		return "ls.total_pv DESC"
	case "totalUv":
		return "ls.total_uv DESC"
	case "totalUip":
		return "ls.total_uip DESC"
	default:
		return "l.create_time DESC"
	}
}

func (q LinkQuery) PageRecycleBin(ctx context.Context, param query.PageRecycleBin) (*types.PageResp[query.Link], error) {
	var records []query.Link
	var total int64

	// 1. 创建基本查询构建器
	baseQuery := q.db.WithContext(ctx).
		Table("t_link l").
		Where("l.recycle_time IS NOT NULL and l.tenant_id = ? and l.delete_time IS NULL", ctx.Value("username"))

	if len(param.Gids) > 0 && param.Gids[0] != "" {
		baseQuery = baseQuery.Where("l.gid IN ?", param.Gids)
	}

	// 2. 查询记录
	queryB := baseQuery.
		Select("l.*, COALESCE(t.today_pv, 0) AS todayPv, COALESCE(t.today_uv, 0) AS todayUv, COALESCE(t.today_uip, 0) AS todayUip").
		Joins("LEFT JOIN t_link_stats_today t ON l.short_uri = t.short_uri AND t.date = current_date").
		Order("l.update_time DESC").
		Limit(param.Limit()).
		Offset(param.Offset())

	if err := queryB.
		Find(&records).Error; err != nil {
		return nil, fmt.Errorf("failed to scan records: %w", err)
	}

	// 需要对record进行转换 添加上一些字段

	// 3. 查询总数
	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count total: %w", err)
	}

	// 4. 返回结果
	return &types.PageResp[query.Link]{
		Current: param.Current,
		Size:    param.Size,
		Total:   total,
		Records: records,
	}, nil
}

func (q LinkQuery) ListGroupLinkCount(ctx context.Context, gids []string) ([]query.GroupLinkCount, error) {
	username := ctx.Value("username").(string)

	var res []query.GroupLinkCount

	// 先从缓存中获取
	cnt, err := q.distributedCache.HGetAll(ctx, constant.LinkGroupCountKey+username)
	if err != nil || len(cnt) == 0 {
		// 如果缓存中没有数据或查询缓存时出错，从数据库查询
		//if err != nil {
		//fmt.Printf("Error accessing cache: %v\n", err)
		//}

		if err = q.db.WithContext(ctx).
			Model(&po.Link{}).
			Select("gid, COUNT(*) AS link_count").Where("gid IN ?", gids).
			Group("gid").
			Find(&res).Error; err != nil {
			return nil, err
		}
		return res, nil
	}

	// 如果缓存中有数据，直接返回
	for _, gid := range gids {
		if v, ok := cnt[gid]; ok {
			count := 0
			if count, err = strconv.Atoi(v); err != nil {
				return nil, err
			}
			res = append(res, query.GroupLinkCount{Gid: gid, LinkCount: count})
		}
	}

	if len(res) == 0 {
		return nil, errno.LinkGroupEmpty
	}

	return res, nil
}
