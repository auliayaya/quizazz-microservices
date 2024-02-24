package handlers

import (
	"context"
	"quizazz/internal/ddd"
	"quizazz/internal/di"
	"quizazz/user-service/internal/constants"
)

func RegisterDomainEventHandlersTx(container di.Container) {
	handlers := ddd.EventHandlerFunc[ddd.AggregateEvent](func(ctx context.Context, event ddd.AggregateEvent) error {
		domainHandlers := di.Get(ctx, constants.DomainEventHandlersKey).(ddd.EventHandler[ddd.AggregateEvent])

		return domainHandlers.HandleEvent(ctx, event)
	})

	subscriber := container.Get(constants.DomainDispatcherKey).(*ddd.EventDispatcher[ddd.AggregateEvent])

	RegisterDomainEventHandlers(subscriber, handlers)
}
