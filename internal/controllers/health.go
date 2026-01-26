package controllers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/fabianoflorentino/golangfromzero/internal/response"
)

type HealthController struct {
	logger *slog.Logger
}

func NewHealthController(log *slog.Logger) *HealthController {
	return &HealthController{
		logger: log,
	}
}

func (hc *HealthController) Check(w http.ResponseWriter, r *http.Request) {
	hc.logger.InfoContext(r.Context(), "health check requested",
		"method", r.Method,
		"path", r.URL.Path,
	)

	healthOK := map[string]any{
		"status":    "OK",
		"message":   "Service is healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	response.JSON(w, http.StatusOK, healthOK)
}
