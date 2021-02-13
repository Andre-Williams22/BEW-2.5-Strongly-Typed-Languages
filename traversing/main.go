package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	directory := "/Users/andrewilliams"
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// fmt.Println("Hello")
		fmt.Println(file.Name())
	}
}

// go build -o main
// ./main
// go run main.go
