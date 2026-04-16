package main

import (
	"fmt"
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

func printGiven20(url string, state *config) {
	location, err := getLocation(url)
	if err != nil{
		fmt.Println(err)
		return
	}
	for _, result := range location.Results {
		fmt.Println(result.Name)
	}
	state.Next = location.Next
	state.Previous = location.Previous
}

func getLocation(url string) (*LocationArea, error) {
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