package handlers

import (
	"github.com/gnshjoo/AppProject/internals/api"
	"github.com/labstack/echo/v4"
	"net/http"
)

func MovieNowPlaying(c echo.Context) error {
	data, err := api.GetMovieNowPlaying()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, data)
}

func MovieUpComing(c echo.Context) error {
	data, err := api.GetMovieUpComing()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, data)
}

func MoviePopular(c echo.Context) error {
	data, err := api.GetMoviePopular()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, data)
}

func MovieDetail(c echo.Context) error {
	id := c.FormValue("id")
	data, err := api.GetMovieDetail(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, data)
}

func MovieSearch(c echo.Context) error {
	query := c.FormValue("query")
	data, err := api.SearchMovie(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, data)
}

func TVOnAirToday(c echo.Context) error {
	data, err := api.GetAiringToday()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, data)
}

func TVTopRated(c echo.Context) error {
	data, err := api.GetTopRated()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, data)
	}

	return c.JSON(http.StatusOK, data)
}

func TVPopular(c echo.Context) error {
	data, err := api.GetPopular()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, data)
}

func TVDetail(c echo.Context) error {
	id := c.FormValue("id")

	data, err := api.GetTVDetail(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, data)
}

func SearchTV(c echo.Context) error {
	query := c.FormValue("query")

	data, err := api.SearchTV(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, data)
}