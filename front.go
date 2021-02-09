package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// FrontRouter represents the Gin router group dedicated for the front-end features.
var FrontRouter *gin.RouterGroup

// InitFrontRouter initializes the FrontRouter variable with Gin router group,
// loads HTML templates and parses SASS stylesheets.
func InitFrontRouter(router *gin.Engine) {
	templatePath := Here + "/templates/"
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		// try finding templates relatively
		templatePath = "templates/"
		if _, err = os.Stat(templatePath); os.IsNotExist(err) {
			log.Fatalln("Could not find template directory.")
		}
	}

	log.Println("Loading HTML templates at", templatePath)
	router.LoadHTMLGlob(templatePath + "*")
	FrontRouter = router.Group("")
	{
		FrontRouter.GET("/", frontIndexEndpoint)
	}
}

// GET / with no special arguments.
func frontIndexEndpoint(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
