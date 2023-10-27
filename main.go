package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var urlMap = make(map[string]string)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/shorten", func(c *gin.Context) {
		longURL := c.PostForm("url")

		shortKey := generateShortKey()
		urlMap[shortKey] = longURL

		shortURL := fmt.Sprintf("http://localhost:8080/%s", shortKey)

		c.HTML(http.StatusOK, "shortened.html", gin.H{
			"shortURL": shortURL,
		})
	})

	router.GET("/:shortKey", func(c *gin.Context) {
		shortKey := c.Param("shortKey")
		longURL, exists := urlMap[shortKey]

		if !exists {
			c.String(http.StatusNotFound, "Short URL not found")
			return
		}

		c.Redirect(http.StatusFound, longURL)
	})

	router.Run(":8080")
}

func generateShortKey() string {
	return uuid.New().String()[:8]
}
