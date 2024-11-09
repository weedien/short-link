package dao

import "time"

type LinkQueryParam struct {
	FullShortUrl string
	Gid          string
	Status       string
	StartDate    time.Time
	EndDate      time.Time
}

type LinkGroupQueryParam struct {
	Gid       string
	Status    string
	StartDate time.Time
	EndDate   time.Time
}
