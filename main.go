package main

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Here pseudo-constant symbolizes the path of currently running file
var Here string
// Cfg represents contents of the configuration file.
var Cfg config
// Database represents active connection with the PostgreSQL database.
var Database *sql.DB

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//gin.SetMode(gin.ReleaseMode)

	Here, _ = os.Executable()
	Cfg = loadConfig("config.json")
	Database = connectToDatabase(Cfg.DBAddress, Cfg.DBUser, Cfg.DBPassword, Cfg.DBName, Cfg.DBSslMode)
	createTablesIfNotExists(Database)

	router := gin.Default()

	staticPath := getDirectoryPath("static/")
	router.Static("/static", staticPath)

	InitFrontRouter(router)
	log.Fatalln(router.Run())
}
