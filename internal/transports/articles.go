package transports

import (
	"net/http"
)

type ArticlesUsecase interface {
}

type Articles struct {
	usecase ArticlesUsecase
}

func NewArticles(au ArticlesUsecase) Articles {
	return Articles{usecase: au}
}

func (a Articles) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (a Articles) GetOne(w http.ResponseWriter, r *http.Request) {

}

func (a Articles) Create(w http.ResponseWriter, r *http.Request) {

}

func (a Articles) Update(w http.ResponseWriter, r *http.Request) {

}

func (a Articles) Delete(w http.ResponseWriter, r *http.Request) {

}
