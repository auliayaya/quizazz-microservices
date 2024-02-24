package postgres

import (
	"context"
	"fmt"
	"quizazz/internal/postgres"
	"quizazz/user-service/internal/domain"
)

type UserRepository struct {
	tableName string
	db        postgres.DB
}

var _ domain.UserRepository = (*UserRepository)(nil)

func NewUserRepository(tableName string, db postgres.DB) UserRepository {
	return UserRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r UserRepository) Find(ctx context.Context, userID string) (*domain.User, error) {
	const query = "SELECT name, email, enabled FROM %s WHERE id = $1 LIMIT 1"

	user := domain.NewUser(userID)

	err := r.db.QueryRowContext(ctx, r.table(query), userID).Scan(&user.Name, &user.Email, &user.Enabled)

	return user, err
}

func (r UserRepository) Save(ctx context.Context, user *domain.User) error {
	const query = "INSERT INTO %s (id, NAME, email, enabled) VALUES ($1, $2, $3, $4)"

	_, err := r.db.ExecContext(ctx, r.table(query), user.ID(), user.Name, user.Email, user.Enabled)

	return err
}

func (r UserRepository) Update(ctx context.Context, user *domain.User) error {
	const query = "UPDATE %s SET NAME = $2, email = $3, enabled = $4 WHERE id = $1"

	_, err := r.db.ExecContext(ctx, r.table(query), user.ID(), user.Name, user.Email, user.Enabled)

	return err
}

func (r UserRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
