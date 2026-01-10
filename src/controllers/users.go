package controllers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/fabianoflorentino/golangfromzero/database"
	"github.com/fabianoflorentino/golangfromzero/repository"
	"github.com/fabianoflorentino/golangfromzero/src/helper"
	"github.com/fabianoflorentino/golangfromzero/src/models"
	"github.com/fabianoflorentino/golangfromzero/src/response"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type UserController struct {
	cfg helper.Config
}

func NewUserController(cfg helper.Config) *UserController {
	return &UserController{cfg: cfg}
}

// Create handles the creation of a new user
func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Validate("new"); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(r.Context(), helper.ConfigTimeout.DatabaseTimeout)
	defer cancel()

	repository := repository.NewUsersRepository(db)
	user.ID, err = repository.Create(ctx, user)
	if err != nil {

		if strings.Contains(err.Error(), "email already used") {
			response.Err(w, http.StatusBadRequest, err)
			return
		}

		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

// SearchByName retrieves all users with contains a search name
func (u *UserController) SearchByName(w http.ResponseWriter, r *http.Request) {
	nameToSearch := strings.ToLower(r.URL.Query().Get("name"))
	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(r.Context(), helper.ConfigTimeout.DatabaseTimeout)
	defer cancel()

	repository := repository.NewUsersRepository(db)
	users, err := repository.SearchByName(ctx, nameToSearch)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	if users == nil {
		response.JSON(w, http.StatusOK, "there no users found")
		return
	}

	response.JSON(w, http.StatusOK, users)
}

// SearchByID retrieves a user by their ID
func (u *UserController) SearchByID(w http.ResponseWriter, r *http.Request) {
	ui := mux.Vars(r)

	id, err := uuid.Parse(ui["userID"])
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(r.Context(), helper.ConfigTimeout.DatabaseTimeout)
	defer cancel()

	repository := repository.NewUsersRepository(db)
	userByID, err := repository.SearchByID(ctx, id)
	if err != nil {
		if strings.Contains("no rows in result set", err.Error()) {
			response.JSON(w, http.StatusOK, "there no users found")
			return
		}

		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, userByID)
}

// Update modifies an existing user
func (u *UserController) Update(w http.ResponseWriter, r *http.Request) {
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

	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Validate("update"); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(r.Context(), helper.ConfigTimeout.DatabaseTimeout)
	defer cancel()

	repository := repository.NewUsersRepository(db)

	if _, err := repository.SearchByID(ctx, id); err != nil {
		if strings.Contains("no rows in result set", err.Error()) {
			response.JSON(w, http.StatusNotFound, "there no users found")
			return
		}
	}

	if err = repository.Update(ctx, id, user); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, "")
}

// Delete removes a user
func (u *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	ui := mux.Vars(r)

	id, err := uuid.Parse(ui["userID"])
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(r.Context(), helper.ConfigTimeout.DatabaseTimeout)
	defer cancel()

	repository := repository.NewUsersRepository(db)
	if err := repository.Delete(ctx, id); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, "")
}
