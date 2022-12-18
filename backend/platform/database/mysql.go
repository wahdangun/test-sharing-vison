package database

import (
	"fmt"
	"os"
	"strconv"
	"test-sharing-vision/utils"
	"time"

	_ "github.com/go-sql-driver/mysql" // load driver for Mysql
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

// MysqlConnection func for connection to Mysql database.
func MysqlConnection() (*sqlx.DB, error) {
	err := godotenv.Load()
	// Define database connection settings.
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))
	fmt.Println("db type", os.Getenv("DB_TYPE"))
	// Build Mysql connection URL.
	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		return nil, err
	}
	fmt.Println(mysqlConnURL)
	// Define database connection for Mysql.
	db, err := sqlx.Connect("mysql", mysqlConnURL)
	if err != nil {
		fmt.Println(mysqlConnURL)
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Set database connection settings:
	// 	- SetMaxOpenConns: the default is 0 (unlimited)
	// 	- SetMaxIdleConns: defaultMaxIdleConns = 2
	// 	- SetConnMaxLifetime: 0, connections are reused forever
	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
