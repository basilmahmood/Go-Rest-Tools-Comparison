package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ginAPI() {
	r := gin.Default()
	r.GET("/repo/:repo/id/:id", func(c *gin.Context) {
		repo := c.Param("repo")
		objectId := c.Param("id")

		c.JSON(200, data[repo+"/"+objectId])
	})
	r.POST("/repo/:repo", func(c *gin.Context) {
		var d interface{}
		repo := c.Param("repo")

		if err := c.ShouldBindJSON(&d); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		data[repo+"/"+strconv.Itoa(id)] = d
		id++
		c.Status(http.StatusOK)
	})
	r.Run()
}
