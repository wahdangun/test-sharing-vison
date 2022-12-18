package queries

import (
	"fmt"
	"test-sharing-vision/app/models"

	"github.com/jmoiron/sqlx"
)

// ArticleQueries struct for queries from article model.
type ArticleQueries struct {
	*sqlx.DB
}

func (q *ArticleQueries) GetArticle() ([]models.Article, error) {
	// Define   variable.
	article := []models.Article{}

	// Define query string.
	query := `SELECT * FROM articles`

	// Send query to database.
	err := q.Get(&article, query)
	if err != nil {
		// Return empty object and error.
		return article, err
	}

	// Return query result.
	return article, nil
}

// get article by id
func (q *ArticleQueries) GetArticleById(id string) (models.Article, error) {
	// Define   variable.
	article := models.Article{}

	// Define query string.
	query := `SELECT * FROM articles WHERE id = ?`

	// Send query to database.
	err := q.Get(&article, query, id)
	if err != nil {
		// Return empty object and error.
		return article, err
	}

	// Return query result.
	return article, nil
}

// get article with pagination
func (q *ArticleQueries) GetArticleWithPagination(offset int, limit int) ([]models.Article, error) {
	// Define   variable.
	article := []models.Article{}

	// Define query string.
	query := `SELECT id, title, content, status, category, created_date, updated_date FROM articles LIMIT ? OFFSET ?`
	// Send query to database.
	err := q.Select(&article, query, limit, offset)
	if err != nil {
		fmt.Println(err)
		// Return empty object and error.
		return article, err
	}

	// Return query result.
	return article, nil
}

// update article by id
func (q *ArticleQueries) UpdateArticleById(article models.Article) (models.Article, error) {
	// Define query string.
	query := `UPDATE articles SET title = ?, content = ?, category = ?, status = ? WHERE id = ?`

	// Send query to database.
	_, err := q.Exec(query, article.Title, article.Content, article.Category, article.Status, article.Id)
	if err != nil {
		// Return empty object and error.
		return article, err
	}
	return article, nil
}

// delete article by id
func (q *ArticleQueries) DeleteArticleById(id string) (models.Article, error) {
	// Define   variable.
	article := models.Article{}

	// Define query string.
	query := `DELETE FROM articles WHERE id = ?`

	// Send query to database.
	_, err := q.Exec(query, id)
	if err != nil {
		// Return empty object and error.
		return article, err
	}
	return article, nil
}

// Create article
func (q *ArticleQueries) CreateArticle(article models.Article) (models.Article, error) {
	// Define query string.
	query := `INSERT INTO articles (title, content, category, status) VALUES (?, ?, ?, ?)`

	// Send query to database.
	_, err := q.Exec(query, article.Title, article.Content, article.Category, article.Status)
	if err != nil {
		// Return empty object and error.
		return article, err
	}
	return article, nil
}
