package main

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func TestValidCep(t *testing.T) {
	go main()
	time.Sleep(1 * time.Second)
	resp, err := http.Get("http://localhost:8080/getTemperature?cep=12345678")
	if err != nil {
		t.Fatalf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
	var data Current
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
}
