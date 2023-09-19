package services

import "github.com/Naokiiiiiii/BlogApiPractice/models"

type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
}

type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}

type Niceservicer interface {
	PostNiceSerice(nice models.Nice) (models.Nice, error)
}
