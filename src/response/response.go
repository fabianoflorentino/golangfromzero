package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func Err(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Err string `json:"error"`
	}{
		Err: err.Error(),
	})
}

func JSON(w http.ResponseWriter, statusCode int, dataResponse any) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(dataResponse); err != nil {
		log.Printf("error to encode response: %s", err)
	}
}
