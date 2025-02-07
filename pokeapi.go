package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationResponse struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type config struct {
	NextURL string
	PrevURL string
	History []string
}

func commandMap(cfg *config) error {
	if cfg.NextURL == "" {
		cfg.NextURL = "https://pokeapi.co/api/v2/location-area/"
	}

	cfg.PrevURL = cfg.NextURL

	resp, err := http.Get(cfg.NextURL)
	if err != nil {
		return fmt.Errorf("failed to fetch location data: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	var data LocationResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
	cfg.NextURL = data.Next
	cfg.PrevURL = data.Previous

	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.PrevURL == "" {
		return fmt.Errorf("no previous locations to go back to")
	}

	// Use the PreviousURL to fetch the previous set of locations
	resp, err := http.Get(cfg.PrevURL)
	if err != nil {
		return fmt.Errorf("failed to fetch location data: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the JSON response
	var data LocationResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Display the location names
	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	// Update PreviousURL to the current NextURL and set the new NextURL
	cfg.PrevURL = cfg.NextURL
	cfg.NextURL = data.Next

	return nil
}
