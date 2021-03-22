package api

import (
	"encoding/json"
	"github.com/gnshjoo/AppProject/internals/config"
	"github.com/gnshjoo/AppProject/internals/models"
	"io/ioutil"
	"net/http"
)

// GetAiringToday get TV api on_air_today list
func GetAiringToday() ([]models.TV, error) {
	 res, err := http.Get(config.URL+"tv/on_the_air?api_key="+config.API_KEY+"&language="+config.LANGUAGE)
	 if err != nil {
	 	return nil, err
	 }

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r models.TVResponse
	json.Unmarshal([]byte(data), &r)

	var v []models.TV
	for _, val := range r.Results {
		v = append(v, val)
	}
	return v, nil
}

//
func GetPopular() ([]models.TV, error) {
	res, err := http.Get(config.URL+"tv/popular?api_key="+config.API_KEY+"&language="+config.LANGUAGE)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r models.TVResponse
	json.Unmarshal(data, &r)

	var v []models.TV
	for _, val := range r.Results {
		v = append(v, val)
	}
	return v, nil
}

func GetTopRated() ([]models.TV, error) {
	res, err := http.Get(config.URL+"tv/popular?api_key="+config.API_KEY+"&language="+config.LANGUAGE)
	if err != nil {
		 return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var r models.TVResponse
	json.Unmarshal(data, &r)

	var v []models.TV
	for _, val := range r.Results {
		v = append(v, val)
	}
	return v, nil
}

func GetTVDetail(id string) (models.TVDetail, error) {
	var v models.TVDetail
	res, err := http.Get(config.URL + "tv/" + id + "?api_key=" + config.API_KEY + "&language="+ config.LANGUAGE)
	if err != nil {
		return v, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return v, err
	}

	json.Unmarshal(data, &v)

	return v, nil
}

func SearchTV(query string) ([]models.TV, error) {
	res, err := http.Get(config.URL+"search/tv?api_key="+config.API_KEY+"&language="+config.LANGUAGE + "&query="+query)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var r models.TVResponse
	json.Unmarshal(data, &r)

	var v []models.TV
	for _, val := range r.Results {
		v = append(v, val)
	}
	return v, nil
}