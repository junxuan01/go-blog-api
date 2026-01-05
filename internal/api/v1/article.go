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
// @Summary      获取文章详情
// @Description  根据文章 ID 获取文章详情
// @Tags         文章
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "文章 ID"
// @Success      200  {object}  util.Response{data=model.Article}
// @Failure      400  {object}  util.Response  "参数错误"
// @Failure      401  {object}  util.Response  "未授权"
// @Failure      404  {object}  util.Response  "文章不存在"
// @Router       /articles/{id} [get]
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
// @Summary      获取文章列表
// @Description  分页获取文章列表
// @Tags         文章
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      dto.ListArticlesRequest  true  "分页参数"
// @Success      200      {object}  util.Response{data=dto.ArticlePageResponse}
// @Failure      400      {object}  util.Response  "参数错误"
// @Failure      401      {object}  util.Response  "未授权"
// @Router       /articles/list [post]
func (ctrl *ArticleController) ListArticles(c *gin.Context) {
	var req dto.ListArticlesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
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
// @Summary      创建文章
// @Description  创建一篇新文章，需要登录
// @Tags         文章
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      dto.CreateArticleRequest  true  "文章内容"
// @Success      200      {object}  util.Response{data=model.Article}
// @Failure      400      {object}  util.Response  "参数错误"
// @Failure      401      {object}  util.Response  "未授权"
// @Router       /articles [post]
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
// @Summary      更新文章
// @Description  更新指定文章，只能更新自己的文章
// @Tags         文章
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      int                       true  "文章 ID"
// @Param        request  body      dto.UpdateArticleRequest  true  "更新内容"
// @Success      200      {object}  util.Response{data=model.Article}
// @Failure      400      {object}  util.Response  "参数错误"
// @Failure      401      {object}  util.Response  "未授权"
// @Failure      403      {object}  util.Response  "无权限"
// @Failure      404      {object}  util.Response  "文章不存在"
// @Router       /articles/{id} [put]
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
// @Summary      删除文章
// @Description  删除指定文章，只能删除自己的文章
// @Tags         文章
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "文章 ID"
// @Success      200  {object}  util.Response  "删除成功"
// @Failure      400  {object}  util.Response  "参数错误"
// @Failure      401  {object}  util.Response  "未授权"
// @Failure      403  {object}  util.Response  "无权限"
// @Failure      404  {object}  util.Response  "文章不存在"
// @Router       /articles/{id} [delete]
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
