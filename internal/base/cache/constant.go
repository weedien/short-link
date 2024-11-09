package cache

import "time"

const (
	DefaultTimeOut    = 3 * time.Second
	DefaultExpiration = 30 * time.Minute
	NeverExpire       = 0
)
