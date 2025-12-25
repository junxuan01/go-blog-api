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
	mockArticle := map[string]interface{}{
		"id":      id,
		"title":   "Gin 实战指南",
		"content": "Gin 是一个高性能的 Go Web 框架...",
	}
	util.Success(c, mockArticle)
}
