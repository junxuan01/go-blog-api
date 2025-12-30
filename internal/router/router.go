package router

import (
	v1 "go-blog-api/internal/api/v1"
	"go-blog-api/internal/middleware"
	"go-blog-api/pkg/config"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 设置运行模式：debug / release，从配置读取
	gin.SetMode(config.AppConfig.Server.Mode)
	// 推荐使用 gin.New() 而不是 gin.Default()，便于精细控制中间件
	r := gin.New()
	r.Use(gin.Logger())   // 日志中间件
	r.Use(gin.Recovery()) // 恢复中间件，防止崩溃

	// 初始化 Controller
	articleCtrl := v1.NewArticleController()
	userCtrl := v1.NewUserController()
	// 路由分组：/api/v1 作为统一前缀，方便做版本控制
	apiV1 := r.Group("/api/v1")
	{

		// /api/v1/auth 用户相关
		auth := apiV1.Group("/auth")
		{
			auth.POST("/login", userCtrl.Login)
			auth.POST("/register", userCtrl.Register)
		}

		// /api/v1/articles 相关接口
		articles := apiV1.Group("/articles")
		articles.Use(middleware.JWT()) // 挂载中间件
		{
			articles.GET("", articleCtrl.ListArticles)
			articles.GET(":id", articleCtrl.GetArticle)
			articles.POST("", articleCtrl.CreateArticle)
			articles.PUT(":id", articleCtrl.UpdateArticle)
			articles.DELETE(":id", articleCtrl.DeleteArticle)
		}
	}
	return r
}
