package api

import (
	"net/http"
	"encoding/json"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocation(url string) (*LocationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	
	var location LocationArea
	if err := json.NewDecoder(res.Body).Decode(&location); err != nil {
		return nil, err
	}
	return &location, nil
}