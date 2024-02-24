package domain

const (
	QuizCreatedEvent      = "quiz.QuizCreated"
	QuizEmailChangedEvent = "quiz.QuizEmailChanged"
	QuizAuthorizedEvent   = "quiz.QuizAuthorized"
	QuizEnabledEvent      = "quiz.QuizEnabled"
	QuizDisabledEvent     = "quiz.QuizDisabled"
)

type QuizCreated struct {
	Quiz *Quiz
}

func (QuizCreated) Key() string {
	return QuizCreatedEvent
}

type QuizEmailChanged struct {
	Quiz *Quiz
}

func (QuizEmailChanged) Key() string {
	return QuizEmailChangedEvent
}

type QuizAuthorized struct {
	Quiz *Quiz
}

func (QuizAuthorized) Key() string {
	return QuizAuthorizedEvent
}

type QuizEnabled struct {
	Quiz *Quiz
}

func (QuizEnabled) Key() string {
	return QuizEnabledEvent
}

type QuizDisabled struct {
	Quiz *Quiz
}

func (QuizDisabled) Key() string {
	return QuizDisabledEvent
}
