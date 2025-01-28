package main

import (
	"encoding/json"
	"net/http"
)

type Input struct {
	Cep string `json:"cep"`
}
type ViaCepResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/busca", getWeatherHandler)

	http.ListenAndServe(":8084", mux)

}

func getWeatherHandler(w http.ResponseWriter, r *http.Request) {

	var inputDTO Input
	err := json.NewDecoder(r.Body).Decode(&inputDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//var cep string

	url_get_loc := "https://viacep.com.br/ws/" + inputDTO.Cep + "/json"

	res, err := http.Get(url_get_loc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	var output ViaCepResponse

	if err = json.NewDecoder(res.Body).Decode(&output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url_get_weather := "http://api.weatherapi.com/v1"

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)

}
http://api.weatherapi.com/v1