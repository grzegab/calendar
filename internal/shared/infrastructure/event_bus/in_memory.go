package event_bus

import (
	"context"
	"sync"
)

type HandlerFunc func(ctx context.Context, event any) error

type InMemoryBus struct {
	mu       sync.RWMutex
	handlers []HandlerFunc
}

func NewInMemory() *InMemoryBus {
	return &InMemoryBus{
		handlers: make([]HandlerFunc, 0),
	}
}

func (b *InMemoryBus) Subscribe(handler HandlerFunc) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.handlers = append(b.handlers, handler)
}
