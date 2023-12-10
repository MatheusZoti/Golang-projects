package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available,", conferenceTickets, remainingTickets)


	var userName string

	// ask user for their name
	userName = "Tom"
	fmt.Println(userName)
}