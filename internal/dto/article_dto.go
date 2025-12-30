package dto

import "go-blog-api/internal/model"

// ========== 请求结构 ==========

// CreateArticleRequest 创建文章请求
type CreateArticleRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=255"`
	Content string `json:"content" binding:"required"`
}

// UpdateArticleRequest 更新文章请求
type UpdateArticleRequest struct {
	Title   string `json:"title" binding:"omitempty,min=1,max=255"`
	Content string `json:"content" binding:"omitempty"`
}

// ListArticlesRequest 文章列表请求
type ListArticlesRequest struct {
	Page     int `form:"page" binding:"omitempty,min=1"`
	PageSize int `form:"page_size" binding:"omitempty,min=1,max=100"`
}

// ========== 响应结构 ==========

// ArticleResponse 文章响应
type ArticleResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

// ArticleListResponse 文章列表响应
type ArticleListResponse struct {
	Articles []model.Article `json:"articles"`
	Total    int64           `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
}
