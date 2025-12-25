package router

import (
	"go-blog-api/internal/api"
	"net/http"
)

// SetupRouter returns the http handler with routes registered.
func SetupRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", api.HealthHandler)
	return mux
}
