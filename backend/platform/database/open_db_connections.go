package database

import (
	"test-sharing-vision/app/queries"

	"github.com/jmoiron/sqlx"
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.ArticleQueries // from Article model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define Database connection variables.
	var (
		db  *sqlx.DB
		err error
	)

	db, err = MysqlConnection()

	if err != nil {
		return nil, err
	}
	return &Queries{
		// Set queries from models:
		ArticleQueries: &queries.ArticleQueries{DB: db}, // from User model

	}, nil
}
