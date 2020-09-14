package stores

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"

	"github.com/nachogoca/golang-example-rest-api-layout/internal/consts"
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

	sqlStmt := `create table articles (	
		id text not null primary key, 
		created_at datetime not null,
		updated_at datetime not null,
		title text,
		content text,
		author text);`
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
	if err := os.Remove("./articles.db"); err != nil {
		logrus.WithError(err).Warn("could not delete sqlite db file")
		return err
	}
	return a.db.Close()
}

// GetAll returns all articles
func (a Articles) GetAll(ctx context.Context) ([]entities.Article, error) {

	query, _, err := sq.Select("id", "created_at", "updated_at", "title", "content", "author").
		From("articles").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("could not build getall query: %w", err)
	}

	logrus.WithField("query", query).Debug("query to get all articles")
	rows, err := a.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("could not execute get all articles query: %w", err)
	}

	var articles []entities.Article
	defer rows.Close()
	for rows.Next() {

		var article entities.Article
		err := rows.Scan(&article.ID,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.Title,
			&article.Content,
			&article.Author)
		if err != nil {
			return nil, fmt.Errorf("could not scan rows: %w", err)
		}
		articles = append(articles, article)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("got err while reading rows: %w", err)
	}

	return articles, nil
}

// GetOne returns one article by id
func (a Articles) GetOne(ctx context.Context, id string) (entities.Article, error) {

	query, args, err := sq.Select("id", "created_at", "updated_at", "title", "content", "author").
		From("articles").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return entities.Article{}, fmt.Errorf("could not build getone query: %w", err)
	}

	row := a.db.QueryRowContext(ctx, query, args...)
	var article entities.Article
	err = row.Scan(&article.ID,
		&article.CreatedAt,
		&article.UpdatedAt,
		&article.Title,
		&article.Content,
		&article.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Article{}, fmt.Errorf("article not found: %s: %w", err.Error(), consts.ErrEntityNotFound)
		}

		return entities.Article{}, fmt.Errorf("could not scan row: %w", err)
	}

	return article, nil

}

// Create inserts an article row
func (a Articles) Create(ctx context.Context, article entities.Article) (entities.Article, error) {

	query, args, err := sq.Insert("articles").
		Columns("id", "created_at", "updated_at", "title", "content", "author").
		Values(article.ID, article.CreatedAt, article.UpdatedAt, article.Title, article.Content, article.Author).
		ToSql()
	if err != nil {
		return entities.Article{}, fmt.Errorf("could not build query: %w", err)
	}

	logrus.WithField("query", query).
		WithField("args", args).
		Debug("query to insert")

	res, err := a.db.ExecContext(ctx, query, args...)
	if err != nil {
		return entities.Article{}, fmt.Errorf("could not exec insert query: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return entities.Article{}, fmt.Errorf("could not verify insertion: %w", err)
	}
	if affected != 1 {
		return entities.Article{}, fmt.Errorf("row was not inserted")
	}

	return a.GetOne(ctx, article.ID)
}

// Update looks for the row with the article id, and updates all the columns
func (a Articles) Update(ctx context.Context, article entities.Article) (entities.Article, error) {
	return entities.Article{}, fmt.Errorf("not implemented yet")

}

// Delete deletes the row
func (a Articles) Delete(ctx context.Context, id string) error {
	return fmt.Errorf("not implemented yet")
}
