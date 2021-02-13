package main

import (
	"fmt"
)

func main() {

	// Average from 2 inputs
	fmt.Print("Enter your age: ")
	var age int
	fmt.Scanln(&age)
	// fmt.Print(input)

	// number, _ := strconv.ParseUint(input, 10, 32)
	// int1 := int(number) //Convert uint64 To int

	fmt.Print("Enter the day of the month you were born: ")
	var day string
	fmt.Scanln(&day)

	fmt.Print("Enter the month you were born: ")
	var month int
	fmt.Scanln(&month)

	fmt.Print("Enter the day of the year you were born: ")
	var year int
	fmt.Scanln(&year)
	// fmt.Print(input2)

	// number2, _ := strconv.ParseUint(year, 10, 32)
	// int2 := int(number2) //Convert uint64 To int
	// // declaring an array of values
	// arra := []int{1, 2, 3, 4}
	arra := []int{age, year}

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
	// avg := (float64(sum)) / (float64(n))

	// typecast all values to float
	// to get the correct result
	fmt.Println("\nThe year is: ", sum)
}

// go build -o main
// ./main
// go run main.go
