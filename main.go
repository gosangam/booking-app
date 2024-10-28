package main

import ("fmt")

func main() {
	fmt.Print("Hello World\n")

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total tickets: %v\n", conferenceTickets)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
	fmt.Println("")

	// User input logic
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Printf("Enter your first name: \n")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: \n")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: \n")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: \n")
	fmt.Scan(&userTickets)

	fmt.Printf("%v %v booked %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive your tickets at %v\n", email)
	fmt.Printf("%v tickets left", conferenceTickets - userTickets)
}	