package v1

import (
	"strconv"

	"go-blog-api/internal/dto"
	"go-blog-api/internal/repository"
	"go-blog-api/internal/service"
	"go-blog-api/pkg/util"

	"github.com/gin-gonic/gin"
)

// ArticleController 负责处理文章相关的 HTTP 请求
type ArticleController struct {
	articleService *service.ArticleService
}

func NewArticleController() *ArticleController {
	repo := repository.NewArticleRepository()
	svc := service.NewArticleService(repo)
	return &ArticleController{articleService: svc}
}

// GetArticle 获取单篇文章
func (ctrl *ArticleController) GetArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.HandleError(c, util.ErrInvalidParam.WithMsg("无效的文章 ID"))
		return
	}

	article, err := ctrl.articleService.GetByID(uint(id))
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.Success(c, article)
}

// ListArticles 获取文章列表
func (ctrl *ArticleController) ListArticles(c *gin.Context) {
	var req dto.ListArticlesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		util.HandleError(c, util.ErrInvalidParam.WithMsg(err.Error()))
		return
	}

	resp, err := ctrl.articleService.List(&req)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.Success(c, resp)
}

// CreateArticle 创建新文章
func (ctrl *ArticleController) CreateArticle(c *gin.Context) {
	// 从 Context 获取当前用户 ID
	userID, exists := c.Get("userID")
	if !exists {
		util.HandleError(c, util.ErrUnauthorized)
		return
	}

	var req dto.CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.HandleError(c, util.ErrInvalidParam.WithMsg(err.Error()))
		return
	}

	article, err := ctrl.articleService.Create(userID.(uint), &req)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.Success(c, article)
}

// UpdateArticle 更新文章
func (ctrl *ArticleController) UpdateArticle(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		util.HandleError(c, util.ErrUnauthorized)
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.HandleError(c, util.ErrInvalidParam.WithMsg("无效的文章 ID"))
		return
	}

	var req dto.UpdateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.HandleError(c, util.ErrInvalidParam.WithMsg(err.Error()))
		return
	}

	article, err := ctrl.articleService.Update(uint(id), userID.(uint), &req)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.Success(c, article)
}

// DeleteArticle 删除文章
func (ctrl *ArticleController) DeleteArticle(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		util.HandleError(c, util.ErrUnauthorized)
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.HandleError(c, util.ErrInvalidParam.WithMsg("无效的文章 ID"))
		return
	}

	if err := ctrl.articleService.Delete(uint(id), userID.(uint)); err != nil {
		util.HandleError(c, err)
		return
	}

	util.Success(c, nil)
}
