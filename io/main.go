package main

import (
	"log"

	"./test"
)

func main() {
	err := test.ReadFull()
	if err != nil {
		log.Fatalln(err)
	}
}
