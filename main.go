package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceTickets uint = 50 //dont change
	var remainingTickets uint = 50
	conferenceName := "Go Conference"
	fmt.Printf("Welcome to %v our booking app\n", conferenceName)
	fmt.Println("we have total of ", conferenceTickets, "remaining", remainingTickets)
	fmt.Println("Book ur tickets for conference")
	for {

		var bookings []string
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
		//boolean
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@gmail.com")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			//bookings[0] = firstName + "" + lastName
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("the whole slice %v\n", bookings)
			fmt.Printf("the single slice %v\n", bookings[0])
			fmt.Printf("Array type:%T\n", bookings)
			fmt.Printf("Array length:%v\n", len(bookings))
			fmt.Printf("Thankyou %v %v for booking %v tickets you will receive a conf mail on %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of our bookings  are:%v\n", bookings)
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
