package handler

import "net/http"
import "newsfeeder/plattform/newsfeed"
import "github.com/gin-gonic/gin"

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post string `json:"post"`
}

func NewsfeedPost(feed newsfeed.NFAdder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsfeedPostRequest{}
		c.Bind(&requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post: requestBody.Post,
		}

		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}