package main

import (
	"log"
	"net/http"

	"ginblog/pkg/routes"

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

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	routes.InitRoutes(r)

	err := r.Run(":9090")
	if err != nil {
		log.Fatalln(err)
	}
}
