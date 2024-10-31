package main

import (
	"fmt"
	"strings"
)

func main() {

	// Array declaration syntax
	// var bookings [50]string

	// Arrays can also have some initial values
	// var bookings = [50]string{"Sangam Kumar", "Rishi"}

	// Slices follow the same syntax for values assignment
	// var bookings = []string{"Sangam Kumar", "Rishi"}

	// It's better to use Slices than using Arrays
	// because Arrays have a fixed size
	// and Slices can grow and shrink
	var bookings []string

	// Variable declaration
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 // Variable declaration and assignment with type

	for {

		// Print statements (Printf means formatted print)
		fmt.Printf("Welcome to %v booking application\n", conferenceName)
		fmt.Printf("Total tickets: %v\n", conferenceTickets)
		fmt.Printf("Remaining tickets: %v\n", remainingTickets)

		// Println will print on the same line and give a line break
		fmt.Println("")

		// User input logic
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Printf("Enter your first name: \n")
		fmt.Scan(&firstName) // Scan can be used to accept inputs

		fmt.Printf("Enter your last name: \n")
		fmt.Scan(&lastName)

		fmt.Printf("Enter your email address: \n")
		fmt.Scan(&email)

		fmt.Printf("Enter number of tickets: \n")
		fmt.Scan(&userTickets)

		if (userTickets > remainingTickets) || (userTickets == 0) {
			fmt.Printf("We only have %v tickets remaining\n", remainingTickets)
			fmt.Printf("Please try again\n")
			continue // continue will skip the rest of the code and restart the loop
		}

		remainingTickets = remainingTickets - userTickets

		// This is the normal way of assigning values to an array
		// bookings[0] = firstName + " " + lastName

		// However, Golang has an append method as well
		// this method will add values to the end of an array
		// so we don't have to worry about the size of the array (or index of the last array item)
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("%v booked %v tickets\n", bookings[0], userTickets)
		fmt.Printf("You will receive your tickets at %v\n", email)
		fmt.Printf("%v tickets left\n", remainingTickets)
		fmt.Printf("Received bookings from %v people till now\n", len(bookings)) // len is used to get the length of an array

		for i, name := range bookings {
			fmt.Printf("Booking %d: %s\n", i, strings.Fields(name)[0])
		}

		if remainingTickets == 0 {
			fmt.Printf("All tickets have been booked. Come back next year")
			break // break will exit the loop
		}
	}

}
