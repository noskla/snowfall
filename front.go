package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// FrontRouter represents the Gin router group dedicated for the front-end features.
var FrontRouter *gin.RouterGroup

// InitFrontRouter initializes the FrontRouter variable with Gin router group,
// loads HTML templates and parses SASS stylesheets.
func InitFrontRouter(router *gin.Engine) {
	templatePath := getDirectoryPath("templates/")

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
