package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	name, _ := os.Hostname()
	c.PureJSON(http.StatusOK, gin.H{
		"data": "Welcome Gin Server: " + name,
	})
}

func Ping(c *gin.Context) {
	c.Status(http.StatusOK)
}
