package main

import (
	"booking-app/helper"
	"fmt" // pkg.go.dev/fmt
	"sync"
	"time"
)

//2:37:20
/*package level variables*/
var conferenceName = "Go Conference"
var remianingTickets uint = conferenceTickets

// var bookings = make([]map[string]string, 0) // var bookings = [50]string{}
var bookings = make([]UserData, 0) // var bookings = [50]string{}
const conferenceTickets = 50

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	//fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remianingTickets, conferenceName)
	//entry point
	//greetUsers(conferenceName, conferenceTickets, remianingTickets)
	greetUsers()

	//for remianingTickets > 0 && len(bookings) < 50 {
	firstname, lastname, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidNumber := helper.ValidInputs(firstname, lastname, email, userTickets, remianingTickets)

	if isValidName && isValidEmail && isValidNumber {
		bookTickets(userTickets, firstname, lastname, email)

		wg.Add(1) //wait before creating new thread
		go sendTicket(userTickets, firstname, lastname, email)

		firstNames := getFirstNames()
		fmt.Printf("Thes first names of bookings are: %v\n", firstNames)

		if remianingTickets == 0 {
			//end program
			fmt.Println("Our Conference is booked out, come back next year")
			//break //stop loop
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name was too short")
		}
		if !isValidEmail {
			fmt.Println("The email you entered didn't contain @ sign")
		}
		if !isValidNumber {
			fmt.Println("please enter a valid number")
		}
		//fmt.Printf("There are only %v tickets left, so you can't book %v tickets\n", remianingTickets, userTickets)
		//continue //loop again from start of loop
	}
	wg.Wait()
	//}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceName, remianingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { //_ blink indentifier for where index is declare but not used
		//var names = strings.Fields(booking) //separate by space
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// capitalize the first letter so export function
func getUserInput() (string, string, string, uint) {
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
	return firstname, lastname, email, userTickets
}

func bookTickets(userTickets uint, firstname string, lastname string, email string) {
	remianingTickets = remianingTickets - userTickets
	//create a map for user
	//var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstname,
		lastName:        lastname,
		email:           email,
		numberOfTickets: userTickets,
	}
	// userData["firstName"] = firstname
	// userData["lastName"] = lastname
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	//slice in Go-> avstration of an array, variable-length
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receieve a confirmation email at %v\n", firstname, lastname, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remianingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("############################")
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("############################")
	wg.Done()
}
