package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nachogoca/golang-example-rest-api-layout/internal/entities"
	"github.com/sirupsen/logrus"
)

const maxContentLen = 1000

//go:generate mockgen -destination=./mocks/articles_mock.go -package=mocks github.com/nachogoca/golang-example-rest-api-layout/internal/usecases ArticlesStore

// ArticlesStore describes all the functions we need from store layer
// create other interfaces as usecases needed
type ArticlesStore interface {
	GetAll(ctx context.Context) ([]entities.Article, error)
	GetOne(ctx context.Context, id string) (entities.Article, error)
	Create(ctx context.Context, article entities.Article) (entities.Article, error)
	Update(ctx context.Context, article entities.Article) (entities.Article, error)
	Delete(ctx context.Context, id string) error
}

// Articles is the usecase that has all the business logic about articles
type Articles struct {
	store ArticlesStore
}

// NewArticles is the Articles constructor
func NewArticles(as ArticlesStore) Articles {
	return Articles{store: as}
}

// GetAll returns all articles
func (a Articles) GetAll(ctx context.Context) ([]entities.Article, error) {

	return nil, fmt.Errorf("not implemented yet")

}

// GetOne returns one article given an id
func (a Articles) GetOne(ctx context.Context, id string) (entities.Article, error) {
	return entities.Article{}, fmt.Errorf("not implemented yet")

}

// Create creates an article
func (a Articles) Create(ctx context.Context, article entities.Article) (entities.Article, error) {

	// example of business logic applied, which should only live in the usecase layer
	if len(article.Content) > maxContentLen {
		return entities.Article{}, fmt.Errorf("article content is longer than allowed")
	}

	id := uuid.New().String()
	art := entities.Article{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     article.Title,
		Content:   article.Content,
		Author:    article.Author,
	}

	created, err := a.store.Create(ctx, art)
	if err != nil {
		return entities.Article{}, fmt.Errorf("could not create store: %w", err)
	}
	logrus.WithField("article", created).Debug("created entity")
	return created, nil
}

// Update updates the attributes of an article
func (a Articles) Update(ctx context.Context, article entities.Article) (entities.Article, error) {
	return entities.Article{}, fmt.Errorf("not implemented yet")

}

// Delete removes an article
func (a Articles) Delete(ctx context.Context, id string) error {
	return fmt.Errorf("not implemented yet")
}
