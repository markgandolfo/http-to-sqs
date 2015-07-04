package main

import (
	"io/ioutil"
	"log"
)

func readFile(file string) string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return string(dat)
}
