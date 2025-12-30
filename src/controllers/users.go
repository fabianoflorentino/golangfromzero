package controllers

import "net/http"

// Create handles the creation of a new user
func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User created"))
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
