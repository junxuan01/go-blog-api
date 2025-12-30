package service

import (
	"go-blog-api/internal/dto"
	"go-blog-api/internal/model"
	"go-blog-api/internal/repository"
	"go-blog-api/pkg/util"
)

type ArticleService struct {
	articleRepo repository.IArticleRepository
}

func NewArticleService(repo repository.IArticleRepository) *ArticleService {
	return &ArticleService{articleRepo: repo}
}

// Create 创建文章
func (s *ArticleService) Create(userID uint, req *dto.CreateArticleRequest) (*model.Article, error) {
	article := &model.Article{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := s.articleRepo.Create(article); err != nil {
		return nil, util.ErrDatabase
	}

	return article, nil
}

// GetByID 获取文章详情
func (s *ArticleService) GetByID(id uint) (*model.Article, error) {
	article, err := s.articleRepo.GetByID(id)
	if err != nil {
		return nil, util.ErrArticleNotFound
	}
	return article, nil
}

// Update 更新文章
func (s *ArticleService) Update(id, userID uint, req *dto.UpdateArticleRequest) (*model.Article, error) {
	// 1. 查询文章
	article, err := s.articleRepo.GetByID(id)
	if err != nil {
		return nil, util.ErrArticleNotFound
	}

	// 2. 检查权限
	if article.UserID != userID {
		return nil, util.ErrForbidden
	}

	// 3. 更新字段
	if req.Title != "" {
		article.Title = req.Title
	}
	if req.Content != "" {
		article.Content = req.Content
	}

	if err := s.articleRepo.Update(article); err != nil {
		return nil, util.ErrDatabase
	}

	return article, nil
}

// Delete 删除文章
func (s *ArticleService) Delete(id, userID uint) error {
	// 1. 查询文章
	article, err := s.articleRepo.GetByID(id)
	if err != nil {
		return util.ErrArticleNotFound
	}

	// 2. 检查权限
	if article.UserID != userID {
		return util.ErrForbidden
	}

	// 3. 删除
	if err := s.articleRepo.Delete(id); err != nil {
		return util.ErrDatabase
	}

	return nil
}

// List 获取文章列表
func (s *ArticleService) List(req *dto.ListArticlesRequest) (*dto.ArticleListResponse, error) {
	// 默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	offset := (req.Page - 1) * req.PageSize

	articles, total, err := s.articleRepo.List(offset, req.PageSize)
	if err != nil {
		return nil, util.ErrDatabase
	}

	return &dto.ArticleListResponse{
		Articles: articles,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
