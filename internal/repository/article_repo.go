package repository

import (
	"go-blog-api/internal/model"
	"go-blog-api/pkg/db"

	"gorm.io/gorm"
)

type IArticleRepository interface {
	Create(article *model.Article) error
	GetByID(id uint) (*model.Article, error)
	Update(article *model.Article) error
	Delete(id uint) error
	List(offset, limit int) ([]model.Article, int64, error)
	ListByUserID(userID uint, offset, limit int) ([]model.Article, int64, error)
}

type ArticleRepository struct {
	db *gorm.DB
}

// 确保 ArticleRepository 实现了接口
var _ IArticleRepository = (*ArticleRepository)(nil)

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{db: db.DB}
}

// Create 创建文章
func (r *ArticleRepository) Create(article *model.Article) error {
	return r.db.Create(article).Error
}

// GetByID 根据 ID 获取文章
func (r *ArticleRepository) GetByID(id uint) (*model.Article, error) {
	var article model.Article
	if err := r.db.Preload("User").First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

// Update 更新文章
func (r *ArticleRepository) Update(article *model.Article) error {
	return r.db.Save(article).Error
}

// Delete 删除文章（软删除）
func (r *ArticleRepository) Delete(id uint) error {
	return r.db.Delete(&model.Article{}, id).Error
}

// List 获取文章列表
func (r *ArticleRepository) List(offset, limit int) ([]model.Article, int64, error) {
	var articles []model.Article
	var total int64

	if err := r.db.Model(&model.Article{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Preload User 信息
	if err := r.db.Preload("User").Offset(offset).Limit(limit).Order("created_at DESC").Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

// ListByUserID 根据用户 ID 获取文章列表
func (r *ArticleRepository) ListByUserID(userID uint, offset, limit int) ([]model.Article, int64, error) {
	var articles []model.Article
	var total int64

	query := r.db.Model(&model.Article{}).Where("user_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Preload User 信息
	if err := query.Preload("User").Offset(offset).Limit(limit).Order("created_at DESC").Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}
