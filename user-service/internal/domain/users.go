package domain

import (
	"github.com/stackus/errors"
	"quizazz/internal/ddd"
)

const UserAggregate = "users.UserAggregate"

type User struct {
	ddd.Aggregate
	Name    string
	Email   string
	Enabled bool
}

var (
	ErrNameCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the user name cannot be blank")
	ErrUserIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the user id cannot be blank")
	ErrEmailCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the SMS number cannot be blank")
	ErrUserAlreadyEnabled  = errors.Wrap(errors.ErrBadRequest, "the user is already enabled")
	ErrUserAlreadyDisabled = errors.Wrap(errors.ErrBadRequest, "the user is already disabled")
	ErrUserNotAuthorized   = errors.Wrap(errors.ErrUnauthorized, "user is not authorized")
)

func NewUser(id string) *User {
	return &User{
		Aggregate: ddd.NewAggregate(id, UserAggregate),
	}
}

func RegisterUser(id, name, email string) (*User, error) {
	if id == "" {
		return nil, ErrUserIDCannotBeBlank
	}
	if name == "" {
		return nil, ErrNameCannotBeBlank
	}
	if email == "" {
		return nil, ErrEmailCannotBeBlank
	}
	user := NewUser(id)
	user.Name = name
	user.Email = email
	user.Enabled = true
	user.AddEvent(UserRegisteredEvent, &UserRegistered{
		User: user,
	})
	return user, nil
}

func (User) Key() string {
	return UserAggregate
}

func (u *User) Authorize( /* TODO authorize what? */ ) error {
	if !u.Enabled {
		return ErrUserNotAuthorized
	}
	u.AddEvent(UserAuthorizedEvent, &UserAuthorized{
		User: u,
	})
	return nil
}

func (u *User) Enable() error {
	if u.Enabled {
		return ErrUserAlreadyEnabled
	}
	u.Enabled = true
	u.AddEvent(UserEnabledEvent, &UserEnabled{
		User: u,
	})
	return nil
}

func (u *User) Disable() error {
	if !u.Enabled {
		return ErrUserAlreadyDisabled
	}
	u.Enabled = false
	u.AddEvent(UserDisabledEvent, &UserDisabled{
		User: u,
	})
	return nil
}
