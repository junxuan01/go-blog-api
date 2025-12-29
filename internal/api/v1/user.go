package v1

import (
	"go-blog-api/internal/repository"
	"go-blog-api/internal/service"
	"go-blog-api/pkg/util"
	"net/http"

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

// Register 用户注册接口
func (ctrl *UserController) Register(c *gin.Context) {
	var req service.RegisterRequest
	// 参数绑定与校验
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, http.StatusBadRequest, 40001, err.Error())
		return
	}

	// 调用业务逻辑
	if err := ctrl.userService.Register(req); err != nil {
		util.Error(c, http.StatusInternalServerError, 50001, err.Error())
		return
	}

	util.Success(c, nil)
}
