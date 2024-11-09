package constant

import "time"

const (
	DefaultTimeOut    = 3 * time.Second
	DefaultExpiration = 30 * time.Minute
	NeverExpire       = 0
)

const (
	SafeGetDistributionLock = "safe_get_distributed_lock_get:"

	// GotoLinkKey 短链接跳转前缀 Key
	GotoLinkKey = "short-link:goto:"

	LinkGroupCountKey = "short-link:group-count:"

	// GotoIsNullLinkKey 短链接空值跳转前缀 Key
	//
	// 在某次查询中，当发现短链接不存在或者已经失效，就会将这个 Key 写入缓存，从而避免查询数据库
	GotoIsNullLinkKey = "short-link:is-null:goto:"

	// LockGotoLinkKey 短链接跳转锁前缀 Key
	LockGotoLinkKey = "short-link:lock:goto:"

	// LockGidUpdateKey 短链接修改分组 ID 锁前缀 Key
	LockGidUpdateKey = "short-link:lock:update-gid"

	// DelayQueueStatKey 短链接延迟队列消费统计 Key
	DelayQueueStatKey = "short-link:delay-queue:stats"

	// LinkStatsUvKey 短链接统计判断是否新用户缓存标识
	LinkStatsUvKey = "short-link:stats:uv:"

	// LinkStatsUipKey 短链接统计判断是否新 IP 缓存标识
	LinkStatsUipKey = "short-link:stats:uip:"

	// LinkStatsStreamTopicKey 短链接监控消息保存队列 Topic 缓存标识
	LinkStatsStreamTopicKey = "short-link:stats-stream"

	// LinkStatsStreamGroupKey 短链接监控消息保存队列 Group 缓存标识
	LinkStatsStreamGroupKey = "short-link:stats-stream:only-group"

	// LinkCreateLockKey 创建短链接锁标识
	LinkCreateLockKey = "short-link:lock:create:%s"
)
