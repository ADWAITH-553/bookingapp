package main

import "fmt"

func main() {
	const conferenceTickets uint = 50 //dont change
	var remainingTickets uint = 50
	conferenceName := "Go Conference"
	fmt.Printf("Welcome to %v our booking app\n", conferenceName)
	fmt.Println("we have total of ", conferenceTickets, "remaining", remainingTickets)
	fmt.Println("Book ur tickets for conference")
	var firstName string
	var lastName string
	var email string
	var userTickets int
	fmt.Println("enter ur first name")
	fmt.Scan(&firstName)
	fmt.Println("enter ur last name")
	fmt.Scan(&lastName)
	fmt.Println("enter ur email address")
	fmt.Scan(&email)
	fmt.Println("enter the number of tickets u need to book")
	fmt.Scan(&userTickets)
	fmt.Printf("Thankyou %v %v for booking %v tickets you will receive a conf mail on %v", firstName, lastName, userTickets, email)
}
