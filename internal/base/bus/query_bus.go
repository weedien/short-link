package bus

import (
	"context"
	"fmt"
)

type QueryBus struct {
	handlers map[string]QueryHandler
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[string]QueryHandler),
	}
}

func (b QueryBus) Register(queryType string, handler QueryHandler) {
	b.handlers[queryType] = handler
}

func (b QueryBus) Unregister(queryType string) {
	delete(b.handlers, queryType)
}

func (b QueryBus) Dispatch(ctx context.Context, query Query) (interface{}, error) {
	queryType := fmt.Sprintf("%T", query)
	handler, ok := b.handlers[queryType]
	if !ok {
		return nil, fmt.Errorf("no query handler for %s", queryType)
	}

	// TODO 在这里输出日志和上传遥测数据
	// query 的结果是不明确的，query 的类型需要调用者去维护，有没有更合适的方式

	return handler.Handle(ctx, query)
}
