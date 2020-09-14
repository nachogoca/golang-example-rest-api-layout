package transports

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/nachogoca/golang-example-rest-api-layout/internal/consts"
	"github.com/nachogoca/golang-example-rest-api-layout/internal/entities"
	"github.com/nachogoca/golang-example-rest-api-layout/internal/middlewares"
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
	ctx := r.Context()
	log := logrus.WithField("request_id", middlewares.GetRequestID(ctx))

	articles, err := a.usecase.GetAll(ctx)
	if err != nil {
		log.WithError(err).Error("could not get all articles")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// To respond empty array instead of nil
	if articles == nil {
		articles = []entities.Article{}
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		log.WithError(err).Error("could not encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetOne returns one article
func (a Articles) GetOne(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logrus.WithField("request_id", middlewares.GetRequestID(ctx))

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.WithField("vars", vars).Error("id not provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	article, err := a.usecase.GetOne(ctx, id)
	if err != nil {
		if errors.Is(err, consts.ErrEntityNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		log.WithError(err).Error("could not find article")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(article); err != nil {
		log.WithError(err).Error("could not encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Create creates an article
func (a Articles) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logrus.WithField("request_id", middlewares.GetRequestID(ctx))

	var article entities.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		log.WithError(err).Error("could not decode request body into article entity")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	created, err := a.usecase.Create(ctx, article)
	if err != nil {
		log.WithError(err).Error("could not create article")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(created); err != nil {
		log.WithError(err).Error("could not encode article response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Update updates an article
func (a Articles) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logrus.WithField("request_id", middlewares.GetRequestID(ctx))

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.WithField("vars", vars).Error("id not provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var article entities.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		log.WithError(err).Error("could not decode request body into article entity")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	article.ID = id

	created, err := a.usecase.Update(ctx, article)
	if err != nil {
		log.WithError(err).Error("could not update article")
		if errors.Is(err, consts.ErrEntityNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(created); err != nil {
		log.WithError(err).Error("could not encode article response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Delete deletes an article
func (a Articles) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logrus.WithField("request_id", middlewares.GetRequestID(ctx))

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.WithField("vars", vars).Error("id not provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := a.usecase.Delete(ctx, id); err != nil {
		log.WithError(err).Error("could not delete article")
		if errors.Is(err, consts.ErrEntityNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
