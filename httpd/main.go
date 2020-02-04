package main

import "newsfeeder/httpd/handler"
import "newsfeeder/plattform/newsfeed"
import "github.com/gin-gonic/gin"

func main() {
	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet)
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// import "fmt"
	// import "newsfeeder/plattform/newsfeed"
	// feed := newsfeed.New()
	// fmt.Println(feed)
	// feed.Add(newsfeed.Item{"Hello", "Das ist ein test"})
	// fmt.Println(feed)
}
