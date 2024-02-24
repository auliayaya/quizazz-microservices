package quizspb

import (
	"quizazz/internal/registry"
	"quizazz/internal/registry/serdes"
)

const (
	QuizAggregateChannel = "mallbots.quizs.events.Quiz"

	QuizCreatedCommand      = "quizsapi.QuizCreatedCommand"
	QuizEmailChangedCommand = "quizsapi.QuizEmailChangedCommand"
	QuizEnabledCommand      = "quizsapi.QuizEnabledCommand"
	QuizDisabledCommand     = "quizsapi.QuizDisabledCommand"

	CommandChannel = "mallbots.quizs.commands"

	AuthorizeQuizCommand = "quizsapi.AuthorizeQuiz"
)

func Registrations(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	// Quiz events
	if err := serde.Register(&QuizCreated{}); err != nil {
		return err
	}
	if err := serde.Register(&QuizEmailChanged{}); err != nil {
		return err
	}
	if err := serde.Register(&QuizEnabled{}); err != nil {
		return err
	}
	if err := serde.Register(&QuizDisabled{}); err != nil {
		return err
	}

	// commands
	if err := serde.Register(&AuthorizeQuiz{}); err != nil {
		return err
	}
	return nil
}

func (*QuizCreated) Key() string      { return QuizCreatedCommand }
func (*QuizEmailChanged) Key() string { return QuizEmailChangedCommand }
func (*QuizEnabled) Key() string      { return QuizEnabledCommand }
func (*QuizDisabled) Key() string     { return QuizDisabledCommand }

func (*AuthorizeQuiz) Key() string { return AuthorizeQuizCommand }
