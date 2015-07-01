package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

var data = `
project:
  path: /api/v1/pathname
  sqs: my_sqs_queue
`

type Config struct {
	PROJECT struct {
		PATH string
		SQS  string
	}
}

func main() {
	config := Config{}

	err := yaml.Unmarshal([]byte(data), &config)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("--- config:\n%+v\n", config)
}
