package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testTempSystem(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate the Weather API response
		responseJSON := `{"location":{"name":"TestLocation"},"current":{"temp_c":20.0,"temp_f":68.0,"temp_k":293.15}}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseJSON))
	}))
	defer testServer.Close()

	// Save the original Weather API URL and set the test server URL
	originalURL := weatherAPIURL
	weatherAPIURL = testServer.URL
	defer func() {
		// Reset the original Weather API URL after the test
		weatherAPIURL = originalURL
	}()

	// Call the function with a test location
	location := "TestLocation"
	result, err := fetchWeatherAPI(location)
	assert.NoError(t, err, "fetchWeatherAPI should not return an error")

	// Verify the expected values using testify's assert package
	expectedResult := &WeatherResponse{
		TempC: 20.0,
		TempF: 68.0,
		TempK: 293.15,
	}

	assert.Equal(t, expectedResult, result, "Unexpected result")
}
