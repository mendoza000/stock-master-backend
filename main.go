package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/mendoza000/stockmaster/routes"
)

var db *sql.DB

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	loadEnv()
	mySqlUrl := os.Getenv("MYSQL_URL") 


	var err error
	db, err = sql.Open("mysql", mySqlUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	r := gin.Default()

	routes.ProductsRoutes(r, db)
	routes.UserRoutes(r, db)


	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}