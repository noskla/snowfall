package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Here pseudo-constant symbolizes the path of currently running file
var Here string

func main() {
	//gin.SetMode(gin.ReleaseMode)
	Here, _ = os.Executable()
	router := gin.Default()

	staticPath := getDirectoryPath("static/")
	router.Static("/static", staticPath)

	InitFrontRouter(router)
	router.Run()
}
