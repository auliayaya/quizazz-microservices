package domain

import (
	"github.com/stackus/errors"
	"quizazz/internal/ddd"
)

const QuizAggregate = "quiz.QuizAggregate"
const QuizQuestionAggregate = "quiz.QuizQuestionAggregate"
const QuizAnswerAggregate = "quiz.QuizAnswerAggregate"

type Quiz struct {
	ddd.Aggregate
	QuizName string
	QuizType string
	Status   string
}

type Question struct {
	ddd.Aggregate
	Question string
}
type Answer struct {
	ddd.Aggregate
	Answers []string
	Keys    string
}

var (
	ErrNameCannotBeBlank     = errors.Wrap(errors.ErrBadRequest, "the quiz name cannot be blank")
	ErrQuizIDCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the quiz id cannot be blank")
	ErrQuizTypeCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the SMS number cannot be blank")
	ErrQuizAlreadyEnabled    = errors.Wrap(errors.ErrBadRequest, "the quiz is already enabled")
	ErrQuizAlreadyDisabled   = errors.Wrap(errors.ErrBadRequest, "the quiz is already disabled")
	ErrQuizNotAuthorized     = errors.Wrap(errors.ErrUnauthorized, "quiz is not authorized")
)

func NewQuiz(id string) *Quiz {
	return &Quiz{
		Aggregate: ddd.NewAggregate(id, QuizAggregate),
	}
}

func CreatedQuiz(id, name, quizType string) (*Quiz, error) {
	if id == "" {
		return nil, ErrQuizIDCannotBeBlank
	}
	if name == "" {
		return nil, ErrNameCannotBeBlank
	}
	if quizType == "" {
		return nil, ErrQuizTypeCannotBeBlank
	}
	quiz := NewQuiz(id)
	quiz.QuizName = name
	quiz.QuizType = quizType

	quiz.AddEvent(QuizCreatedEvent, &QuizCreated{
		Quiz: quiz,
	})
	return quiz, nil
}

func (Quiz) Key() string {
	return QuizAggregate
}

func (u *Quiz) Authorize( /* TODO authorize what? */ ) error {
	if u.QuizType != "" {
		return ErrQuizNotAuthorized
	}
	u.AddEvent(QuizAuthorizedEvent, &QuizAuthorized{
		Quiz: u,
	})
	return nil
}

func (u *Quiz) Enable() error {
	//if u.Enabled {
	//	return ErrQuizAlreadyEnabled
	//}
	//u.Enabled = true
	u.AddEvent(QuizEnabledEvent, &QuizEnabled{
		Quiz: u,
	})
	return nil
}

func (u *Quiz) Disable() error {
	//if !u.Enabled {
	//	return ErrQuizAlreadyDisabled
	//}
	//u.Enabled = false
	u.AddEvent(QuizDisabledEvent, &QuizDisabled{
		Quiz: u,
	})
	return nil
}
