package main

import (
	"encoding/json"
	"fmt"
)

type Configuration struct {
	Path  string
	Queue string
}

func parseJSON(data string) []Configuration {

	config := []Configuration{}
	err := json.Unmarshal([]byte(data), &config)
	if err != nil {
		fmt.Println("Could not load in config, error:", err)
	}
	return config
}
