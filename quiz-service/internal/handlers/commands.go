package handlers

import (
	"context"
	"github.com/docker/cli/cli/command/commands"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"quizazz/internal/am"
	"quizazz/internal/ddd"
	"quizazz/internal/errorsotel"
	"quizazz/internal/registry"
	"quizazz/quiz-service/internal/application"
	"quizazz/quiz-service/quizspb"
	"time"
)

type commandHandlers struct {
	app application.App
}

func NewCommandHandlers(reg registry.Registry, app application.App, replyPublisher am.ReplyPublisher, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewCommandHandler(reg, replyPublisher, commandHandlers{
		app: app,
	}, mws...)
}

func RegisterCommandHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	_, err := subscriber.Subscribe(quizspb.CommandChannel, handlers, am.MessageFilter{
		quizspb.QuizCreatedCommand,
		//quizspb.CancelShoppingListCommand,
		//quizspb.InitiateShoppingCommand,
	}, am.GroupName("quiz-commands"))

	return err
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (reply ddd.Reply, err error) {
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
	case quizspb.QuizCreatedCommand:
		return h.doCreateQuiz(ctx, cmd)
		//case quizspb.CancelShoppingListCommand:
		//	return h.doCancelShoppingList(ctx, cmd)
	}

	return nil, nil
}

func (h commandHandlers) doCreateQuiz(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*quizspb.CreateQuizRequest)

	id := uuid.New().String()

	//items := make([]commands.OrderItem, 0, len(payload.GetItems()))
	//for _, item := range payload.GetItems() {
	//	items = append(items, commands.OrderItem{
	//		StoreID:   item.GetStoreId(),
	//		ProductID: item.GetProductId(),
	//		Quantity:  int(item.GetQuantity()),
	//	})
	//}

	err := h.app.CreateQuiz(ctx, commands.CreateShoppingList{
		ID:      id,
		OrderID: payload.GetOrderId(),
		Items:   items,
	})

	return ddd.NewReply(quizspb.CreatedShoppingListReply, &quizspb.CreatedShoppingList{Id: id}), err
}

func (h commandHandlers) doCancelShoppingList(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*quizspb.CancelShoppingList)

	err := h.app.CancelShoppingList(ctx, commands.CancelShoppingList{ID: payload.GetId()})

	// returning nil returns a simple Success or Failure reply; err being nil determines which
	return nil, err
}

func (h commandHandlers) doInitiateShopping(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*quizspb.InitiateShopping)

	err := h.app.InitiateShopping(ctx, commands.InitiateShopping{ID: payload.GetId()})

	// returning nil returns a simple Success or Failure reply; err being nil determines which
	return nil, err
}
