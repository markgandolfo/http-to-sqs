package main

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	configFile := flag.String("config", "config.json", "Your config file")
	flag.Parse()

	data := readFile(*configFile)
	config := parseJSON(data)
	fmt.Println(config)

	router := httprouter.New()
	router.GET("/*route", httpHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
