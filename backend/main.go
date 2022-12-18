package main

import (
	routes "test-sharing-vision/app/route"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	//database.Migration(20)
	gin.Default()
	r := routes.SetupRoutes()
	r.Run()
}
