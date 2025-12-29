package service

import (
	"errors"
	"go-blog-api/internal/model"
	"go-blog-api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

// UserService 负责和“用户相关”的业务逻辑
type UserService struct {
	userRepo repository.IUserRepository
}

// NewUserService 构造函数，目前内部自己创建依赖
// 后面我们会讨论如何通过依赖注入把这个依赖从外部传进来
func NewUserService(userRepo repository.IUserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

// Register 注册新用户
func (s *UserService) Register(req RegisterRequest) error {

	// 1. 检查用户名是否存在
	if _, err := s.userRepo.GetByUsername(req.Username); err == nil {
		return errors.New("username already exists")
	}
	// 2. 检查邮箱是否存在
	if _, err := s.userRepo.GetByEmail(req.Email); err == nil {
		return errors.New("email already exists")
	}
	// 3. 密码加密
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// 4. 创建用户
	user := &model.User{
		Username: req.Username,
		Password: string(hashedPwd),
		Email:    req.Email,
		Avatar:   "https://example.com/default-avatar.png",
	}
	return s.userRepo.CreateUser(user)
}
