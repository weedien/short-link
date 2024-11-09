package app

import (
	"shortlink/internal/link/app/command"
	"shortlink/internal/link/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateLink      command.CreateLinkHandler
	CreateLinkBatch command.CreateLinkBatchHandler
	UpdateLink      command.UpdateLinkHandler

	SaveToRecycleBin      command.SaveToRecycleBinHandler
	RemoveFromRecycleBin  command.RemoveFromRecycleBinHandler
	RecoverFromRecycleBin command.RecoverFromRecycleBinHandler
}

type Queries struct {
	PageLink       query.PageLinkHandler
	ListGroupCount query.ListGroupCountHandler
	GetOriginalUrl query.GetOriginalUrlHandler

	PageRecycleBin query.PageRecycleBinHandler
}
