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
		return nil, util.ErrInvalidCredentials
	}

	// 2. 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, util.ErrInvalidCredentials
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

// GetByID 获取用户详情
func (s *UserService) GetByID(id uint) (*model.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, util.ErrUserNotFound
	}
	return user, nil
}

// Update 更新用户信息
func (s *UserService) Update(id uint, req *dto.UpdateUserRequest) (*model.User, error) {
	// 1. 查询用户
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, util.ErrUserNotFound
	}

	// 2. 检查邮箱是否被其他用户使用
	if req.Email != "" && req.Email != user.Email {
		if existingUser, _ := s.userRepo.GetByEmail(req.Email); existingUser != nil && existingUser.ID != id {
			return nil, util.ErrEmailExists
		}
		user.Email = req.Email
	}

	// 3. 更新头像
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}

	// 4. 保存
	if err := s.userRepo.Update(user); err != nil {
		return nil, util.ErrDatabase
	}

	return user, nil
}

// Delete 删除用户
func (s *UserService) Delete(id uint) error {
	// 1. 检查用户是否存在
	if _, err := s.userRepo.GetByID(id); err != nil {
		return util.ErrUserNotFound
	}

	// 2. 删除
	if err := s.userRepo.Delete(id); err != nil {
		return util.ErrDatabase
	}

	return nil
}

// List 获取用户列表
func (s *UserService) List(req *dto.ListUsersRequest) (*dto.PageResponse[model.User], error) {
	req.SetDefaults()

	users, total, err := s.userRepo.List(req.Offset(), req.PageSize, req.Keyword)
	if err != nil {
		return nil, util.ErrDatabase
	}

	return dto.NewPageResponse(users, total, req.Page, req.PageSize), nil
}

// Logout 注销用户（当前实现为无状态，服务端无需额外操作）
func (s *UserService) Logout(userID uint) error {
	// 如果将来需要在服务端维护 token 黑名单或会话表，可在这里实现
	return nil
}
