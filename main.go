package main

import (
	"booking_app/helper"
	"fmt"
	"strings"
)

const conferenceTickets uint = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetuser()
	//fmt.Printf("Welcome to %v our booking app\n", conferenceName)

	for {

		firstName, lastName, email, userTickets := getUserInputs()
		//boolean
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			firstNames := Fnames()
			fmt.Printf("The first names of our bookings  are:%v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("our conference is bokked out. come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Printf("your inputs are invalid tryagain firstname or lastname entered is two short\n")
			}
			if !isValidEmail {
				fmt.Println("email entered is invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("ticket number entered is invalid")
			}
		}

	}
}

func greetuser() {
	fmt.Printf("Welcome to our %v booking application", conferenceName)
	fmt.Println("we have total of ", conferenceTickets, "remaining", remainingTickets)
	fmt.Println("Book ur tickets for conference")
}
func Fnames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	//
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("enter ur first name")
	fmt.Scan(&firstName)

	fmt.Println("enter ur last name")
	fmt.Scan(&lastName)

	fmt.Println("enter ur email address")
	fmt.Scan(&email)

	fmt.Println("enter the number of tickets u need to book")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thankyou %v %v for booking %v tickets you will receive a conf mail on %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("The first names of our bookings  are:%v\n", bookings)
}
