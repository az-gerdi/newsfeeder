package handler

import "net/http"
import "github.com/gin-gonic/gin"

func PingGet(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string {
		"hello": "gin retuns hallo message",
	})
}