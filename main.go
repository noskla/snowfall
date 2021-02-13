package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// Here pseudo-constant symbolizes the path of currently running file
var Here string
// Cfg represents contents of the configuration file.
var Cfg config

func main() {
	//gin.SetMode(gin.ReleaseMode)
	Here, _ = os.Executable()
	Cfg = loadConfig("config.json")

	router := gin.Default()

	staticPath := getDirectoryPath("static/")
	router.Static("/static", staticPath)

	InitFrontRouter(router)
	log.Fatalln(router.Run())
}
