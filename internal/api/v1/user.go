package v1

import (
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
