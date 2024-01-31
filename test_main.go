package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTemperatureHandler(t *testing.T) {
	// Define a sample address
	address := &ViaCEP{
		Cep:         "01001000",
		Logradouro:  "Praça da Sé",
		Complemento: "",
		Bairro:      "Sé",
		Localidade:  "São Paulo",
		Uf:          "SP",
		Ibge:        "3550308",
		Gia:         "1004",
		Ddd:         "11",
		Siafi:       "7107",
	}

	// Define a sample weather response
	weather := &Current{
		WeatherResponse: WeatherResponse{
			TempC: 30.5,
			TempF: 86.9,
			TempK: 303.65,
		},
	}

	// Create a test server and register the handler
	ts := httptest.NewServer(http.HandlerFunc(getTemperatureHandler))
	defer ts.Close()

	// Set up the expected response
	expectedResponse := &WeatherResponse{
		TempC: 30.5,
		TempF: 86.9,
		TempK: 303.65,
	}
	expectedBody, _ := json.Marshal(expectedResponse)

	// Make a request to the test server with a valid zip code
	req, _ := http.NewRequest("GET", ts.URL+"/getTemperature?cep="+address.Cep, nil)
	w := httptest.NewRecorder()
	getTemperatureHandler(w, req)

	// Check that the response status code is 200
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check that the response body matches the expected response
	if string(w.Body.Bytes()) != string(expectedBody) {
		t.Errorf("Expected response body %s, got %s", expectedBody, w.Body.Bytes())
	}
}

func getTemperatureHandler(w http.ResponseWriter, r *http.Request) {
	// Call the main function's handler with a mock fetchViaCep function
	// that returns the sample address
	FetchViaCep = func(cep string) (*ViaCEP, error) {
		return address, nil
	}

	// Call the main function's handler with a mock fetchWeatherAPI function
	// that returns the sample weather response
	FetchWeatherAPI = func(location string) (*Current, error) {
		return weather, nil
	}

	main()
}

// Define a global variable to hold the sample address
var address *ViaCEP

// Define a global variable to hold the sample weather response
var weather *Current

// Mock the fetchViaCep function for testing
var FetchViaCep = func(cep string) (*ViaCEP, error) {
	address = &ViaCEP{
		Cep:         "01001000",
		Logradouro:  "Praça da Sé",
		Complemento: "",
		Bairro:      "Sé",
		Localidade:  "São Paulo",
		Uf:          "SP",
		Ibge:        "3550308",
		Gia:         "1004",
		Ddd:         "11",
		Siafi:       "7107",
	}
	return address, nil
}

// Mock the fetchWeatherAPI function for testing
var FetchWeatherAPI = func(location string) (*Current, error) {
	weather = &Current{
		WeatherResponse: WeatherResponse{
			TempC: 30.5,
			TempF: 86.9,
			TempK: 303.65,
		},
	}
	return weather, nil
}
