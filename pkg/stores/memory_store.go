package stores

import (
	"chipoc/pkg/models"
	"errors"
	"fmt"
	"math/rand"
)

var articles = []*models.Article{
	{ID: "1", UserID: 100, Title: "Hi", Slug: "hi"},
	{ID: "2", UserID: 200, Title: "sup", Slug: "sup"},
	{ID: "3", UserID: 300, Title: "alo", Slug: "alo"},
	{ID: "4", UserID: 400, Title: "bonjour", Slug: "bonjour"},
	{ID: "5", UserID: 500, Title: "whats up", Slug: "whats-up"},
}

// User fixture data
var users = []*models.User{
	{ID: 100, Name: "Peter"},
	{ID: 200, Name: "Julia"},
}

type memoryStore struct {
	ArticlesStore
}

func NewMemoryStore() ArticlesStore {
	return &memoryStore{}
}

func (s *memoryStore) ListArticles() ([]*models.Article, error) {
	return articles, nil
}

func (s *memoryStore) NewArticle(article *models.Article) (string, error) {
	article.ID = fmt.Sprintf("%d", rand.Intn(100)+10)
	articles = append(articles, article)
	return article.ID, nil
}

func (s *memoryStore) GetArticle(id string) (*models.Article, error) {
	for _, a := range articles {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, errors.New("article not found.")
}

func (s *memoryStore) GetArticleBySlug(slug string) (*models.Article, error) {
	for _, a := range articles {
		if a.Slug == slug {
			return a, nil
		}
	}
	return nil, errors.New("article not found.")
}


func (s *memoryStore)  UpdateArticle(id string, article *models.Article) (*models.Article, error) {
	for i, a := range articles {
		if a.ID == id {
			articles[i] = article
			return article, nil
		}
	}
	return nil, errors.New("article not found.")
}

func (s *memoryStore) RemoveArticle(id string) (*models.Article, error) {
	for i, a := range articles {
		if a.ID == id {
			articles = append((articles)[:i], (articles)[i+1:]...)
			return a, nil
		}
	}
	return nil, errors.New("article not found.")
}

func (s *memoryStore) GetUser(id int64) (*models.User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found.")
}