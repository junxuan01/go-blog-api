package main

import (
	"fmt"
	"net/http"

	"go-blog-api/internal/model"
	"go-blog-api/internal/router"
	"go-blog-api/pkg/config"
	"go-blog-api/pkg/db"

	_ "go-blog-api/docs" // Swagger docs
)

// @title           Go Blog API
// @version         1.0
// @description     一个学习 Go Web 开发的博客后端 API 项目
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://github.com/guoxiangwen/go-blog-api
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description 输入 Bearer {token} 格式

func main() {
	// 1. 初始化配置
	config.InitConfig()

	// 2. 初始化数据库连接
	db.InitDB()

	// 3. 自动迁移数据表（等创建Model后再启用）
	db.AutoMigrate(&model.User{}, &model.Article{}, &model.Comment{})

	// 4. 初始化 Gin 路由
	r := router.InitRouter()

	// 5. 启动服务
	addr := ":" + config.AppConfig.Server.Port
	fmt.Printf("Server starting on %s\n", addr)

	srv := &http.Server{
		Addr:    addr,
		Handler: r, // 把 Gin Engine 作为 HTTP Handler 传入
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
