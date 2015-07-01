package main

import (
	"fmt"
	"log"
	"gopkg.in/yaml.v2"
)


// var data = `
// a: Easy!
// b:
//   c: 2
//   d: [3, 4]
// `

// type T struct {
//         A string
//         B struct{C int; D []int ",flow"}
// }


var data = `---
project_name:
  path: /api/v1/create_article
  sqs: create_article
`

type Config struct {
	name string
	mapping struct {
		path string
		queue string
	}
}


func main() {
	config := Config{}

	err := yaml.Unmarshal([]byte(data), &config)

        if err != nil {
                log.Fatalf("error: %v", err)
        }
	
        fmt.Printf("--- config:\n%v\n\n", config)
}
