package util

import "github.com/spf13/viper"

type Config struct {
	WmataAPI string `mapstructure:"API_KEY"`
}

type Train struct {
	Cars []Car `json:"Trains"`
}
type Car struct {
	CarID           string `json:"Car"`
	Destination     string `json:"Destination"`
	DestinationCode string `json:"DestinationCode"`
	DestinationName string `json:"DestinationName"`
	Group           string `json:"Group"`
	Line            string `json:"Line"`
	LocationCode    string `json:"LocationCode"`
	LocationName    string `json:"LocationName"`
	Min             string `json:"Min"`
}

type CurrentStatus struct {
	Status       string
	LocationName string
	Destination  string
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
