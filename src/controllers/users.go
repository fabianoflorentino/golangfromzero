package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fabianoflorentino/golangfromzero/database"
	"github.com/fabianoflorentino/golangfromzero/repository"
	"github.com/fabianoflorentino/golangfromzero/src/models"
)

// Create handles the creation of a new user
func Create(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	var u models.User
	if err := json.Unmarshal(requestBody, &u); err != nil {
		http.Error(w, "failed, invalid JSON", http.StatusBadRequest)
		return
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, "database unavailable", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	id, err := repository.Create(u)
	if err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"id":      id,
		"message": "user created",
	})
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
