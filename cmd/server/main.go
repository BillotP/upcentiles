package main

import (
	"log"
	"os"
	"upcentile/internal/api"
	"upcentile/internal/handler"

	"github.com/labstack/echo/v4"
)

var defaultPort = "8080"

func main() {
	var srv = echo.New()
	srv.HideBanner = true
	srv.HidePort = true
	srv.GET("/analysis", handler.GetAnalysisHandler)
	log.Printf("[INFO] server v%s listening on port %s", api.VERSION, defaultPort)
	err := srv.Start(":" + defaultPort)

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		os.Exit(1)
	}
}
