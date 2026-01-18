package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Err write a error message in JSON with status code and error message.
func Err(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Err string `json:"error"`
	}{
		Err: err.Error(),
	})
}

// JSON set headers and write the status code, encode the data response.
func JSON(w http.ResponseWriter, statusCode int, dataResponse any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(dataResponse); err != nil {
		log.Printf("error to encode response: %s", err)
	}
}
