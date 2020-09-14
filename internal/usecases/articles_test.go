package usecases

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nachogoca/golang-example-rest-api-layout/internal/entities"
	"github.com/nachogoca/golang-example-rest-api-layout/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
)

func TestArticles_Create(t *testing.T) {
	type fields struct {
		store     ArticlesStore
		mockStore func(m *mocks.MockArticlesStore)
	}
	type args struct {
		article entities.Article
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Article
		wantErr bool
	}{
		{
			name: "Success: Create article",
			fields: fields{
				mockStore: func(m *mocks.MockArticlesStore) {
					m.EXPECT().
						Create(gomock.Any(), gomock.AssignableToTypeOf(entities.Article{})).
						Return(entities.Article{
							Title:   "title",
							Content: "content",
							Author:  "author",
						}, nil)
				},
			},
			args: args{
				article: entities.Article{
					Title:   "title",
					Content: "content",
					Author:  "author",
				},
			},
			want: entities.Article{
				Title:   "title",
				Content: "content",
				Author:  "author",
			},
			wantErr: false,
		},
		{
			name: "Failure: Store returns error",
			fields: fields{
				mockStore: func(m *mocks.MockArticlesStore) {
					m.EXPECT().
						Create(gomock.Any(), gomock.AssignableToTypeOf(entities.Article{})).
						Return(entities.Article{}, fmt.Errorf("internal error"))
				},
			},
			args: args{
				article: entities.Article{
					Title:   "title",
					Content: "content",
					Author:  "author",
				},
			},
			want:    entities.Article{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			m := mocks.NewMockArticlesStore(ctrl)
			if tt.fields.mockStore != nil {
				tt.fields.mockStore(m)
			}

			a := Articles{
				store: m,
			}

			got, err := a.Create(context.Background(), tt.args.article)
			if (err != nil) != tt.wantErr {
				t.Errorf("Articles.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, got, tt.want)
		})
	}
}
