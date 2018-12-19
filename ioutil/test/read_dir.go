package test

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadDir() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
