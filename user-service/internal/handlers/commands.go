package handlers

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"quizazz/internal/am"
	"quizazz/internal/ddd"
	"quizazz/internal/errorsotel"
	"quizazz/internal/registry"
	"quizazz/user-service/internal/application"
	"quizazz/user-service/userspb"
	"time"
)

type commandHandlers struct {
	app application.App
}

func RegisterCommandHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	_, err := subscriber.Subscribe(userspb.CommandChannel, handlers, am.MessageFilter{
		userspb.AuthorizeUserCommand,
	}, am.GroupName("user-commands"))
	return err
}

func (c commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (reply ddd.Reply, err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling command",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled command", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling command", trace.WithAttributes(
		attribute.String("Command", cmd.CommandName()),
	))

	switch cmd.CommandName() {
	case userspb.AuthorizeUserCommand:
		return c.doAuthorizeUser(ctx, cmd)
	}

	return nil, nil
}
func (c commandHandlers) doAuthorizeUser(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*userspb.AuthorizeUser)

	return nil, c.app.AuthorizeUser(ctx, application.AuthorizeUser{ID: payload.GetId()})
}

func NewCommandHandlers(reg registry.Registry, app application.App, replyPublisher am.ReplyPublisher, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewCommandHandler(reg, replyPublisher, commandHandlers{
		app: app,
	}, mws...)
}
