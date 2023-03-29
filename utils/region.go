package utils

import (
	"encoding/json"
	"goshop-api/models"
	"io/ioutil"
	"net/http"
)

const BASE_URL = "https://emsifa.github.io/api-wilayah-indonesia/api"

func GetProvinceById(province_id string) (*models.Province, error) {
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

func GetAllProvince() ([]models.Province, error) {
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
