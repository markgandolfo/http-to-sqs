package main

import (
	"fmt"
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

func main() {
	config := parseYaml(data)

	fmt.Printf("--- config:\n%+v\n", config)
}
