package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Province struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type City struct {
	Id          string `json:"id"`
	Province_id string `json:"province_id"`
	Name        string `json:"name"`
}

func GetProvinceById(id int) (*Province, error) {

	url, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")

	if err != nil {
		return nil, err
	}

	defer url.Body.Close()

	var province []Province

	if err := json.NewDecoder(url.Body).Decode(&province); err != nil {
		return nil, err
	}

	for _, x := range province {
		if x.Id == fmt.Sprintf("%d", id) {
			return &x, nil
		}
	}

	return nil, fmt.Errorf("Province With %d Not Found", id)

}

func GetCityById(provinceId, cityId int) (*City, error) {

	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%d.json", provinceId)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var city []City

	if err := json.NewDecoder(resp.Body).Decode(&city); err != nil {
		return nil, err
	}

	for _, x := range city {
		if x.Id == fmt.Sprintf("%d", cityId) {
			return &x, nil
		}
	}

	return nil, fmt.Errorf("City  With %d Not Found", cityId)
}
