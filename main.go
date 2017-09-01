package main

import (
	"errors"

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

func errorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) != 0 {
		c.JSON(-1, c.Errors)
	}
}

func main() {
	db := initDB()
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
	db := initDB()
	var works []Work
	if err := db.Select(&works); err != nil {
		c.AbortWithError(500, err)
		return
	}

	if len(works) == 0 {
		c.AbortWithError(404, errors.New("not found."))
		return
	}

	c.JSON(200, works)
}

func workGET(c *gin.Context) {
	db := initDB()
	var works []Work
	if err := db.Select(&works, db.Where("id", "=", c.Param("id"))); err != nil {
		c.AbortWithError(500, err)
		return
	}

	if len(works) == 0 {
		c.AbortWithError(404, errors.New("not found."))
		return
	}

	c.JSON(200, works[0])
}

func workPOST(c *gin.Context) {
	db := initDB()
	var work Work
	c.BindJSON(&work)
	if err := work.Validate(); err != nil {
		c.AbortWithError(400, err)
		return
	}

	if _, err := db.Insert(&work); err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, work)
}

func workPUT(c *gin.Context) {
	db := initDB()
	var works []Work
	if err := db.Select(&works, db.Where("id", "=", c.Param("id"))); err != nil {
		c.AbortWithError(500, err)
		return
	}

	if len(works) == 0 {
		c.AbortWithError(404, errors.New("not found."))
		return
	}

	work := works[0]
	c.BindJSON(&work)
	if err := work.Validate(); err != nil {
		c.AbortWithError(400, err)
		return
	}

	if _, err := db.Update(&work); err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, work)
}

func workDELETE(c *gin.Context) {
	db := initDB()
	var works []Work
	if err := db.Select(&works, db.Where("id", "=", c.Param("id"))); err != nil {
		c.AbortWithError(500, err)
		return
	}

	if len(works) == 0 {
		c.AbortWithError(404, errors.New("not found."))
		return
	}

	work := works[0]
	if _, err := db.Delete(&work); err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}
