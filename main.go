package main

import (
	"net/http"
	"url/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/shorten", handler.Shorten)

	router.GET("/:shortKey", handler.ShortKey)

	router.Run(":8080")
}
