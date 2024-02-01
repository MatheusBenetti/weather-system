package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFetchViaCep_SuccessfulRequest(t *testing.T) {
	cep := "95670084"
	expected := &ViaCEP{
		Cep:         "95670-084",
		Logradouro:  "Rua Senador Salgado Filho",
		Complemento: "",
		Bairro:      "Centro",
		Localidade:  "Gramado",
		Uf:          "RS",
		Ibge:        "4309100",
		Gia:         "",
		Ddd:         "54",
		Siafi:       "8681",
	}

	result, err := fetchViaCep(cep)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestTemperatureConversion(t *testing.T) {
	// Call the function under test with a known Celsius temperature
	fahrenheit := celsiusToFahrenheit(25)
	kelvin := celsiusToKelvin(25)

	// Check if the conversion is correct
	assert.Equal(t, 77.0, fahrenheit)
	assert.Equal(t, 298.0, kelvin)
}

func TestServerListening(t *testing.T) {
	// Start the server
	go main()

	// Wait for the server to start listening
	time.Sleep(1 * time.Second)

	// Send a GET request to the server
	resp, err := http.Get("http://localhost:8080/getTemperature?cep=95670084")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}
}
