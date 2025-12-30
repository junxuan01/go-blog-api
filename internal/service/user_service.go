package service

import (
	"go-blog-api/internal/dto"
	"go-blog-api/internal/model"
	"go-blog-api/internal/repository"
	"go-blog-api/pkg/util"

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

// Login 用户登录
func (s *UserService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// 1. 查询用户
	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, util.ErrUserNotFound.WithMsg("用户名或密码错误")
	}

	// 2. 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, util.ErrUnauthorized.WithMsg("用户名或密码错误")
	}

	// 3. 生成 Token
	token, err := util.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}

// Register 注册新用户
func (s *UserService) Register(req dto.RegisterRequest) error {
	// 1. 检查用户名是否存在
	if _, err := s.userRepo.GetByUsername(req.Username); err == nil {
		return util.ErrUsernameExists
	}
	// 2. 检查邮箱是否存在
	if _, err := s.userRepo.GetByEmail(req.Email); err == nil {
		return util.ErrEmailExists
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
