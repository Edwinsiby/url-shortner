package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var urlMap = make(map[string]string)

func generateShortKey() string {
	return uuid.New().String()[:8]
}

func Shorten(c *gin.Context) {
	longURL := c.PostForm("url")

	shortKey := generateShortKey()
	urlMap[shortKey] = longURL

	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortKey)

	c.HTML(http.StatusOK, "shortened.html", gin.H{
		"shortURL": shortURL,
	})
}

func ShortKey(c *gin.Context) {
	shortKey := c.Param("shortKey")
	longURL, exists := urlMap[shortKey]

	if !exists {
		c.String(http.StatusNotFound, "Short URL not found")
		return
	}

	c.Redirect(http.StatusFound, longURL)
}
