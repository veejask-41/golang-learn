package main

import "fmt"

func main() {
	var confName = "SingSong"
	const confTickets = 100
	var remainingTickets = 100

	fmt.Printf("Welcome to %v conference booking app\n", confName)
	fmt.Printf("We had %v tickets, and only %v tickets are remaining\n", confTickets, remainingTickets)

	var userName string

	fmt.Println("Please enter your name")
	fmt.Scanln(&userName)

	fmt.Printf("Hello %v, how many tickets would you like to book?\n", userName)

	var bookings []string

	fmt.Println(bookings)

	bookings = append(bookings, userName)

	fmt.Println(bookings)
}
