package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

var data = `
mappings:
  project1:
    path: /api/v1/pathname1
    sqs: my_sqs_queue1
  project2:
    path: /api/v1/pathname2
    sqs: my_sqs_queue2
`

type Config struct {
	Mappings map[string]Route
}

type Route struct {
	PATH string
	SQS  string
}

func main() {
	config := Config{}

	err := yaml.Unmarshal([]byte(data), &config)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("--- config:\n%+v\n", config)
}
