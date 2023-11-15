package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	bootServer()
}

func bootServer() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"test": "Main website",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
