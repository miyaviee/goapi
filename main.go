package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naoina/genmai"

	_ "github.com/lib/pq"
)

var db *genmai.DB

func initDB() *genmai.DB {
	db, err := genmai.New(&genmai.PostgresDialect{}, "postgres://postgres:@localhost/test?sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}

func errorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) != 0 {
		c.JSON(-1, c.Errors)
	}
}

func main() {
	db = initDB()
	if err := db.CreateTableIfNotExists(&Work{}); err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()
	r.Use(errorHandler)

	r.GET("/works", workIndex)
	r.GET("/works/:id", workGET)
	r.POST("/works", workPOST)
	r.PUT("/works/:id", workPUT)
	r.DELETE("/works/:id", workDELETE)

	r.Run(":8080")
}

func workIndex(c *gin.Context) {
	var works []Work
	if err := findWorks(&works); err != nil {
		c.AbortWithError(err.Code, err.Error)
		return
	}

	c.JSON(200, works)
}

func workGET(c *gin.Context) {
	var work Work
	if err := findWork(&work, c.Param("id")); err != nil {
		c.AbortWithError(err.Code, err.Error)
		return
	}

	c.JSON(200, work)
}

func workPOST(c *gin.Context) {
	var work Work
	c.BindJSON(&work)
	if err := work.Validate(); err != nil {
		c.AbortWithError(err.Code, err.Error)
		return
	}

	if err := saveWork(&work); err != nil {
		c.AbortWithError(err.Code, err.Error)
		return
	}

	c.JSON(200, work)
}

func workPUT(c *gin.Context) {
	var work Work
	if err := findWork(&work, c.Param("id")); err != nil {
		c.AbortWithError(err.Code, err.Error)
		return
	}

	c.BindJSON(&work)
	if err := work.Validate(); err != nil {
		c.AbortWithError(err.Code, err.Error)
		return
	}

	if err := saveWork(&work); err != nil {
		c.AbortWithError(err.Code, err.Error)
		return
	}

	c.JSON(200, work)
}

func workDELETE(c *gin.Context) {
	var work Work
	if err := findWork(&work, c.Param("id")); err != nil {
		c.AbortWithError(err.Code, err.Error)
		return
	}

	if err := deleteWork(&work); err != nil {
		c.AbortWithError(err.Code, err.Error)
		return
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}
