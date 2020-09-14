package usecases

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/nachogoca/golang-example-rest-api-layout/internal/consts"
	"github.com/nachogoca/golang-example-rest-api-layout/internal/entities"
	"github.com/nachogoca/golang-example-rest-api-layout/internal/middlewares"
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
	log := logrus.WithField("request_id", middlewares.GetRequestID(ctx))

	articles, err := a.store.GetAll(ctx)
	if err != nil {
		log.WithError(err).Error("could not get all articles")
		return nil, fmt.Errorf("could not get all articles: %w", err)
	}

	log.WithField("articles", len(articles)).Info("found articles")
	return articles, nil
}

// GetOne returns one article given an id
func (a Articles) GetOne(ctx context.Context, id string) (entities.Article, error) {
	log := logrus.WithField("request_id", middlewares.GetRequestID(ctx))

	article, err := a.store.GetOne(ctx, id)
	if err != nil {
		if errors.Is(err, consts.ErrEntityNotFound) {
			log.WithError(err).WithField("id", id).Warn("could not get article")
			return entities.Article{}, fmt.Errorf("article id %s not found %w", id, err)
		}

		log.WithError(err).WithField("id", id).Error("could not get article")
		return entities.Article{}, fmt.Errorf("could not get article id %s: %w", id, err)
	}

	log.WithField("article", article).Info("article retrieved")
	return article, nil
}

// Create creates an article
func (a Articles) Create(ctx context.Context, article entities.Article) (entities.Article, error) {
	log := logrus.WithField("request_id", middlewares.GetRequestID(ctx))

	// example of business logic applied, which should only live in the usecase layer
	if len(article.Content) > maxContentLen {
		return entities.Article{}, fmt.Errorf("article content is longer than allowed")
	}

	id := uuid.New().String()
	art := entities.Article{
		ID:        id,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title:     article.Title,
		Content:   article.Content,
		Author:    article.Author,
	}

	created, err := a.store.Create(ctx, art)
	if err != nil {
		return entities.Article{}, fmt.Errorf("could not create article: %w", err)
	}
	log.WithField("article", created).Info("article created")
	return created, nil
}

// Update updates the attributes of an article
func (a Articles) Update(ctx context.Context, article entities.Article) (entities.Article, error) {
	log := logrus.WithField("request_id", middlewares.GetRequestID(ctx))

	toUpdate, err := a.GetOne(ctx, article.ID)
	if err != nil {
		if errors.Is(err, consts.ErrEntityNotFound) {
			log.WithError(err).WithField("id", article.ID).Warn("could not get article")
			return entities.Article{}, fmt.Errorf("article id %s not found %w", article.ID, err)
		}

		log.WithError(err).WithField("id", article.ID).Error("could not get article")
		return entities.Article{}, fmt.Errorf("could not get article id %s: %w", article.ID, err)
	}
	log.WithField("article", article).Debug("found article to update")
	toUpdate.Title = article.Title
	toUpdate.Content = article.Content
	toUpdate.Author = article.Author
	toUpdate.UpdatedAt = time.Now().UTC()

	updated, err := a.store.Update(ctx, toUpdate)
	if err != nil {
		return entities.Article{}, fmt.Errorf("could not update article: %w", err)
	}
	log.WithField("article", updated).Info("article updated")

	return updated, nil
}

// Delete removes an article
func (a Articles) Delete(ctx context.Context, id string) error {
	log := logrus.WithField("request_id", middlewares.GetRequestID(ctx))

	toDelete, err := a.GetOne(ctx, id)
	if err != nil {
		if errors.Is(err, consts.ErrEntityNotFound) {
			log.WithError(err).WithField("id", id).Warn("could not get article")
			return fmt.Errorf("article id %s not found %w", id, err)
		}

		log.WithError(err).WithField("id", id).Error("could not get article")
		return fmt.Errorf("could not get article id %s: %w", id, err)
	}
	log.WithField("article", toDelete).Debug("found article to delete")

	if err := a.store.Delete(ctx, id); err != nil {
		return fmt.Errorf("could not delete article: %w", err)
	}
	log.WithField("id", id).Info("article deleted")

	return nil
}
