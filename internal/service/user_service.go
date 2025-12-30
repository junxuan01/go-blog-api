package service

import (
	"errors"
	"go-blog-api/internal/model"
	"go-blog-api/internal/repository"
	"go-blog-api/pkg/util"

	"golang.org/x/crypto/bcrypt"
)

// UserService 负责和“用户相关”的业务逻辑
type UserService struct {
	userRepo repository.IUserRepository
}
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginResponse struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

// NewUserService 构造函数，目前内部自己创建依赖
// 后面我们会讨论如何通过依赖注入把这个依赖从外部传进来
func NewUserService(userRepo repository.IUserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// Login 用户登录
func (s *UserService) Login(req *LoginRequest) (*LoginResponse, error) {
	// 1. 查询用户
	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	// 2. 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	// 3. 生成 Token
	token, err := util.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: token,
		User:  *user,
	}, nil
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
