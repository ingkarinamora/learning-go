package main

import (
	"testing"
)

type mockWeatherProvider struct{}

func (m mockWeatherProvider) temperature(city string) (float64, error) {
	if city == "TestCity" {
		return 300.0, nil
	}
	return 0, nil
}

func TestHello(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hello)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Hello World"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestTemperature(t *testing.T) {
	mw := multiWeatherProvider{
		mockWeatherProvider{},
	}

	temp, err := mw.temperature("TestCity")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	expected := 300.0
	if temp != expected {
		t.Errorf("expected %v, got %v", expected, temp)
	}
}
