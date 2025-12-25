package main

import (
	"fmt"
	"net/http"

	"go-blog-api/internal/router"
	"go-blog-api/pkg/config"
)

func main() {
	// 1. 初始化配置
	config.InitConfig()

	// 2. 初始化 Gin 路由
	r := router.InitRouter()

	// 3. 启动服务
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
