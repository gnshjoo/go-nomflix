package main

import (
	"github.com/gnshjoo/AppProject/internals/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Debug = true

	// middleware
	e.Use(middleware.Logger())

	// routes
	v1 := e.Group("/v1")
	routes.MovieRoute(v1)
	routes.TVRoute(v1)



	e.Logger.Fatal(e.Start(":8080"))
}
