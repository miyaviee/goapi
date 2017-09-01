package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naoina/genmai"

	_ "github.com/lib/pq"
)

func initDB() *genmai.DB {
	db, err := genmai.New(&genmai.PostgresDialect{}, "postgres://postgres:@localhost/test?sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}

func systemError(c *gin.Context) {
	c.JSON(500, gin.H{
		"error":   true,
		"message": "system error.",
	})
}

func notFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"error":   true,
		"message": "not found.",
	})
}

func badRequest(c *gin.Context) {
	c.JSON(400, gin.H{
		"error":   true,
		"message": "bad request.",
	})
}

func main() {
	db := initDB()
	if err := db.CreateTableIfNotExists(&Work{}); err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/works", workIndex)
	r.GET("/works/:id", workGET)
	r.POST("/works", workPOST)
	r.PUT("/works/:id", workPUT)
	r.DELETE("/works/:id", workDELETE)

	r.Run(":8080")
}

func workIndex(c *gin.Context) {
	db := initDB()
	var works []Work
	if err := db.Select(&works); err != nil {
		systemError(c)
		return
	}

	if len(works) == 0 {
		notFound(c)
		return
	}

	c.JSON(200, works)
}

func workGET(c *gin.Context) {
	db := initDB()
	var works []Work
	if err := db.Select(&works, db.Where("id", "=", c.Param("id"))); err != nil {
		systemError(c)
		return
	}

	if len(works) == 0 {
		notFound(c)
		return
	}

	c.JSON(200, works[0])
}

func workPOST(c *gin.Context) {
	db := initDB()
	var work Work
	c.BindJSON(&work)
	if err := work.Validate(); err != nil {
		badRequest(c)
		return
	}

	if _, err := db.Insert(&work); err != nil {
		systemError(c)
		return
	}

	c.JSON(200, work)
}

func workPUT(c *gin.Context) {
	db := initDB()
	var works []Work
	if err := db.Select(&works, db.Where("id", "=", c.Param("id"))); err != nil {
		systemError(c)
		return
	}

	if len(works) == 0 {
		notFound(c)
		return
	}

	work := works[0]
	c.BindJSON(&work)
	if err := work.Validate(); err != nil {
		badRequest(c)
		return
	}

	if _, err := db.Update(&work); err != nil {
		systemError(c)
		return
	}

	c.JSON(200, work)
}

func workDELETE(c *gin.Context) {
	db := initDB()
	var works []Work
	if err := db.Select(&works, db.Where("id", "=", c.Param("id"))); err != nil {
		systemError(c)
		return
	}

	if len(works) == 0 {
		notFound(c)
		return
	}

	work := works[0]
	if _, err := db.Delete(&work); err != nil {
		systemError(c)
		return
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}
