package application

import (
	"context"
	"quizazz/internal/ddd"
	"quizazz/user-service/internal/domain"
)

type (
	RegisterUser struct {
		ID    string
		Name  string
		Email string
	}
	AuthorizeUser struct {
		ID string
	}
	GetUser struct {
		ID string
	}
	EnableUser struct {
		ID string
	}
	DisableUser struct {
		ID string
	}
	App interface {
		RegisterUser(ctx context.Context, register RegisterUser) error
		AuthorizeUser(ctx context.Context, authorize AuthorizeUser) error
		GetUser(ctx context.Context, get GetUser) (*domain.User, error)
		EnableUser(ctx context.Context, enable EnableUser) error
		DisableUser(ctx context.Context, disable DisableUser) error
	}
	Application struct {
		users           domain.UserRepository
		domainPublisher ddd.EventPublisher[ddd.AggregateEvent]
	}
)

func New(users domain.UserRepository, domainPublisher ddd.EventPublisher[ddd.AggregateEvent]) *Application {
	return &Application{
		users:           users,
		domainPublisher: domainPublisher,
	}
}

func (a Application) RegisterUser(ctx context.Context, register RegisterUser) error {
	user, err := domain.RegisterUser(register.ID, register.Name, register.Email)
	if err != nil {
		return err
	}
	if err = a.users.Save(ctx, user); err != nil {
		return err
	}
	if err = a.domainPublisher.Publish(ctx, user.Events()...); err != nil {
		return err
	}
	return nil
}

func (a Application) AuthorizeUser(ctx context.Context, authorize AuthorizeUser) error {
	user, err := a.users.Find(ctx, authorize.ID)
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

func (a Application) GetUser(ctx context.Context, get GetUser) (*domain.User, error) {
	return a.users.Find(ctx, get.ID)
}

func (a Application) EnableUser(ctx context.Context, enable EnableUser) error {
	user, err := a.users.Find(ctx, enable.ID)
	if err != nil {
		return err
	}
	if err = user.Enable(); err != nil {
		return err
	}
	if err = a.users.Update(ctx, user); err != nil {
		return err
	}
	if err = a.domainPublisher.Publish(ctx, user.Events()...); err != nil {
		return err
	}
	return nil
}

func (a Application) DisableUser(ctx context.Context, disable DisableUser) error {
	user, err := a.users.Find(ctx, disable.ID)
	if err != nil {
		return err
	}
	if err = user.Disable(); err != nil {
		return err
	}
	if err = a.users.Update(ctx, user); err != nil {
		return err
	}
	return nil
}

var _ App = (*Application)(nil)
