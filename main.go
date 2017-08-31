package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naoina/genmai"

	_ "github.com/lib/pq"
)

var db *genmai.DB

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Hello, World",
		})
	})
	db, err := genmai.New(&genmai.PostgresDialect{}, "postgres://postgres:@localhost/test?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.CreateTableIfNotExists(&Work{}); err != nil {
		panic(err)
	}

	r.Run(":8080")
}
