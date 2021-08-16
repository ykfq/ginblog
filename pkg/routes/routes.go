package routes

import (
	"ginblog/pkg/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// Handle the index route
	r.GET("/index", showIndex)
	r.GET("/", handlers.ShowIndexPage)
	r.GET("/article/view/:article_id", handlers.GetArticle)
}

func showIndex(c *gin.Context) {
	// Call the HTML method fo Context to render a template
	c.HTML(
		// Set the http status to 200(ok)
		http.StatusOK,
		// use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case ,"title")
		gin.H{
			"title": "Home Page",
		},
	)
}
