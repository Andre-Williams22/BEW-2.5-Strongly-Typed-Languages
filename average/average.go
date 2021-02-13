package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Average from 3 inputs
	fmt.Print("Enter 1st weight: ")
	var input string
	fmt.Scanln(&input)
	fmt.Print(input)

	number, _ := strconv.ParseUint(input, 10, 32)
	input1 := int(number) //Convert uint64 To int

	fmt.Print("Enter 2nd weight: ")
	var input2 string
	fmt.Scanln(&input2)
	fmt.Print(input2)

	number2, _ := strconv.ParseUint(input2, 10, 32)
	int2 := int(number2) //Convert uint64 To int

	fmt.Print("Enter 3rd weight: ")
	var input3 string
	fmt.Scanln(&input2)
	fmt.Print(input3)

	number3, _ := strconv.ParseUint(input, 10, 32)
	int3 := int(number3) //Convert uint64 To int

	fmt.Print("Enter 4th weight: ")
	var input4 string
	fmt.Scanln(&input4)
	fmt.Print(input4)

	number4, _ := strconv.ParseUint(input, 10, 32)
	int4 := int(number4) //Convert uint64 To int

	fmt.Print("Enter 5th weight: ")
	var input5 string
	fmt.Scanln(&input5)
	fmt.Print(input5)

	number5, _ := strconv.ParseUint(input, 10, 32)
	int5 := int(number5) //Convert uint64 To int

	// // declaring an array of values
	// arra := []int{1, 2, 3, 4}
	arra := []int{input1, int2, int3, int4, int5}

	// size of the array
	n := len(arra)

	// declaring a variable
	// to store the sum
	sum := 0

	// traversing through the
	// array using for loop
	for i := 0; i < n; i++ {

		// adding the values of
		// array to the variable sum
		sum += (arra[i])
	}
	// declaring a variable
	// avg to find the average
	avg := (float64(sum)) / (float64(n))

	// typecast all values to float
	// to get the correct result
	fmt.Println("\nAverage Weight = ", avg)
}

// go build -o main
// ./main
// go run main.go
