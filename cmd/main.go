package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/james-see/wmata-go/util"
)

func main() {
	// load the env from app.env
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	url := "https://api.wmata.com/StationPrediction.svc/json/GetPrediction/C05"

	payload := strings.NewReader("{body}")

	req, _ := http.NewRequest("GET", url, payload)

	req.Header.Add("api_key", config.WmataAPI)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
