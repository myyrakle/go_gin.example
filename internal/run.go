package internal

import (
	"net/http"
	 _ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello!",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	router.Run()
}
