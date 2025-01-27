package main

import (
	"encoding/json"
	"net/http"
)

type Input struct {
	Cep string `json:"cep"`
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/busca", getWeatherHandler)

}

func getWeatherHandler(w http.ResponseWriter, r *http.Request) {

	var inputDTO Input
	if err := json.NewDecoder(r.Body).Decode(&inputDTO); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
