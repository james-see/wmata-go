package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/james-see/wmata-go/util"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/echo", echoHandler)
	err := http.ListenAndServe("127.0.0.1:5000", mux)
	if err != nil {
		fmt.Println(err)
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// load the env from app.env
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// make the request
	url := "https://api.wmata.com/StationPrediction.svc/json/GetPrediction/C05"
	payload := strings.NewReader("body")
	req, _ := http.NewRequest("GET", url, payload)
	req.Header.Add("api_key", config.WmataAPI)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// put the data into a struct that holds the json array
	trains := util.Train{}
	err = json.NewDecoder(res.Body).Decode(&trains)
	fmt.Println(trains)
	if err != nil {
		panic(err)
	}

	// put the data back from the full struct and arrays into a single struct to filter out what we need
	var destination util.CurrentStatus
	var listofdestinations []util.CurrentStatus
	for _, car := range trains.Cars {
		if car.DestinationCode == "K08" || car.DestinationCode == "N12" {
			// only need to grab one of the statuses
			destination.Status = car.Min
			destination.Destination = car.Destination
			destination.LocationName = car.LocationName
			listofdestinations = append(listofdestinations, destination)
		}

	}
	for _, destiny := range listofdestinations {
		fmt.Printf("Destination: %v\n", destiny.Destination)
		fmt.Printf("Time to Board: %v\n", destiny.Status)
	}
}
