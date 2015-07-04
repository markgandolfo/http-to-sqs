package main

import (
	"fmt"
)

func main() {
	data := readFile("config.yml")
	config := parseYaml(data)

	fmt.Printf("--- config:\n%+v\n", config)
}
