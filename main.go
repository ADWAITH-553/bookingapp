package main

import "fmt"

func main() {
	const conferenceTickets = 50 //dont change
	var remainingTickets = 50
	var conferenceName = "Go Conference"
	fmt.Printf("Welcome to %v our booking app\n", conferenceName)
	fmt.Println("we have total of ", conferenceTickets, "remaining", remainingTickets)
	fmt.Println("Book ur tickets for conference")

}
