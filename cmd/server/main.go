package main

import (
	"fmt"
	"go-blog-api/pkg/config"
	"go-blog-api/pkg/util"
	"net/http"
	"strconv"
)

func main() {
	// 1. 初始化配置
	config.InitConfig()
	fmt.Println("Configuration loaded successfully")

	// 2. 启动 Web 服务 (暂时用 net/http 演示)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		ret := util.AddCount(1, 3) // 示例调用 util 包中的函数
		w.Write([]byte("pong"))
		w.Write([]byte(strconv.Itoa(ret)))
	})

	addr := ":" + config.AppConfig.Server.Port
	fmt.Printf("Server is running on http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
