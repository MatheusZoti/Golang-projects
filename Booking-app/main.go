package main

import (
	"fmt"
	"strconv"
	"strings"
)

// VARIAVEIS 'GLOBAIS' dentro desse package (PACKAGE LEVEL)
const conferenceTickets uint = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}



func main() {

	greetUsers( conferenceName, conferenceTickets, remainingTickets)

	for {
		// CAPTURA O INPUT DO USUÁRIO
		firstName, lastName, email, userTickets := getUserInput()

		// VALIDAÇÃO DO INPUT
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)


		if isValidEmail && isValidName && isValidTicketNumber {

			// Atualizar as variaveis: 1- Tickets que sobraram; 2- Array dos compradores
			bookTicket(userTickets, firstName, lastName, email)
		
			// FUNÇÃO DE PRINTAR OS PRIMEIROS NOMES NO ARRAY
			firstNames := getFirstNames()
			fmt.Printf("A lista dos nomes é essa: %v\n", firstNames)

			
			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your name is not valid.")
			}
			if !isValidEmail {
				fmt.Println("Your email is not valid.")
			}
			if !isValidTicketNumber {
				fmt.Println("Your Ticket number is not valid.")
			}
		}
	}
}



func greetUsers(conferenceName string, conferenceTickets, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
}
 

func getFirstNames() []string {
	firstNames := []string{}
	for _, bookingValue := range bookings {
		var names = strings.Fields(bookingValue)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}


func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("What is your first name?\n")
	fmt.Scan(&firstName)
	fmt.Printf("What is your last name?\n")
	fmt.Scan(&lastName)
	fmt.Printf("What is your email address:\n")
	fmt.Scan(&email)
	fmt.Printf("How many tickets do you want?\n")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}


func bookTicket(userTickets uint, firstName, lastName, email string) {
	remainingTickets = remainingTickets - userTickets

	// Create a map for each user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You'll receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}