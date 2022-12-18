package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migration(step int) {
	db, err := sql.Open("mysql", "root:mysqlpw@tcp(host.docker.internal:3306)/article")
	if err != nil {
		panic(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}
	m, err2 := migrate.NewWithDatabaseInstance(
		"file://I:/test-sharing-vison/test-sharing-vison/backend/migrations",
		"mysql",
		driver,
	)
	if err2 != nil {
		panic(err2)
	}

	m.Steps(step)
	m.Up()
}
