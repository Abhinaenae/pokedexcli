package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	//check cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		//fmt.Println("cache hit")
		locationAreasResp := LocationAreasResponse{}
		if err := json.Unmarshal(data, &locationAreasResp); err != nil {
			return LocationAreasResponse{}, err
		}
		return locationAreasResp, nil
	}
	//fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", err)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResp := LocationAreasResponse{}
	if err := json.Unmarshal(data, &locationAreasResp); err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationAreas, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	//check cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		//fmt.Println("cache hit")
		locationAreas := LocationAreas{}
		if err := json.Unmarshal(data, &locationAreas); err != nil {
			return LocationAreas{}, err
		}
		return locationAreas, nil
	}
	//fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreas{}, fmt.Errorf("bad status code: %v", err)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	locationAreas := LocationAreas{}
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreas, nil
}
