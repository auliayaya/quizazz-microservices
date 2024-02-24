package domain

const (
	UserRegisteredEvent   = "users.UserRegistered"
	UserEmailChangedEvent = "users.UserEmailChanged"
	UserAuthorizedEvent   = "users.UserAuthorized"
	UserEnabledEvent      = "users.UserEnabled"
	UserDisabledEvent     = "users.UserDisabled"
)

type UserRegistered struct {
	User *User
}

func (UserRegistered) Key() string {
	return UserRegisteredEvent
}

type UserEmailChanged struct {
	User *User
}

func (UserEmailChanged) Key() string {
	return UserEmailChangedEvent
}

type UserAuthorized struct {
	User *User
}

func (UserAuthorized) Key() string {
	return UserAuthorizedEvent
}

type UserEnabled struct {
	User *User
}

func (UserEnabled) Key() string {
	return UserEnabledEvent
}

type UserDisabled struct {
	User *User
}

func (UserDisabled) Key() string {
	return UserDisabledEvent
}
