package routes

import (
	"github.com/gnshjoo/AppProject/internals/handlers"
	"github.com/labstack/echo/v4"
)

func MovieRoute(e *echo.Group) {
	movieAddr := e.Group("/movie")
	movieAddr.GET("/now_playing", handlers.MovieNowPlaying)
	movieAddr.GET("/popular", handlers.MoviePopular)
	movieAddr.GET("/upcoming", handlers.MovieUpComing)
	movieAddr.GET("/movie/:id", handlers.MovieDetail)
	movieAddr.GET("/search/movie/:query", handlers.MovieSearch)
}

func TVRoute(e *echo.Group) {
	tvAddr := e.Group("/tv")
	tvAddr.GET("/top", handlers.TVTopRated)
	tvAddr.GET("/popular", handlers.TVPopular)
	tvAddr.GET("/on_air", handlers.TVOnAirToday)
	tvAddr.GET("/tv/:id", handlers.SearchTV)
	tvAddr.GET("/search/tv/:query", handlers.SearchTV)
}