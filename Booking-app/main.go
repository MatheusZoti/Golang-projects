package main

import "fmt"

func main() {
	var userName string
	var userTickets int
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets


	//bloco principal
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	//descobrir o tipo usando %T no printf()
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)


	// ask user for their name
	fmt.Printf("What is your name?\n")
	fmt.Scan(&userName)



	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
}