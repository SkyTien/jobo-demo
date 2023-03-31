package main

import (
	_ "Goal-Back-End/conf"
	"Goal-Back-End/database"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	fmt.Println("Starting...")
	gin.SetMode(gin.DebugMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	//router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	time.Sleep(5 * time.Second)
	//defer db.SqlDB.Close()
	fmt.Println("Init DB...")
	database.Init()

	// Listen and Server in 0.0.0.0:8080
	fmt.Println("Running...")
	router.Run("0.0.0.0:8080")
}
