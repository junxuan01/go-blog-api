# go-blog-api

一个简单的 Go Web API 项目，采用三层架构（API/Service/Repository），用于学习 Go Web 开发。

## 项目结构

```
go-blog-api/
├── cmd/
│   └── server/
│       └── main.go        # 程序入口，只负责启动
├── internal/              # 私有应用代码，外部无法导入
│   ├── api/               # API 接口层 (Controller)
│   ├── service/           # 业务逻辑层 (Service)
│   ├── repository/        # 数据访问层 (Repository/DAO)
│   ├── model/             # 数据库模型 (GORM Model)
│   └── router/            # 路由配置
├── pkg/                   # 公共库代码，可以被外部项目引用
│   └── util/              # 工具类
├── configs/               # 配置文件
│   └── config.yaml
├── go.mod
└── go.sum
```

## 快速开始

### 运行项目

```bash
go run cmd/server/main.go
```

### 测试健康检查接口

```bash
curl http://localhost:8080/health
```

你应该会看到类似以下的返回：

```json
{"status":"ok","time":"2025-12-17T22:04:00+08:00"}
```

## 开发计划

- [x] 项目结构初始化
- [ ] 添加数据库连接（GORM）
- [ ] 实现用户认证
- [ ] 实现博客文章 CRUD
- [ ] 添加中间件（日志、认证等）
- [ ] 添加单元测试

## 技术栈

- Go 1.25+
- net/http（标准库路由）
- 后续：GORM（ORM）、Gin（可选框架）等

