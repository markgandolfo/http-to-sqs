package main

import (
	"fmt"
)

func main() {
	data := readFile("config.json")
	config := parseJSON(data)

	fmt.Printf("config:\n%+v\n", config)
}
