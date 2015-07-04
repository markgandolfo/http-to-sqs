package main

import (
	"gopkg.in/yaml.v2"
	"log"
)

type Config struct {
	Mappings map[string]Route
}

type Route struct {
	Path string
	Sqs  string
}

func parseYaml(data string) Config {
	config := Config{}

	err := yaml.Unmarshal([]byte(data), &config)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config
}
