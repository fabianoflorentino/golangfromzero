package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

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
		if strings.Contains(err.Error(), "users_email_key") {
			return uuid.Nil, errors.New("email already used")
		}

		return uuid.Nil, err
	}

	return id, nil
}

func (r UserRepository) SearchByName(name string) ([]models.User, error) {
	query := `SELECT id, name, email, created_at from users WHERE name LIKE $1`

	var users []models.User
	ctx := context.Background()

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

func (r UserRepository) SearchByID(id uuid.UUID) (*models.User, error) {
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`

	var user models.User
	ctx := context.Background()

	if err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
		return nil, err
	}
	r.db.Close()

	return &user, nil
}

func (r UserRepository) Update(id uuid.UUID, user models.User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`

	ctx := context.Background()

	if _, err := r.db.Exec(ctx, query, user.Name, user.Email, id); err != nil {
		return err
	}
	r.db.Close()

	return nil
}

func (r UserRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`

	ctx := context.Background()

	if _, err := r.db.Exec(ctx, query, id); err != nil {
		return err
	}
	r.db.Close()

	return nil
}
