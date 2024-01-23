package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/getTemperature", func(w http.ResponseWriter, r *http.Request) {
		cep := r.URL.Query().Get("cep")

		if len(cep) != 8 {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("invalid zipcode"))
			return
		}

		viaCEP, err := fetchViaCep(cep)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("can not found zipcode"))
			return
		}

		weather, err := fetchWeatherAPI(viaCEP.Localidade)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error getting weather data"))
			return
		}

		// Responda com as temperaturas formatadas
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(weather)
	})
}

func fetchViaCep(cep string) (*ViaCEP, error) {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")

	if err != nil {
		return nil, fmt.Errorf("failed to make request to ViaCEP API: %v", err)
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var data ViaCEP
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return &data, nil
}

func fetchWeatherAPI(location string) {
	return &WeatherResponse{
		TempC: 28.5,
		TempF: celsiusToFahrenheit(28.5),
		TempK: celsiusToKelvin(28.5),
	}, nil
}
