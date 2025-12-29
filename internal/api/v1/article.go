package v1

import (
	"go-blog-api/pkg/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ArticleController 负责处理文章相关的 HTTP 请求
type ArticleController struct{}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

// GetArticle 获取单篇文章
// @Summary 获取文章详情
// @Tags 文章
// @Param id path int true "文章ID"
// @Success 200 {object} util.Response
// @Router /api/v1/articles/{id} [get]
func (ctrl *ArticleController) GetArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.Error(c, http.StatusBadRequest, 40001, "Invalid article ID")
		return
	}
	// 模拟从Service层获取数据
	mockArticle := map[string]any{
		"id":      id,
		"title":   "Gin 实战指南",
		"content": "Gin 是一个高性能的 Go Web 框架...",
	}
	util.Success(c, mockArticle)
}

// ListArticles 获取文章列表
// @Summary 获取文章列表
// @Tags 文章
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} util.Response
// @Router /api/v1/articles [get]
func (ctrl *ArticleController) ListArticles(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, err1 := strconv.Atoi(pageStr)
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err1 != nil || err2 != nil || page <= 0 || pageSize <= 0 {
		util.Error(c, http.StatusBadRequest, 40002, "Invalid pagination parameters")
		return
	}

	// 模拟从Service层获取数据
	mockArticles := []map[string]any{
		{"id": 1, "title": "Gin 实战指南"},
		{"id": 2, "title": "Go 并发编程"},
	}
	util.Success(c, gin.H{
		"page":      page,
		"page_size": pageSize,
		"articles":  mockArticles,
		"total":     2,
	})
}
