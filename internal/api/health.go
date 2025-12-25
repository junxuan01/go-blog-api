package api

import (
	"encoding/json"
	"go-blog-api/internal/service"
	"net/http"
)

// HealthHandler 简单的健康检查接口
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := service.NewHealthService().Get()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
