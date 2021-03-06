package handler

import "net/http"
import "newsfeeder/plattform/newsfeed"
import "github.com/gin-gonic/gin"

func NewsfeedGet(feed newsfeed.NFGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}