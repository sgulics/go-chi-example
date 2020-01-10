package services

import (
	"chipoc/pkg/models"
	"chipoc/pkg/stores"
	"github.com/sirupsen/logrus"
)

type ArticlesService interface {
	NewArticle(article *models.Article) (string, error)
	GetArticle(id string) (*models.Article, error)
	GetArticleBySlug(slug string) (*models.Article, error)
	UpdateArticle(id string, article *models.Article) (*models.Article, error)
	RemoveArticle(id string) (*models.Article, error)
	GetUser(id int64) (*models.User, error)
	ListArticles() ([]*models.Article, error)
}

type articlesService struct{
	store stores.ArticlesStore
	logger *logrus.Logger
}

func NewArticlesService(store stores.ArticlesStore, logger *logrus.Logger) *articlesService {
	return &articlesService{store: store, logger: logger}
}

func (s *articlesService) NewArticle(article *models.Article) (string, error) {
	return s.store.NewArticle(article)
}

func (s *articlesService) GetArticle(id string) (*models.Article, error) {
	log := s.logger.WithFields(logrus.Fields{
		"articleId": id,
	})
	log.Info("GetArticle")
	return s.store.GetArticle(id)
}

func (s *articlesService) GetArticleBySlug(slug string) (*models.Article, error) {
	return s.store.GetArticleBySlug(slug)
}

func (s *articlesService) UpdateArticle(id string, article *models.Article) (*models.Article, error) {
	return s.store.UpdateArticle(id, article)
}

func (s *articlesService) RemoveArticle(id string) (*models.Article, error) {
	return s.store.RemoveArticle(id)
}

func (s *articlesService) GetUser(id int64) (*models.User, error) {
	return s.store.GetUser(id)
}

func (s *articlesService) ListArticles() ([]*models.Article, error) {
	return s.store.ListArticles()
}
