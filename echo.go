package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func echoAPI() {
	e := echo.New()
	e.GET("/repo/:repo/id/:id", func(c echo.Context) error {
		repo := c.Param("repo")
		objectId := c.Param("id")

		return c.JSON(200, data[repo+"/"+objectId])
	})
	e.POST("/repo/:repo", func(c echo.Context) error {
		var d interface{}
		repo := c.Param("repo")

		if err := c.Bind(&d); err != nil {
			c.Response().Status = http.StatusBadRequest
			return nil
		}

		data[repo+"/"+strconv.Itoa(id)] = d
		id++
		c.Response().Status = http.StatusOK
		return nil
	})
	e.Logger.Fatal(e.Start(":8080"))
}
