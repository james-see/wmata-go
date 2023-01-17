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
	// load the env from app.env

	mux := http.NewServeMux()
	mux.HandleFunc("/echo", echoHandler)
	http.ListenAndServe("localhost:5000", mux)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	url := "https://api.wmata.com/StationPrediction.svc/json/GetPrediction/C05"

	payload := strings.NewReader("")

	req, _ := http.NewRequest("GET", url, payload)

	req.Header.Add("api_key", config.WmataAPI)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)
	//
	trains := util.Train{}
	err = json.NewDecoder(res.Body).Decode(&trains)
	if err != nil {
		panic(err)
	}

	// car := util.Car{}
	// if err := json.Unmarshal(trains, &car); err != nil {
	// 	panic(err)
	// }
	destination := &util.CurrentStatus{}

	for _, car := range trains.Cars {
		if car.DestinationCode == "K08" {
			destination.Status = car.Min
			destination.Destination = car.Destination
			destination.LocationName = car.LocationName
			break
		}

	}
	fmt.Printf("Destination: %v\n", destination.Destination)
	fmt.Printf("Time to Board: %v\n", destination.Status)

	// fmt.Println(trains.Cars[0])

	// for index, train := range trains {
	// 	fmt.Println(train.CarID, index)
	// }
	// var p fastjson.Parser
	// v, _ := p.Parse(body)
	// fmt.Printf("foo=%v\n", v.GetStringBytes("0"))
	// fmt.Printf("foo.0=%v\n", fastjson.GetBytes(body))

	// fmt.Println(res)
	// fmt.Println(string(body))

}
