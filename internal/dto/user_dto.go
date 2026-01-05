package dto

import "go-blog-api/internal/model"

// ========== 请求结构 ==========

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Email  string `json:"email" binding:"omitempty,email"`
	Avatar string `json:"avatar" binding:"omitempty,url"`
}

// ListUsersRequest 用户列表请求（嵌入通用分页）
type ListUsersRequest struct {
	PageRequest
	Keyword string `json:"keyword" binding:"omitempty"` // 搜索关键词（用户名/邮箱）
}

// ========== 响应结构 ==========

// LoginResponse 登录响应
type LoginResponse struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}
