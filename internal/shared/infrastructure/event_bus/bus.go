package event_bus

type Bus interface {
	Publish(event any)
}

type EventBus interface {
	Subscribe(handler HandlerFunc)
}
