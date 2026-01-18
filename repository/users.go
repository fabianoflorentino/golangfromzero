package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/fabianoflorentino/golangfromzero/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUsersRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db}
}

func (r UserRepository) Create(ctx context.Context, user models.User) (uuid.UUID, error) {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`

	var id uuid.UUID

	if err := r.db.QueryRow(ctx, query, user.Name, user.Email, user.Password).Scan(&id); err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			return uuid.Nil, errors.New("email already used")
		}

		return uuid.Nil, err
	}

	return id, nil
}

// SearchByName search a user using a string to filter and find a user by name.
func (r UserRepository) SearchByName(ctx context.Context, name string) ([]models.User, error) {
	query := `SELECT id, name, email, created_at from users WHERE name LIKE $1`

	var users []models.User

	name = fmt.Sprintf("%%%s%%", name)

	rows, err := r.db.Query(ctx, query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// SearchByID search a user using a id (UUID) to filter and find a user by ID.
func (r UserRepository) SearchByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`

	var user models.User

	if err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

// Update update's user information; can be update name and email
func (r UserRepository) Update(ctx context.Context, id uuid.UUID, user models.User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`

	if _, err := r.db.Exec(ctx, query, user.Name, user.Email, id); err != nil {
		return err
	}

	return nil
}

// Delete delete's user from database.
func (r UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`

	if _, err := r.db.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}
