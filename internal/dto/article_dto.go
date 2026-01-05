package dto

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

// ListArticlesRequest 文章列表请求（嵌入通用分页）
type ListArticlesRequest struct {
	PageRequest
}
