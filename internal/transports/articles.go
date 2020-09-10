package transports

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/nachogoca/golang-example-rest-api-layout/internal/entities.go"
)

// run go generate ./... and the mocks will be generated
//
//go:generate mockgen -destination=./mocks/articles_mock.go -package=mocks github.com/nachogoca/golang-example-rest-api-layout/internal/transports ArticlesUsecase

// ArticlesUsecase describes all the functions we need from usecase layer
// create other interfaces as usecases needed
type ArticlesUsecase interface {
	GetAll(ctx context.Context) ([]entities.Article, error)
	GetOne(ctx context.Context, id string) (entities.Article, error)
	Create(ctx context.Context, article entities.Article) (entities.Article, error)
	Update(ctx context.Context, article entities.Article) (entities.Article, error)
	Delete(ctx context.Context, id string) error
}

// All HTTP (or whatever other communication protocol) specifics are handled here
// Requests are parsed from the required content type
// and responses are parsed to the required output content type
// Ensures that usecase functions are business only

// Articles is the transport struct
type Articles struct {
	usecase ArticlesUsecase
}

// NewArticles is the Articles transport constructor
func NewArticles(au ArticlesUsecase) Articles {
	return Articles{usecase: au}
}

// GetAll returns all articles
func (a Articles) GetAll(w http.ResponseWriter, r *http.Request) {
	articles, err := a.usecase.GetAll(r.Context())
	if err != nil {
		logrus.WithError(err).Error("could not get all articles")
		// TODO err according to err
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(articles)
	if err != nil {
		logrus.WithError(err).Error("could not marshal response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

// GetOne returns one article
func (a Articles) GetOne(w http.ResponseWriter, r *http.Request) {

}

// Create creates an article
func (a Articles) Create(w http.ResponseWriter, r *http.Request) {

}

// Update updates an article
func (a Articles) Update(w http.ResponseWriter, r *http.Request) {

}

// Delete deletes an article
func (a Articles) Delete(w http.ResponseWriter, r *http.Request) {

}
