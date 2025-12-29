package repository

import (
	"go-blog-api/internal/model"
	"go-blog-api/pkg/db"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *model.User) error
	GetByUsername(username string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

// 确保 UserRepository 实现了接口
var _ IUserRepository = (*UserRepository)(nil)

func NewUserRepository() *UserRepository {
	return &UserRepository{db: db.DB}
}

// CreateUser 创建新用户
func (r *UserRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

// GetByUsername 根据用户名获取用户
func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail 根据邮箱获取用户
func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
