package postgres

import (
	"context"
	"fmt"
	"quizazz/internal/postgres"
	"quizazz/quiz-service/internal/domain"
)

type QuizRepository struct {
	tableName string
	db        postgres.DB
}

var _ domain.QuizRepository = (*QuizRepository)(nil)

func NewQuizRepository(tableName string, db postgres.DB) QuizRepository {
	return QuizRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r QuizRepository) Find(ctx context.Context, quizID string) (*domain.Quiz, error) {
	const query = "SELECT name, quiz_type, status FROM %s WHERE id = $1 LIMIT 1"

	quiz := domain.NewQuiz(quizID)

	err := r.db.QueryRowContext(ctx, r.table(query), quizID).Scan(&quiz.QuizName, &quiz.QuizType, &quiz.Status)

	return quiz, err
}

func (r QuizRepository) Save(ctx context.Context, quiz *domain.Quiz) error {
	const query = "INSERT INTO %s (id, NAME, quiz_type, status) VALUES ($1, $2, $3, $4)"

	_, err := r.db.ExecContext(ctx, r.table(query), quiz.ID(), quiz.QuizName, quiz.QuizType, quiz.Status)

	return err
}

func (r QuizRepository) Update(ctx context.Context, quiz *domain.Quiz) error {
	const query = "UPDATE %s SET NAME = $2, quiz_type = $3, status = $4 WHERE id = $1"

	_, err := r.db.ExecContext(ctx, r.table(query), quiz.ID(), quiz.QuizName, quiz.QuizType, quiz.Status)

	return err
}

func (r QuizRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
