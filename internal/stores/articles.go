package stores

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"

	"github.com/nachogoca/golang-example-rest-api-layout/internal/entities"
)

type Articles struct {
	db *sql.DB
}

func NewArticle() (Articles, error) {

	os.Remove("./articles.db")

	db, err := sql.Open("sqlite3", "./articles.db")
	if err != nil {
		logrus.WithError(err).Error("could not open sqlite file")
		return Articles{}, fmt.Errorf("could not open sqlite file: %w", err)
	}

	sqlStmt := `create table articles (id integer not null primary key, name text);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		logrus.WithError(err).WithField("statement", sqlStmt).Error("could not execute init statement")
		return Articles{}, fmt.Errorf("could not execute init statement %w: %s", err, sqlStmt)
	}

	if err := db.Ping(); err != nil {
		logrus.WithError(err).Error("could not ping database")
		return Articles{}, fmt.Errorf("could not ping database: %w", err)
	}

	return Articles{db}, nil
}

func (a Articles) Close() error {
	os.Remove("./articles.db")
	return a.db.Close()
}

// GetAll returns all articles
func (a Articles) GetAll(ctx context.Context) ([]entities.Article, error) {

	return nil, fmt.Errorf("not implemented yet")

}

// GetOne returns one article by id
func (a Articles) GetOne(ctx context.Context, id string) (entities.Article, error) {
	return entities.Article{}, fmt.Errorf("not implemented yet")

}

// Create inserts an article row
func (a Articles) Create(ctx context.Context, article entities.Article) (entities.Article, error) {

	query, args, err := sq.Insert("articles").
		Columns("id", "created_at", "updated_at", "title", "content", "author").
		Values(article.ID, article.CreatedAt, article.UpdatedAt, article.Title, article.Content, article.Author).
		ToSql()
	if err != nil {
		logrus.WithError(err).Error("could not build insert query")
		return entities.Article{}, fmt.Errorf("could not build query: %w", err)
	}

	logrus.WithField("query", query).
		WithField("args", args).
		Debug("query to insert")

	res, err := a.db.ExecContext(ctx, query, args...)
	if err != nil {
		logrus.WithError(err).Error("could not exec query")
		return entities.Article{}, fmt.Errorf("could not exec insert query: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		logrus.WithError(err).Error("could not get inserted id")
		return entities.Article{}, fmt.Errorf("could not get inserted id: %w", err)
	}

	logrus.WithField("id", id).Debug("inserted id")

	return entities.Article{}, fmt.Errorf("not implemented yet")

}

// Update looks for the row with the article id, and updates all the columns
func (a Articles) Update(ctx context.Context, article entities.Article) (entities.Article, error) {
	return entities.Article{}, fmt.Errorf("not implemented yet")

}

// Delete deletes the row
func (a Articles) Delete(ctx context.Context, id string) error {
	return fmt.Errorf("not implemented yet")
}
