package main

import "fmt" // pkg.go.dev/fmt

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remianingTickets uint = conferenceTickets
	var bookings [50]string // var bookings = [50]string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remianingTickets, conferenceName)
	//entry point
	fmt.Printf("Welcome to %v booking system\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remianingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstname string
	var lastname string
	var email string
	var userTickets uint

	//ask user for their information
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstname) //&go and find memory address
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastname)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter your number of tickets: ")
	fmt.Scan(&userTickets)

	remianingTickets = remianingTickets - userTickets
	bookings[0] = firstname + " " + lastname

	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("The first value array: %v\n", bookings[0])
	// fmt.Printf("Array type: %T\n", bookings)
	// fmt.Printf("Array length: %v\n", len(bookings))
	//slice in Go-> avstration of an array, variable-length
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receieve a confirmation email at %v\n", firstname, lastname, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remianingTickets, conferenceName)
}
