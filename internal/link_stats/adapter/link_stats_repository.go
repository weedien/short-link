package adapter

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log/slog"
	"shortlink/internal/base/lock"
	"shortlink/internal/base/toolkit"
	"shortlink/internal/link/constant"
	"shortlink/internal/link/domain/event"
	"shortlink/internal/link_stats/adapter/po"
)

type LinkStatsRepository struct {
	db     *gorm.DB
	rdb    *redis.Client
	locker lock.DistributedLock
}

func NewLinkStatsRepository(db *gorm.DB, rdb *redis.Client) LinkStatsRepository {
	return LinkStatsRepository{db: db, rdb: rdb}
}

func (r LinkStatsRepository) SaveLinkStats(ctx context.Context, statsInfo event.UserVisitInfo) error {
	lockKey := constant.LockGidUpdateKey + statsInfo.ShortUri
	if _, err := r.locker.Acquire(ctx, lockKey, -1); err != nil {
		return err
	}
	defer func() {
		if err := r.locker.Release(ctx, lockKey); err != nil {
			slog.Error("release lock failed", "lockKey", lockKey, "err", err)
		}
	}()

	shortUri := statsInfo.ShortUri
	currentDate := statsInfo.CurrentDate
	hour := currentDate.Hour() + 1
	weekDay := int(currentDate.Weekday())

	// 访问统计
	// 确定两个值的信息，uvFirstFlag 和 uipFirstFlag
	uv, uip := 0, 0
	uvAdded, err := r.rdb.SAdd(ctx, constant.LinkStatsUvKey+shortUri, statsInfo.UV).Result()
	if err != nil {
		return err
	}
	if uvAdded > 0 {
		uv = 1
	}
	uipAdded, err := r.rdb.SAdd(ctx, constant.LinkStatsUipKey+shortUri, statsInfo.RemoteAddr).Result()
	if err != nil {
		return err
	}
	if uipAdded > 0 {
		uip = 1
	}
	linkAccessStatPo := po.LinkAccessStat{
		Pv:       1,
		Uv:       uv,
		Uip:      uip,
		Hour:     hour,
		Week:     weekDay,
		ShortUri: shortUri,
		Date:     currentDate,
	}
	if err := r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "short_uri"}, {Name: "date"}, {Name: "hour"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"pv":  gorm.Expr("pv + ?", 1),
			"uv":  gorm.Expr("uv + ?", uv),
			"uip": gorm.Expr("uip + ?", uip),
		}),
	}).Create(&linkAccessStatPo).Error; err != nil {
		return err
	}
	// 今日统计
	linkStatToday := po.LinkStatsToday{
		TodayPv:  1,
		TodayUv:  uv,
		TodayUip: uip,
		Date:     currentDate,
		ShortUri: shortUri,
	}
	if err := r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "short_uri"}, {Name: "date"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"today_pv":  gorm.Expr("today_pv + ?", 1),
			"today_uv":  gorm.Expr("today_uv + ?", uv),
			"today_uip": gorm.Expr("today_uip + ?", uip),
		}),
	}).Create(&linkStatToday).Error; err != nil {
		return err
	}
	// 地区信息
	location := toolkit.GetLocationByIP(statsInfo.RemoteAddr)
	linkLocaleStatPo := po.LinkLocaleStat{
		ShortUri: shortUri,
		Date:     currentDate,
		Cnt:      1,
		Province: location.RegionName,
		City:     location.City,
		Country:  location.Country,
	}
	if err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "short_uri"}, {Name: "date"}, {Name: "province"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"cnt": gorm.Expr("cnt + ?", 1)}),
	}).Create(&linkLocaleStatPo).Error; err != nil {
		return err
	}
	// 操作系统信息
	linkOsStatPo := po.LinkOsStat{
		Os:       statsInfo.OS,
		Cnt:      1,
		ShortUri: shortUri,
		Date:     currentDate,
	}
	if err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "os"}, {Name: "short_uri"}, {Name: "date"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"cnt": gorm.Expr("cnt + ?", 1)}),
	}).Create(&linkOsStatPo).Error; err != nil {
		return err
	}
	// 浏览器信息
	linkBrowserStatPo := po.LinkBrowserStat{
		Browser:  statsInfo.Browser,
		Cnt:      1,
		ShortUri: shortUri,
		Date:     currentDate,
	}
	if err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "browser"}, {Name: "short_uri"}, {Name: "date"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"cnt": gorm.Expr("cnt + ?", 1)}),
	}).Create(&linkBrowserStatPo).Error; err != nil {
		return err
	}
	// 设备信息
	linkDeviceStatPo := po.LinkDeviceStat{
		Device:   statsInfo.Device,
		Cnt:      1,
		ShortUri: shortUri,
		Date:     currentDate,
	}
	if err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "device"}, {Name: "short_uri"}, {Name: "date"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"cnt": gorm.Expr("cnt + ?", 1)}),
	}).Create(&linkDeviceStatPo).Error; err != nil {
		return err
	}
	// 网络信息
	linkNetworkStatPo := po.LinkNetworkStat{
		Network:  statsInfo.Network,
		Cnt:      1,
		ShortUri: shortUri,
		Date:     currentDate,
	}
	if err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "network"}, {Name: "short_uri"}, {Name: "date"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"cnt": gorm.Expr("cnt + ?", 1)}),
	}).Create(&linkNetworkStatPo).Error; err != nil {
		return err
	}
	// 访问日志
	linkAccessLogPo := po.LinkAccessLog{
		ShortUri: shortUri,
		User:     statsInfo.UV,
		IP:       statsInfo.RemoteAddr,
		Browser:  statsInfo.Browser,
		Os:       statsInfo.OS,
		Network:  statsInfo.Network,
		Device:   statsInfo.Device,
		Locale:   location.Country + "-" + location.RegionName + "-" + location.City,
	}
	if err := r.db.Create(&linkAccessLogPo).Error; err != nil {
		return err
	}
	// 更新shortLink表中的状态pv, uv, uip
	linkGotoPo := po.LinkGoto{ShortUri: shortUri}
	if err := r.db.First(&linkGotoPo).Error; err != nil {
		return err
	}
	r.db.Model(&po.Link{}).
		Where("gid = ? and short_uri = ?", linkGotoPo.Gid, shortUri).
		Updates(map[string]interface{}{
			"total_pv":  gorm.Expr("total_pv + ?", 1),
			"total_uv":  gorm.Expr("total_uv + ?", uv),
			"total_uip": gorm.Expr("total_uip + ?", uip),
		})

	return nil
}
