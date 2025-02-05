package main

import "testing"

func TestIsValidCep(t *testing.T) {
	tests := []struct {
		cep      string
		expected bool
	}{
		{"12345678", true},
		{"12345-678", false},
		{"abcdefgh", false},
		{"1234", false},
	}

	for _, test := range tests {
		if isValidCep(test.cep) != test.expected {
			t.Errorf("CEP %s validation failed", test.cep)
		}
	}
}

func TestRoundToTwo(t *testing.T) {
	if roundToTwo(28.5678) != 28.57 {
		t.Errorf("Expected 28.57, got %f", roundToTwo(28.5678))
	}
}

func TestGetLocation(t *testing.T) {
	cep := "01001000"
	location, err := getLocation(cep)
	if err != nil {
		t.Errorf("Error getting location: %v", err)
	}
	if location == "" {
		t.Errorf("Expected valid location, got empty string")
	}
}

func TestGetCurrentCelsiusTemp(t *testing.T) {
	location := "Fortaleza"
	temp, err := getCurrentCelsiusTemp(location)
	if err != nil {
		t.Errorf("Error fetching temperature: %v", err)
	}
	if temp == 0 {
		t.Errorf("Expected valid temperature, got 0")
	}
}
