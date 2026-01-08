package router

import (
	v1 "go-blog-api/internal/api/v1"
	"go-blog-api/internal/middleware"
	"go-blog-api/pkg/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	// 设置运行模式：debug / release，从配置读取
	gin.SetMode(config.AppConfig.Server.Mode)
	// 推荐使用 gin.New() 而不是 gin.Default()，便于精细控制中间件
	r := gin.New()
	r.Use(gin.Logger())   // 日志中间件
	r.Use(gin.Recovery()) // 恢复中间件，防止崩溃
	// CORS 中间件，允许本地前端开发访问
	r.Use(middleware.CORS())

	// Swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 初始化 Controller
	articleCtrl := v1.NewArticleController()
	userCtrl := v1.NewUserController()
	// 路由分组：/api/v1 作为统一前缀，方便做版本控制
	apiV1 := r.Group("/api/v1")
	{

		// /api/v1/auth auth相关
		auth := apiV1.Group("/auth")
		{
			auth.POST("/login", userCtrl.Login)
			auth.POST("/register", userCtrl.Register)
		}

		// /api/v1/articles 相关接口
		articles := apiV1.Group("/articles")
		articles.Use(middleware.JWT()) // 挂载中间件
		{
			articles.POST("/list", articleCtrl.ListArticles)
			articles.GET(":id", articleCtrl.GetArticle)
			articles.POST("", articleCtrl.CreateArticle)
			articles.PUT(":id", articleCtrl.UpdateArticle)
			articles.DELETE(":id", articleCtrl.DeleteArticle)
		}

		// /api/v1/users 用户管理接口
		users := apiV1.Group("/users")
		users.Use(middleware.JWT()) // 需要登录
		{
			users.POST("/list", userCtrl.ListUsers)
			users.GET(":id", userCtrl.GetUser)
			users.PUT(":id", userCtrl.UpdateUser)
			users.DELETE(":id", userCtrl.DeleteUser)
		}
	}
	return r
}
