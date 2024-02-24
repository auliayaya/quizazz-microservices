package application

import (
	"context"
	"quizazz/internal/ddd"
	"quizazz/quiz-service/internal/domain"
)

type (
	CreateQuiz struct {
		ID       string
		QuizName string
		QuizType string
	}
	AuthorizeQuiz struct {
		ID string
	}
	GetQuiz struct {
		ID string
	}
	EnableQuiz struct {
		ID string
	}
	DisableQuiz struct {
		ID string
	}
	App interface {
		CreateQuiz(ctx context.Context, register CreateQuiz) error
		AuthorizeQuiz(ctx context.Context, authorize AuthorizeQuiz) error
		GetQuiz(ctx context.Context, get GetQuiz) (*domain.Quiz, error)
		EnableQuiz(ctx context.Context, enable EnableQuiz) error
		DisableQuiz(ctx context.Context, disable DisableQuiz) error
	}
	Application struct {
		quizzes         domain.QuizRepository
		domainPublisher ddd.EventPublisher[ddd.AggregateEvent]
	}
)

func New(quizzes domain.QuizRepository, domainPublisher ddd.EventPublisher[ddd.AggregateEvent]) *Application {
	return &Application{
		quizzes:         quizzes,
		domainPublisher: domainPublisher,
	}
}

func (a Application) CreateQuiz(ctx context.Context, register CreateQuiz) error {
	quiz, err := domain.CreatedQuiz(register.ID, register.QuizName, register.QuizType)
	if err != nil {
		return err
	}
	if err = a.quizzes.Save(ctx, quiz); err != nil {
		return err
	}
	if err = a.domainPublisher.Publish(ctx, quiz.Events()...); err != nil {
		return err
	}
	return nil
}

func (a Application) AuthorizeQuiz(ctx context.Context, authorize AuthorizeQuiz) error {
	user, err := a.quizzes.Find(ctx, authorize.ID)
	if err != nil {
		return err
	}
	if err = user.Authorize(); err != nil {
		return err
	}
	if err = a.domainPublisher.Publish(ctx, user.Events()...); err != nil {
		return err
	}
	return nil
}

func (a Application) GetQuiz(ctx context.Context, get GetQuiz) (*domain.Quiz, error) {
	return a.quizzes.Find(ctx, get.ID)
}

func (a Application) EnableQuiz(ctx context.Context, enable EnableQuiz) error {
	user, err := a.quizzes.Find(ctx, enable.ID)
	if err != nil {
		return err
	}
	if err = user.Enable(); err != nil {
		return err
	}
	if err = a.quizzes.Update(ctx, user); err != nil {
		return err
	}
	if err = a.domainPublisher.Publish(ctx, user.Events()...); err != nil {
		return err
	}
	return nil
}

func (a Application) DisableQuiz(ctx context.Context, disable DisableQuiz) error {
	user, err := a.quizzes.Find(ctx, disable.ID)
	if err != nil {
		return err
	}
	if err = user.Disable(); err != nil {
		return err
	}
	if err = a.quizzes.Update(ctx, user); err != nil {
		return err
	}
	return nil
}

var _ App = (*Application)(nil)
