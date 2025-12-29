package main

import (
	"fmt"
	"net/http"

	"go-blog-api/internal/model"
	"go-blog-api/internal/router"
	"go-blog-api/pkg/config"
	"go-blog-api/pkg/db"
)

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
