package main

import (
	"upcentile/internal/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	var srv = echo.New()
	srv.HideBanner = true
	srv.GET("/analysis", handler.GetAnalysisHandler)
	srv.Start(":8080")
}
