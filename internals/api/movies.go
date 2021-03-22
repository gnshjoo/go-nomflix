package api

import (
	"encoding/json"
	"github.com/gnshjoo/AppProject/internals/config"
	"github.com/gnshjoo/AppProject/internals/models"
	"io/ioutil"
	"net/http"
)

// GetMovieNowPlaying get movie api now playing list
func GetMovieNowPlaying() ([]models.Movie, error) {

	res, err := http.Get(config.URL + "movie/now_playing/?api_key=" + config.API_KEY + "&language=" + config.LANGUAGE)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r models.MovieResponse
	json.Unmarshal(data, &r)

	var v []models.Movie
	for _, val := range r.Results {
		v = append(v, val)
	}
	return v, nil
}

// GetMovieUpComing get movie api now upcoming list
func GetMovieUpComing() ([]models.Movie, error) {

	res, err := http.Get(config.URL + "movie/upcoming/?api_key=" + config.API_KEY + "&language=" + config.LANGUAGE)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r models.MovieResponse
	json.Unmarshal([]byte(data), &r)

	var v []models.Movie
	for _, val := range r.Results {
		v = append(v, val)
	}
	return v, nil
}

// GetMoviePopular get movie api popular list
func GetMoviePopular() ([]models.Movie, error) {
	res, err := http.Get(config.URL + "movie/popular/?api_key=" + config.API_KEY + "&language=" + config.LANGUAGE)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r models.MovieResponse
	json.Unmarshal(data, &r)

	var v []models.Movie
	for _, val := range r.Results {
		v = append(v, val)
	}
	return v, nil
}

// GetMovieDetail
func GetMovieDetail(id string) (models.MovieDetail, error) {
	var r models.MovieDetail

	res, err := http.Get(config.URL + "movie/"+ id +"?api_key="+config.API_KEY+"&language="+config.LANGUAGE)
	if err != nil {
		return r, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, err
	}

	json.Unmarshal(data, &r)

	return r, nil
}

// SearchMovie
func SearchMovie(query string) ([]models.Movie, error) {
	res, err := http.Get(config.URL + "search/movie/?api_key=" + config.API_KEY + "&language=" + config.LANGUAGE + "&query="+ query)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r models.MovieResponse
	json.Unmarshal(data, &r)

	var v []models.Movie
	for _, val := range r.Results {
		v = append(v, val)
	}
	return v, nil
}