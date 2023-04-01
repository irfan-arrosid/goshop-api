package utils

import (
	"encoding/json"
	"goshop-api/models"
	"io/ioutil"
	"net/http"
)

const BASE_URL = "https://emsifa.github.io/api-wilayah-indonesia/api"

func ProvinceGetById(province_id string) (*models.Province, error) {
	response, err := http.Get(BASE_URL + "/province/" + province_id + ".json")
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject models.Province
	json.Unmarshal(responseData, &responseObject)

	return &responseObject, nil
}

func ProvinceGetAll() ([]models.Province, error) {
	response, err := http.Get(BASE_URL + "/province.json")
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject []models.Province
	json.Unmarshal(responseData, &responseObject)

	return responseObject, nil
}

func CityGetById(city_id string) (*models.City, error) {
	response, err := http.Get(BASE_URL + "/regency/" + city_id + ".json")
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject models.City
	json.Unmarshal(responseData, &responseObject)

	return &responseObject, nil
}

func CityGetAll() ([]models.City, error) {
	response, err := http.Get(BASE_URL + "/regency.json")
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject []models.City
	json.Unmarshal(responseData, &responseObject)

	return responseObject, nil
}
