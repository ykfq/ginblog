package handlers

import (
	"net/http"
	"strconv"

	"ginblog/pkg/models"

	"github.com/gin-gonic/gin"
)

// Render one of HTML, JSON or CSV based on the Accept header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Response with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Response with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Response with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	// Call the HTML methord of the Context to render a template
	// c.HTML(
	// 	// Set the HTTP status code to 200(OK)
	// 	http.StatusOK,
	// 	// Use the index.html template
	// 	"index.html",
	// 	// Pass the data that the page uses
	// 	gin.H{
	// 		"title":   "Home Page",
	// 		"payload": articles,
	// 	},
	// )
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

func GetArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleId, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := models.GetArticleByID(articleId); err == nil {
			// Call the HTML method of the context to render a template
			c.HTML(
				// Set the HTTP status code to 200 (OK)
				http.StatusOK,
				// Use the article.html template
				"article.html",
				// Pass the data that the page uses
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
