package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fabianoflorentino/golangfromzero/database"
	"github.com/fabianoflorentino/golangfromzero/repository"
	"github.com/fabianoflorentino/golangfromzero/src/models"
	"github.com/fabianoflorentino/golangfromzero/src/response"
)

// Create handles the creation of a new user
func Create(w http.ResponseWriter, r *http.Request) {
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

	if err := user.Validate(); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

// GetAll retrieves all users
func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All users"))
}

// GetByID retrieves a user by their ID
func GetByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User retrieved"))
}

// Update modifies an existing user
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User updated"))
}

// Delete removes a user
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User deleted"))
}
