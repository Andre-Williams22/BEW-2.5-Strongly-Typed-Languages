package main

import (
	"fmt"
	"strconv"
)

type House struct {
	city          string
	address       string
	price         int
	numberOfRooms int
}

func main() {
	// instantiate House struct
	var houses []House

	// Take inputs
	fmt.Print("Enter your how many houses you'd like to enter for: ")
	var num_times string
	fmt.Scanln(&num_times)
	// fmt.Print(input)

	nums, _ := strconv.ParseUint(num_times, 10, 32)
	n := int(nums) //Convert uint64 To int

	i := 0

	max_times := n

	for i < max_times {

		// Take new input
		fmt.Print("Enter the address house: ")
		var address string
		fmt.Scanln(&address)

		// Take new input
		fmt.Print("Enter the city: ")
		var city string
		fmt.Scanln(&city)

		// Take new input
		fmt.Print("Enter the number of rooms: ")
		var rooms int
		fmt.Scanln(&rooms)

		// Take new input
		fmt.Print("Enter the price: ")
		var price int
		fmt.Scanln(&price)

		// create new house variable
		house := House{
			numberOfRooms: rooms,
			city:          city,
			address:       address,
			price:         price,
		}
		// add house to the struct
		houses = append(houses, house)

		// print space
		fmt.Print("\n ")

		// increment counter
		i += 1

	}
	// loop through houses
	for home := range houses {
		// print content
		fmt.Printf("address: %s\t city: %s\t %d rooms\t price: $%d\n ", houses[home].address, houses[home].city, houses[home].numberOfRooms, houses[home].price)

	}

}

// go build -o main
// ./main
// go run main.go
