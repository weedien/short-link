package bus

import (
	"context"
	"fmt"
)

type CommandBus struct {
	handlers map[string]CommandHandler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]CommandHandler),
	}
}

func (bus *CommandBus) Register(name string, handler CommandHandler) {
	bus.handlers[name] = handler
}

func (bus *CommandBus) Unregister(name string) {
	delete(bus.handlers, name)
}

func (bus *CommandBus) Dispatch(ctx context.Context, cmd Command) error {
	commandType := fmt.Sprintf("%T", cmd)
	handler, ok := bus.handlers[commandType]
	if !ok {
		return fmt.Errorf("no command handler for %s", commandType)
	}
	return handler.Handle(ctx, cmd)
}
