package link_stats

import "errors"

// Stats 统计信息
//
// link_stats 是对 link 领域的一个子领域，它主要负责统计相关的信息
type Stats struct {
	clickNum int // 点击次数
	totalPv  int // 总页面访问次数
	totalUv  int // 总独立访客数
	totalUip int // 总独立IP数
	todayPv  int // 今日页面访问次数
	todayUv  int // 今日独立访客数
	todayUip int // 今日独立IP数
}

func NewInitStats() *Stats {
	return &Stats{
		clickNum: 0,
		totalPv:  0,
		totalUv:  0,
		totalUip: 0,
		todayPv:  0,
		todayUv:  0,
		todayUip: 0,
	}
}

// ClickNum 获取点击次数
func (s *Stats) ClickNum() int {
	return s.clickNum
}

// TotalPv 获取总页面访问次数
func (s *Stats) TotalPv() int {
	return s.totalPv
}

// TotalUv 获取总独立访客数
func (s *Stats) TotalUv() int {
	return s.totalUv
}

// TotalUip 获取总独立IP数
func (s *Stats) TotalUip() int {
	return s.totalUip
}

func UnmarshalStatsFromDB(
	clickNum int,
	totalPv int,
	totalUv int,
	totalUip int,
	todayPv int,
	todayUv int,
	todayUip int,
) (*Stats, error) {

	// 确保所有的数据都是非负数
	if clickNum < 0 || totalPv < 0 || totalUv < 0 ||
		totalUip < 0 || todayPv < 0 || todayUv < 0 || todayUip < 0 {
		return nil, errors.New("all stats should be non-negative")
	}

	return &Stats{
		clickNum: clickNum,
		totalPv:  totalPv,
		totalUv:  totalUv,
		totalUip: totalUip,
		todayPv:  todayPv,
		todayUv:  todayUv,
		todayUip: todayUip,
	}, nil
}
