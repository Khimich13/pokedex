package api

import (
	"net/http"
	"encoding/json"
	"github.com/Khimich13/pokedex/internal/pokecache"
	"time"
	"io"
	"fmt"
)

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	URL            string `json:"url"`
	BaseExperience int    `json:"base_experience"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
} 

const LocationAreaUrl string = "https://pokeapi.co/api/v2/location-area/"
const PokemonUrl      string = "https://pokeapi.co/api/v2/pokemon/"
var cache = pokecache.NewCache(5 * time.Second)

func GetData[T any](url string) (*T, error) {
	var body []byte
	if val, exist := cache.Get(url); exist {
		body = val
	} else {
		var err error
		body, err = GetBodyFromUrl(url)
		if err != nil {
			return nil, err
		}
		cache.Add(url, body)
	}
	
	var data T
	if err := json.Unmarshal(body, &data); err != nil {
    	return nil, err
	}
	return &data, nil
}

func GetBodyFromUrl(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("error occured: %s", res.Status)
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}