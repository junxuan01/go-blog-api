package dto

import "go-blog-api/internal/model"

// PageRequest 通用分页请求
type PageRequest struct {
	Page     int `json:"page" binding:"omitempty,min=1"`
	PageSize int `json:"page_size" binding:"omitempty,min=1,max=100"`
}

// SetDefaults 设置默认分页参数
func (p *PageRequest) SetDefaults() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
}

// Offset 计算数据库查询偏移量
func (p *PageRequest) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// PageResponse 通用分页响应（泛型）
type PageResponse[T any] struct {
	List     []T   `json:"list"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
}

// NewPageResponse 创建分页响应
func NewPageResponse[T any](list []T, total int64, page, pageSize int) *PageResponse[T] {
	return &PageResponse[T]{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
}

// ========== Swagger 文档用的具体类型别名 ==========
// 由于 swaggo 不完全支持泛型，这里定义具体类型用于文档生成

// ArticlePageResponse 文章分页响应（Swagger 用）
type ArticlePageResponse = PageResponse[model.Article]

// UserPageResponse 用户分页响应（Swagger 用）
type UserPageResponse = PageResponse[model.User]

// CommentPageResponse 评论分页响应（Swagger 用）
type CommentPageResponse = PageResponse[model.Comment]
