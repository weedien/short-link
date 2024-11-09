package query

import (
	"context"
	"errors"
	"log/slog"
	"reflect"
	"shortlink/internal/base/base_event"
	"shortlink/internal/base/cache"
	"shortlink/internal/base/decorator"
	"shortlink/internal/base/errno"
	"shortlink/internal/base/metrics"
	"shortlink/internal/link/common/constant"
	"shortlink/internal/link/domain/event"
	"shortlink/internal/link/domain/link"
)

type getOriginalUrlHandler struct {
	readModel        GetOriginalUrlReadModel
	eventBus         base_event.EventBus
	distributedCache cache.DistributedCache
}

type GetOriginalUrlHandler decorator.QueryHandler[GetOriginalUrl, string]

func NewGetOriginalUrlHandler(
	readModel GetOriginalUrlReadModel,
	eventBus base_event.EventBus,
	distributedCache cache.DistributedCache,
	logger *slog.Logger,
	metrics metrics.Client,
) GetOriginalUrlHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[GetOriginalUrl, string](
		getOriginalUrlHandler{readModel: readModel, eventBus: eventBus, distributedCache: distributedCache},
		logger,
		metrics,
	)
}

type GetOriginalUrl struct {
	ShortUri      string
	UserVisitInfo event.UserVisitInfo
}

type GetOriginalUrlReadModel interface {
	GetLink(ctx context.Context, shortUri string) (*link.Link, error)
}

func (h getOriginalUrlHandler) Handle(ctx context.Context, q GetOriginalUrl) (res string, err error) {

	// $$ 发布事件 UserVisitEvent
	e := event.NewUserVisitEvent(q.UserVisitInfo)
	if err = h.eventBus.Publish(ctx, e); err != nil {
		return "", err
	}

	fetchFn := func() (res interface{}, err error) {
		lk := &link.Link{}
		if lk, err = h.readModel.GetLink(ctx, q.ShortUri); err != nil || lk == nil {
			return
		}
		originalUrl := lk.OriginalUrl()
		if originalUrl == "" {
			err = errno.LinkNotExists
			return
		}
		switch lk.Status() {
		case link.StatusActive:
			res = link.NewCacheValue(lk)
			//if _, err = res.Validate(); err != nil {
			//	return nil, err
			//}
			return
		case link.StatusExpired:
			err = errno.LinkExpired
			return
		case link.StatusForbidden:
			err = errno.LinkForbidden
			return
		case link.StatusReserved:
			err = errno.LinkReserved
			return
		default:
			err = errno.LinkNotExists
			return
		}
	}

	var result interface{}
	result, err = h.distributedCache.SafeGetWithCacheCheckFilter(
		ctx,
		constant.GotoLinkKey+q.ShortUri,
		reflect.TypeOf(link.CacheValue{}),
		fetchFn,
		constant.NeverExpire,
		cache.ShortUriCreateBloomFilter,
		q.ShortUri,
		constant.GotoIsNullLinkKey+q.ShortUri,
	)
	if err != nil {
		if errors.Is(err, errno.RedisKeyNotExist) {
			err = errno.LinkNotExists
		}
		return
	}

	if cacheValue, ok := result.(*link.CacheValue); ok {
		res = cacheValue.OriginalUrl
	}

	return
}
