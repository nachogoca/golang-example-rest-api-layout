package stores

import (
	"context"
	"fmt"

	"github.com/nachogoca/golang-example-rest-api-layout/internal/entities"
)

type Articles struct {
}

func NewArticle() Articles {
	return Articles{}
}

func (a Articles) GetAll(ctx context.Context) ([]entities.Article, error) {

	return nil, fmt.Errorf("not implemented yet")

}
func (a Articles) GetOne(ctx context.Context, id string) (entities.Article, error) {
	return entities.Article{}, fmt.Errorf("not implemented yet")

}
func (a Articles) Create(ctx context.Context, article entities.Article) (entities.Article, error) {
	return entities.Article{}, fmt.Errorf("not implemented yet")

}
func (a Articles) Update(ctx context.Context, article entities.Article) (entities.Article, error) {
	return entities.Article{}, fmt.Errorf("not implemented yet")

}
func (a Articles) Delete(ctx context.Context, id string) error {
	return fmt.Errorf("not implemented yet")
}
