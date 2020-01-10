package stores

import (
	"chipoc/pkg/models"
)

type ArticlesStore interface {
	NewArticle(article *models.Article) (string, error)
	GetArticle(id string) (*models.Article, error)
	GetArticleBySlug(slug string) (*models.Article, error)
	UpdateArticle(id string, article *models.Article) (*models.Article, error)
	RemoveArticle(id string) (*models.Article, error)
	GetUser(id int64) (*models.User, error)
	ListArticles() ([]*models.Article, error)
}



//func dbNewArticle(article *models.Article) (string, error) {
//	article.ID = fmt.Sprintf("%d", rand.Intn(100)+10)
//	articles = append(articles, article)
//	return article.ID, nil
//}

//func dbGetArticle(id string) (*models.Article, error) {
//	for _, a := range articles {
//		if a.ID == id {
//			return a, nil
//		}
//	}
//	return nil, errors.New("article not found.")
//}

//func dbGetArticleBySlug(slug string) (*models.Article, error) {
//	for _, a := range articles {
//		if a.Slug == slug {
//			return a, nil
//		}
//	}
//	return nil, errors.New("article not found.")
//}

//func dbUpdateArticle(id string, article *models.Article) (*models.Article, error) {
//	for i, a := range articles {
//		if a.ID == id {
//			articles[i] = article
//			return article, nil
//		}
//	}
//	return nil, errors.New("article not found.")
//}

//func dbRemoveArticle(id string) (*models.Article, error) {
//	for i, a := range articles {
//		if a.ID == id {
//			articles = append((articles)[:i], (articles)[i+1:]...)
//			return a, nil
//		}
//	}
//	return nil, errors.New("article not found.")
//}
//
//func dbGetUser(id int64) (*User, error) {
//	for _, u := range users {
//		if u.ID == id {
//			return u, nil
//		}
//	}
//	return nil, errors.New("user not found.")
//}
