package v1

import (
	"strconv"

	"go-blog-api/internal/dto"
	"go-blog-api/internal/repository"
	"go-blog-api/internal/service"
	"go-blog-api/pkg/util"

	"github.com/gin-gonic/gin"
)

// UserController 负责处理用户相关的 HTTP 请求
type UserController struct {
	userService *service.UserService
}

// NewUserController 构造函数，目前内部直接创建 UserService
// 后续可以通过依赖注入将 userService 作为参数传入
func NewUserController() *UserController {
	repo := repository.NewUserRepository()
	service := service.NewUserService(repo)
	return &UserController{
		userService: service,
	}
}

// Login 用户登录接口
// @Summary      用户登录
// @Description  使用用户名和密码登录，返回 JWT Token
// @Tags         认证
// @Accept       json
// @Produce      json
// @Param        request  body      dto.LoginRequest  true  "登录信息"
// @Success      200      {object}  util.Response{data=dto.LoginResponse}
// @Failure      400      {object}  util.Response  "参数错误"
// @Failure      401      {object}  util.Response  "用户名或密码错误"
// @Router       /auth/login [post]
func (ctrl *UserController) Login(c *gin.Context) {
	var req dto.LoginRequest
	// 参数绑定与校验
	if err := c.ShouldBindJSON(&req); err != nil {
		util.HandleError(c, util.ErrInvalidParam.WithMsg(err.Error()))
		return
	}

	// 调用业务逻辑
	resp, err := ctrl.userService.Login(&req)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.Success(c, resp)
}

// Register 用户注册接口
// @Summary      用户注册
// @Description  注册新用户账号
// @Tags         认证
// @Accept       json
// @Produce      json
// @Param        request  body      dto.RegisterRequest  true  "注册信息"
// @Success      200      {object}  util.Response  "注册成功"
// @Failure      400      {object}  util.Response  "参数错误"
// @Failure      409      {object}  util.Response  "用户名或邮箱已存在"
// @Router       /auth/register [post]
func (ctrl *UserController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	// 参数绑定与校验
	if err := c.ShouldBindJSON(&req); err != nil {
		util.HandleError(c, util.ErrInvalidParam.WithMsg(err.Error()))
		return
	}

	// 调用业务逻辑
	if err := ctrl.userService.Register(req); err != nil {
		util.HandleError(c, err)
		return
	}

	util.Success(c, nil)
}

// GetMe 获取当前登录用户信息
// @Summary      获取当前用户信息
// @Description  获取当前登录用户的详细信息
// @Tags         认证
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  util.Response{data=model.User}
// @Failure      401  {object}  util.Response  "未授权"
// @Failure      404  {object}  util.Response  "用户不存在"
// @Router       /auth/me [get]
func (ctrl *UserController) GetMe(c *gin.Context) {
	// 从 JWT 中间件获取用户 ID
	userID, exists := c.Get("userID")
	if !exists {
		util.HandleError(c, util.ErrUnauthorized)
		return
	}

	user, err := ctrl.userService.GetByID(userID.(uint))
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.Success(c, user)
}

// GetUser 获取用户详情
// @Summary      获取用户详情
// @Description  根据用户 ID 获取用户信息
// @Tags         用户
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "用户 ID"
// @Success      200  {object}  util.Response{data=model.User}
// @Failure      400  {object}  util.Response  "参数错误"
// @Failure      401  {object}  util.Response  "未授权"
// @Failure      404  {object}  util.Response  "用户不存在"
// @Router       /users/{id} [get]
func (ctrl *UserController) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.HandleError(c, util.ErrInvalidParam.WithMsg("无效的用户 ID"))
		return
	}

	user, err := ctrl.userService.GetByID(uint(id))
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.Success(c, user)
}

// ListUsers 获取用户列表
// @Summary      获取用户列表
// @Description  分页获取用户列表，支持关键词搜索
// @Tags         用户
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      dto.ListUsersRequest  true  "分页参数"
// @Success      200      {object}  util.Response{data=dto.UserPageResponse}
// @Failure      400      {object}  util.Response  "参数错误"
// @Failure      401      {object}  util.Response  "未授权"
// @Router       /users/list [post]
func (ctrl *UserController) ListUsers(c *gin.Context) {
	var req dto.ListUsersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.HandleError(c, util.ErrInvalidParam.WithMsg(err.Error()))
		return
	}

	resp, err := ctrl.userService.List(&req)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.Success(c, resp)
}

// UpdateUser 更新用户信息
// @Summary      更新用户信息
// @Description  更新当前登录用户的信息
// @Tags         用户
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      int                    true  "用户 ID"
// @Param        request  body      dto.UpdateUserRequest  true  "更新内容"
// @Success      200      {object}  util.Response{data=model.User}
// @Failure      400      {object}  util.Response  "参数错误"
// @Failure      401      {object}  util.Response  "未授权"
// @Failure      403      {object}  util.Response  "无权限"
// @Failure      404      {object}  util.Response  "用户不存在"
// @Router       /users/{id} [put]
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	// 获取当前登录用户
	currentUserID, exists := c.Get("userID")
	if !exists {
		util.HandleError(c, util.ErrUnauthorized)
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.HandleError(c, util.ErrInvalidParam.WithMsg("无效的用户 ID"))
		return
	}

	// 只能修改自己的信息
	if uint(id) != currentUserID.(uint) {
		util.HandleError(c, util.ErrForbidden.WithMsg("只能修改自己的信息"))
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.HandleError(c, util.ErrInvalidParam.WithMsg(err.Error()))
		return
	}

	user, err := ctrl.userService.Update(uint(id), &req)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.Success(c, user)
}

// DeleteUser 删除用户
// @Summary      删除用户
// @Description  删除指定用户（仅管理员或本人）
// @Tags         用户
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "用户 ID"
// @Success      200  {object}  util.Response  "删除成功"
// @Failure      400  {object}  util.Response  "参数错误"
// @Failure      401  {object}  util.Response  "未授权"
// @Failure      403  {object}  util.Response  "无权限"
// @Failure      404  {object}  util.Response  "用户不存在"
// @Router       /users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	// 获取当前登录用户
	currentUserID, exists := c.Get("userID")
	if !exists {
		util.HandleError(c, util.ErrUnauthorized)
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.HandleError(c, util.ErrInvalidParam.WithMsg("无效的用户 ID"))
		return
	}

	// 只能删除自己（后续可扩展管理员权限）
	if uint(id) != currentUserID.(uint) {
		util.HandleError(c, util.ErrForbidden.WithMsg("只能删除自己的账号"))
		return
	}

	if err := ctrl.userService.Delete(uint(id)); err != nil {
		util.HandleError(c, err)
		return
	}

	util.Success(c, nil)
}
