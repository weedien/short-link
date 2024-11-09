package service

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log/slog"
	"shortlink/internal/base/base_event"
	"shortlink/internal/base/cache"
	"shortlink/internal/base/lock"
	"shortlink/internal/base/metrics"
	"shortlink/internal/link/adapter"
	"shortlink/internal/link/adapter/read"
	"shortlink/internal/link/app"
	"shortlink/internal/link/app/command"
	"shortlink/internal/link/app/query"
	"shortlink/internal/link/common/config"
	"shortlink/internal/link/domain/link"
)

func NewLinkApplication(
	db *gorm.DB,
	rdb *redis.Client,
	locker lock.DistributedLock,
	eventBus base_event.EventBus,
) (a app.Application) {

	logger := slog.Default()
	metricsClient := metrics.NoOp{}

	c := config.Get().AppLink

	factoryConfig := link.FactoryConfig{
		Domain:            c.Domain,
		UseSSL:            config.Get().Server.UseSsl,
		Whitelist:         c.Whitelist,
		MaxAttempts:       c.MaxAttempts,
		MaxLinksPerGroup:  c.MaxLinksPerGroup,
		DefaultFavicon:    c.DefaultFaviconUrl,
		DefaultGid:        "isChfh",
		DefaultExpiration: 15,
		DefaultCreateType: link.CreateByApi,
		DefaultValidType:  link.ValidTypeTemporary,
	}
	linkFactory, err := link.NewFactory(factoryConfig)
	if err != nil {
		panic("failed to create link factory: " + err.Error())
	}

	distributedCache := cache.NewRedisDistributedCache(rdb, locker)
	repository := adapter.NewLinkRepository(linkFactory, db, distributedCache)
	readModel := read.NewLinkQuery(db, linkFactory, distributedCache)

	a = app.Application{
		Commands: app.Commands{
			CreateLink:      command.NewCreateLinkHandler(linkFactory, repository, locker, logger, metricsClient),
			CreateLinkBatch: command.NewCreateLinkBatchHandler(linkFactory, repository, logger, metricsClient),
			UpdateLink:      command.NewUpdateLinkHandler(repository, logger, metricsClient),

			SaveToRecycleBin:      command.NewSaveToRecycleBinHandler(repository, logger, metricsClient),
			RemoveFromRecycleBin:  command.NewRemoveFromRecycleBinHandler(repository, logger, metricsClient),
			RecoverFromRecycleBin: command.NewRecoverFromRecycleBinHandler(repository, logger, metricsClient),
		},
		Queries: app.Queries{
			PageLink:       query.NewPageLinkHandler(readModel, logger, metricsClient),
			ListGroupCount: query.NewListGroupCountHandler(readModel, logger, metricsClient),
			GetOriginalUrl: query.NewGetOriginalUrlHandler(readModel, eventBus, distributedCache, logger, metricsClient),

			PageRecycleBin: query.NewPageRecycleBinHandler(readModel, logger, metricsClient),
		},
	}

	return
}
