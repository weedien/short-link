package bus

import "context"

type Command interface{}

type Query interface{}

type CommandHandler interface {
	Handle(ctx context.Context, command Command) error
}

type QueryHandler interface {
	Handle(ctx context.Context, query Query) (interface{}, error)
}
