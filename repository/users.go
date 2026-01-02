package repository

import (
	"context"

	"github.com/fabianoflorentino/golangfromzero/src/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUsersRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db}
}

func (r UserRepository) Create(user models.User) (uuid.UUID, error) {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	ctx := context.Background()

	var id uuid.UUID

	if err := r.db.QueryRow(ctx, query, user.Name, user.Email, user.Password).Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
