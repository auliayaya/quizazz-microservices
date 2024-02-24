package handlers

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"quizazz/internal/am"
	"quizazz/internal/ddd"
	"quizazz/internal/errorsotel"
	"quizazz/quiz-service/internal/domain"
	"quizazz/quiz-service/quizspb"
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
		domain.QuizCreatedEvent,
		domain.QuizEmailChangedEvent,
		domain.QuizEnabledEvent,
		domain.QuizDisabledEvent,
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
	case domain.QuizCreatedEvent:
		return h.onQuizCreatedd(ctx, event)
	case domain.QuizEmailChangedEvent:
		return h.onQuizEmailChanged(ctx, event)
	case domain.QuizEnabledEvent:
		return h.onQuizEnabled(ctx, event)
	case domain.QuizDisabledEvent:
		return h.onQuizDisabled(ctx, event)
	}
	return nil
}

func (h domainHandlers[T]) onQuizCreatedd(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.QuizCreated)
	return h.publisher.Publish(ctx, quizspb.QuizAggregateChannel,
		ddd.NewEvent(quizspb.QuizCreatedCommand, &quizspb.QuizCreated{
			Id:       payload.Quiz.ID(),
			Name:     payload.Quiz.QuizName,
			QuizType: payload.Quiz.QuizType,
		}),
	)
}

func (h domainHandlers[T]) onQuizEmailChanged(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.QuizCreated)
	return h.publisher.Publish(ctx, quizspb.QuizAggregateChannel,
		ddd.NewEvent(quizspb.QuizEmailChangedCommand, &quizspb.QuizEmailChanged{
			Id:       payload.Quiz.ID(),
			QuizType: payload.Quiz.QuizType,
		}),
	)
}

func (h domainHandlers[T]) onQuizEnabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, quizspb.QuizAggregateChannel,
		ddd.NewEvent(quizspb.QuizEnabledCommand, &quizspb.QuizEnabled{
			Id: event.AggregateID(),
		}),
	)
}

func (h domainHandlers[T]) onQuizDisabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, quizspb.QuizAggregateChannel,
		ddd.NewEvent(quizspb.QuizDisabledCommand, &quizspb.QuizDisabled{
			Id: event.AggregateID(),
		}),
	)
}
