package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/fabianoflorentino/golangfromzero/internal/models"
	"github.com/fabianoflorentino/golangfromzero/internal/response"
	"github.com/fabianoflorentino/golangfromzero/repository"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TimeoutConfig struct {
	DatabaseTimeout time.Duration
}

var DefaultTimout = TimeoutConfig{
	DatabaseTimeout: 5 * time.Second,
}

// UserController represents a user controller that receveis a configuration and database connections
type UserController struct {
	db     *pgxpool.Pool
	repo   *repository.UserRepository
	logger *slog.Logger
}

// NewUserController initialize a new controller configuration and database connection
func NewUserController(db *pgxpool.Pool, repo *repository.UserRepository, logger *slog.Logger) *UserController {
	return &UserController{
		db:     db,
		repo:   repo,
		logger: logger,
	}
}

// Create handles the creation of a new user
func (user *UserController) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), DefaultTimout.DatabaseTimeout)
	defer cancel()

	requestBody, err := io.ReadAll(r.Body)

	user.logger.InfoContext(ctx, "create user request started",
		"method", r.Method,
		"path", r.URL.Path,
	)

	if err != nil {
		user.logger.ErrorContext(ctx, "failed to read the request body",
			"method", r.Method,
			"path", r.URL.Path,
			"error", err,
		)

		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var u models.User

	if err := json.Unmarshal(requestBody, &u); err != nil {
		user.logger.ErrorContext(ctx, "failed to unmarshal the request body",
			"method", r.Method,
			"path", r.URL.Path,
			"error", err,
		)

		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := u.Validate(models.ValidationCreate); err != nil {
		user.logger.ErrorContext(ctx, "failed to validate the user data",
			"method", r.Method,
			"path", r.URL.Path,
			"error", err,
		)

		response.Err(w, http.StatusBadRequest, err)
		return
	}

	u.ID, err = user.repo.Create(ctx, u)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrEmailAlreadyExist):
			user.logger.ErrorContext(ctx, repository.ErrEmailAlreadyExist.Error(),
				"method", r.Method,
				"path", r.URL.Path,
				"error", err,
			)

			response.Err(w, http.StatusConflict, err)
			return
		default:
			user.logger.ErrorContext(ctx, "failed to create user",
				"method", r.Method,
				"path", r.URL.Path,
				"error", err,
			)

			response.Err(w, http.StatusInternalServerError, err)
			return
		}
	}

	user.logger.InfoContext(ctx, "user created successfully",
		"method", r.Method,
		"path", r.URL.Path,
		"user_id", u.ID,
	)

	w.WriteHeader(http.StatusCreated)
}

// SearchByName retrieves all users with contains a search name
func (user *UserController) SearchByName(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), DefaultTimout.DatabaseTimeout)
	defer cancel()

	nameToSearch := strings.ToLower(r.URL.Query().Get("name"))

	user.logger.InfoContext(ctx, "search users request started",
		"method", r.Method,
		"path", r.URL.Path,
		"has_name_filter", nameToSearch != "",
	)

	users, err := user.repo.SearchByName(ctx, nameToSearch)
	if err != nil {
		user.logger.ErrorContext(ctx, "failed to search users",
			"method", r.Method,
			"path", r.URL.Path,
			"error", err,
		)

		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	if len(users) == 0 {
		user.logger.InfoContext(ctx, "no users found",
			"method", r.Method,
			"path", r.URL.Path,
		)

		response.JSON(w, http.StatusOK, "there no users found")
		return
	}

	user.logger.InfoContext(ctx, "users found",
		"method", r.Method,
		"path", r.URL.Path,
		"count", len(users),
	)

	response.JSON(w, http.StatusOK, users)
}

// SearchByID retrieves a user by their ID
func (user *UserController) SearchByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), DefaultTimout.DatabaseTimeout)
	defer cancel()

	ui := mux.Vars(r)

	user.logger.InfoContext(ctx, "search user by ID request started",
		"method", r.Method,
		"path", r.URL.Path,
		"user_id", ui["userID"],
	)

	id, err := uuid.Parse(ui["userID"])
	if err != nil {
		user.logger.ErrorContext(ctx, "invalid user ID format",
			"method", r.Method,
			"path", r.URL.Path,
			"error", err,
		)

		response.Err(w, http.StatusBadRequest, err)
		return
	}

	userByID, err := user.repo.SearchByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNoRows) {
			user.logger.ErrorContext(ctx, "user not found",
				"method", r.Method,
				"path", r.URL.Path,
				"user_id", ui["userID"],
			)

			response.JSON(w, http.StatusOK, "there no users found")
			return
		}

		user.logger.ErrorContext(ctx, "failed to search user by ID",
			"method", r.Method,
			"path", r.URL.Path,
			"error", err,
		)

		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	user.logger.InfoContext(ctx, "user found",
		"method", r.Method,
		"path", r.URL.Path,
		"user_id", ui["userID"],
	)

	response.JSON(w, http.StatusOK, userByID)
}

// Update modifies an existing user
func (user *UserController) Update(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), DefaultTimout.DatabaseTimeout)
	defer cancel()

	ui := mux.Vars(r)

	id, err := uuid.Parse(ui["userID"])
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var u models.User
	if err := json.Unmarshal(requestBody, &u); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := u.Validate(models.ValidationUpdate); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if _, err := user.repo.SearchByID(ctx, id); err != nil {
		if strings.Contains("no rows in result set", err.Error()) {
			response.JSON(w, http.StatusNotFound, "there no users found")
			return
		}
	}

	if err = user.repo.Update(ctx, id, u); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete removes a user
func (user *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	ui := mux.Vars(r)

	id, err := uuid.Parse(ui["userID"])
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), DefaultTimout.DatabaseTimeout)
	defer cancel()

	if err := user.repo.Delete(ctx, id); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
