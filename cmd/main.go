package main

import (
	"log"
)

func main() {
	if err := RunGRPC(); err != nil {
		log.Fatal(err)
	}
}
