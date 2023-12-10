package main

import "fmt"



func main() {
	bookings := []string{}
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets


	//Mensagem inicial e contagem dos tickets disponiveis
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

	// Informações do usuário e compra dos tickets
	fmt.Printf("What is your first name?\n")
	fmt.Scan(&firstName)
	fmt.Printf("What is your last name?\n")
	fmt.Scan(&lastName)
	fmt.Printf("What is your email address:\n")
	fmt.Scan(&email)
	fmt.Printf("How many tickets do you want?\n")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You'll receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("Essas são todas as pessoas que compraram nossos tickets: %v\n", bookings)
}