package main

import (
	"log"

	"./test"
)

func main() {
	err := test.Pipe()
	if err != nil {
		log.Fatalln(err)
	}
}
