package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"regexp"
)

type Input struct {
	Cep string `json:"cep"`
}
type ApiViaCepResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

type ApiWeatherResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}
type GetWeatherResponse struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/busca", getWeatherHandler)

	http.ListenAndServe(":8080", mux)

}

func getWeatherHandler(w http.ResponseWriter, r *http.Request) {

	var inputDTO Input
	err := json.NewDecoder(r.Body).Decode(&inputDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//var cep string
	if !isValidCep(inputDTO.Cep) {
		http.Error(w, "Invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	location, err := getLocation(inputDTO.Cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	celsiuTemp, err := getCurrentCelsiusTemp(location)
	if err != nil {
		if err.Error() == "can not fint zipcode" {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tempFahrenheit := celsiuTemp*1.8 + 32
	tempKelvin := celsiuTemp + 273
	output := GetWeatherResponse{
		roundToTwo(celsiuTemp),
		roundToTwo(tempFahrenheit),
		roundToTwo(tempKelvin),
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)

}

func roundToTwo(value float64) float64 {
	return math.Round(value*100) / 100
}

func getCurrentCelsiusTemp(location string) (float64, error) {
	apiKey := "f875c284c1114aec9c5220427250402"
	encodedLocation := url.QueryEscape(location)

	url_get_weather := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, encodedLocation)

	resp, err := http.Get(url_get_weather)
	if err != nil {

		return http.StatusNotFound, errors.New("can not find zipcode")
	}
	defer resp.Body.Close()
	var apiWeatherResponse ApiWeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&apiWeatherResponse)
	if err != nil {
		return 0, err
	}
	return apiWeatherResponse.Current.TempC, nil
}

func getLocation(cep string) (string, error) {
	url_get_loc := "https://viacep.com.br/ws/" + cep + "/json"

	res, err := http.Get(url_get_loc)
	if err != nil {

		return "", errors.New("Error getting location")
	}
	defer res.Body.Close()
	var output ApiViaCepResponse

	if err = json.NewDecoder(res.Body).Decode(&output); err != nil {
		return "", errors.New("Error getting location")
	}
	return output.Localidade, nil
}

func isValidCep(cep string) bool {
	rgx := regexp.MustCompile(`^\d{8}$`)
	return rgx.MatchString(cep)
}
