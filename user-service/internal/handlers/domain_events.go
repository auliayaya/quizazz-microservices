package handlers

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"quizazz/internal/am"
	"quizazz/internal/ddd"
	"quizazz/internal/errorsotel"
	"quizazz/user-service/internal/domain"
	"quizazz/user-service/userspb"
	"time"
)

type domainHandlers[T ddd.AggregateEvent] struct {
	publisher am.EventPublisher
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*domainHandlers[ddd.AggregateEvent])(nil)

func NewDomainEventHandlers(publisher am.EventPublisher) ddd.EventHandler[ddd.AggregateEvent] {
	return &domainHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.AggregateEvent], handlers ddd.EventHandler[ddd.AggregateEvent]) {
	subscriber.Subscribe(handlers,
		domain.UserRegisteredEvent,
		domain.UserEmailChangedEvent,
		domain.UserEnabledEvent,
		domain.UserDisabledEvent,
	)
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling domain event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled domain event", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling domain event", trace.WithAttributes(
		attribute.String("Event", event.EventName()),
	))

	switch event.EventName() {
	case domain.UserRegisteredEvent:
		return h.onUserRegistered(ctx, event)
	case domain.UserEmailChangedEvent:
		return h.onUserEmailChanged(ctx, event)
	case domain.UserEnabledEvent:
		return h.onUserEnabled(ctx, event)
	case domain.UserDisabledEvent:
		return h.onUserDisabled(ctx, event)
	}
	return nil
}

func (h domainHandlers[T]) onUserRegistered(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.UserRegistered)
	return h.publisher.Publish(ctx, userspb.UserAggregateChannel,
		ddd.NewEvent(userspb.UserRegisteredEvent, &userspb.UserRegistered{
			Id:    payload.User.ID(),
			Name:  payload.User.Name,
			Email: payload.User.Email,
		}),
	)
}

func (h domainHandlers[T]) onUserEmailChanged(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.UserRegistered)
	return h.publisher.Publish(ctx, userspb.UserAggregateChannel,
		ddd.NewEvent(userspb.UserEmailChangedEvent, &userspb.UserEmailChanged{
			Id:    payload.User.ID(),
			Email: payload.User.Email,
		}),
	)
}

func (h domainHandlers[T]) onUserEnabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, userspb.UserAggregateChannel,
		ddd.NewEvent(userspb.UserEnabledEvent, &userspb.UserEnabled{
			Id: event.AggregateID(),
		}),
	)
}

func (h domainHandlers[T]) onUserDisabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, userspb.UserAggregateChannel,
		ddd.NewEvent(userspb.UserDisabledEvent, &userspb.UserDisabled{
			Id: event.AggregateID(),
		}),
	)
}
