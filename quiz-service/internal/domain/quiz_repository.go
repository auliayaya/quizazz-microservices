package domain

import "context"

type QuizRepository interface {
	Save(ctx context.Context, quiz *Quiz) error
	Find(ctx context.Context, quizID string) (*Quiz, error)
	Update(ctx context.Context, quiz *Quiz) error
}
