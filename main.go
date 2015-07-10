package main

import (
	"flag"
	"fmt"
)

func main() {
	configFile := flag.String("config", "config.json", "Your config file")
	flag.Parse()

	data := readFile(*configFile)
	config := parseJSON(data)

	fmt.Printf("config:\n%+v\n", config)

	startServer(config)
}
