package app

import "shortlink/internal/link_stats/app/query"

type Application struct {
	Queries Queries
}

type Queries struct {
	GetLinkStats               query.GetLinkStatsHandler
	GroupLinkStats             query.GroupLinkStatsHandler
	GetLinkStatsAccessRecord   query.GetLinkStatsAccessRecordHandler
	GroupLinkStatsAccessRecord query.GroupLinkStatsAccessRecordHandler
}
