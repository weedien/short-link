package dao

import (
	"gorm.io/gorm"
)

type LinkStatsTodayDao struct {
	db *gorm.DB
}

func NewLinkStatsTodayDao(db *gorm.DB) LinkStatsTodayDao {
	return LinkStatsTodayDao{db: db}
}

//// LinkTodayState 记录今日统计监控数据
//func (m *LinkStatsTodayDao) LinkTodayState(linkTodayStat po.LinkStatsToday) error {
//	rawSql := `
//INSERT INTO t_link_stats_today (full_short_url, date, today_uv, today_pv, today_uip, create_time, update_time, del_flag)
//VALUES ($1, $2, $3, $4, $5, NOW(), NOW(), 0)
//ON CONFLICT (full_short_url, date)
//DO UPDATE SET today_uv = t_link_stats_today.today_uv + EXCLUDED.today_uv,
//              today_pv = t_link_stats_today.today_pv + EXCLUDED.today_pv,
//              today_uip = t_link_stats_today.today_uip + EXCLUDED.today_uip;
//`
//	return m.db.Exec(rawSql, linkTodayStat.ShortUri, linkTodayStat.Date, linkTodayStat.TodayUv, linkTodayStat.TodayPv, linkTodayStat.TodayUip, linkTodayStat.TodayUv, linkTodayStat.TodayPv, linkTodayStat.TodayUip).Error
//}
