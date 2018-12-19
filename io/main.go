package main

import (
	"log"

	"./test"
)

func main() {
	err := test.Seek()
	if err != nil {
		log.Fatalln(err)
	}
}
