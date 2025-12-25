package service

import "time"

// HealthResponse 表示健康检查的返回信息
type HealthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

type HealthService struct{}

func NewHealthService() *HealthService { return &HealthService{} }

func (s *HealthService) Get() HealthResponse {
	return HealthResponse{Status: "ok", Time: time.Now().Format(time.RFC3339)}
}
