package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

type WeatherResponse struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

type Current struct {
	WeatherResponse `json:"current"`
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(weather.WeatherResponse)
	})

	log.Println("Server listening on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
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

func fetchWeatherAPI(location string) (*Current, error) {
	req, err := http.Get("http://api.weatherapi.com/v1/current.json?q=" + location + "&key=50dbab8a6094453b8d4214401242301")

	if err != nil {
		return nil, fmt.Errorf("failed to make request to Weather API: %v", err)
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var data Current
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	var fahrenheit float64 = data.WeatherResponse.TempC
	var kelvin float64 = data.WeatherResponse.TempC

	data.WeatherResponse.TempF = celsiusToFahrenheit(fahrenheit)
	data.WeatherResponse.TempK = celsiusToKelvin(kelvin)
	return &data, nil
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 1.8) + 32
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273
}
