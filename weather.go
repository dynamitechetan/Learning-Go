package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const weatherUrl string = "http://api.openweathermap.org/data/2.5/weather?units=imperial"

type Payload struct {
	Coord   map[string]float64
	Weather []Weather
	Main    map[string]float64
	Wind    map[string]float64
	Clouds  map[string]float64
	Cod     int // Status code, i.e., 200 -> OK
	Name    string
}

type Weather struct {
	Id          float64
	Main        string
	Description string
}

func main() {
	// Command line arguments besides program name
	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("Usage: weather <zip code>")
		os.Exit(1)
	}
	args := os.Args[1:]
	zip := args[0]

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		apiKey = "YOUR API KEY HERE"
	}

	url := fmt.Sprintf("%s&zip=%s&appid=%s", weatherUrl, zip, apiKey)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var p Payload
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println()

	if p.Cod == 200 {
		fmt.Printf("Weather for %s\n", p.Name)
		fmt.Printf("%s: %s\n", p.Weather[0].Main, p.Weather[0].Description)
		fmt.Printf("Temp: %.2f F\n", p.Main["temp"])
		fmt.Printf("Humidity: %.1f %%\n", p.Main["humidity"])
		fmt.Printf("Wind speed: %.1f\n", p.Wind["speed"])
	} else {
		fmt.Println("There was an error fetching weather data.")
	}

	fmt.Println()
}