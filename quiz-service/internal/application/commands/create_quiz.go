package commands

import (
	"context"
	"github.com/stackus/errors"
	"quizazz/internal/ddd"
	"quizazz/quiz-service/internal/domain"
)

type CreateQuiz struct {
	ID       string
	QuizName string
	QuizType string
}

type CreateQuizHandler struct {
	shoppingLists   domain.QuizRepository
	domainPublisher ddd.EventPublisher[ddd.AggregateEvent]
}

func NewCreateQuizHandler(shoppingLists domain.QuizRepository, domainPublisher ddd.EventPublisher[ddd.AggregateEvent],
) CreateQuizHandler {
	return CreateQuizHandler{
		shoppingLists:   shoppingLists,
		domainPublisher: domainPublisher,
	}
}

func (h CreateQuizHandler) CreateQuiz(ctx context.Context, cmd CreateQuiz) error {
	list, _ := domain.CreatedQuiz(cmd.ID, cmd.QuizName, cmd.QuizType)

	//for _, item := range cmd.Items {
	//	// horribly inefficient
	//	store, err := h.stores.Find(ctx, item.StoreID)
	//	if err != nil {
	//		return errors.Wrap(err, "building shopping list")
	//	}
	//	product, err := h.products.Find(ctx, item.ProductID)
	//	if err != nil {
	//		return errors.Wrap(err, "building shopping list")
	//	}
	//	err = list.AddItem(store, product, item.Quantity)
	//	if err != nil {
	//		return errors.Wrap(err, "building shopping list")
	//	}
	//}

	if err := h.shoppingLists.Save(ctx, list); err != nil {
		return errors.Wrap(err, "scheduling quiz")
	}

	// publish domain events
	if err := h.domainPublisher.Publish(ctx, list.Events()...); err != nil {
		return err
	}

	return nil
}
