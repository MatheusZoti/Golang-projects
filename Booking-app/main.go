package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	bookings := []string{}


	//Mensagem inicial e contagem dos tickets disponiveis
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)


	//For LOOP
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		
		// Informações do usuário e compra dos tickets
		fmt.Printf("What is your first name?\n")
		fmt.Scan(&firstName)
		fmt.Printf("What is your last name?\n")
		fmt.Scan(&lastName)
		fmt.Printf("What is your email address:\n")
		fmt.Scan(&email)
		fmt.Printf("How many tickets do you want?\n")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTicketNumber {
			// Atualizar as variaveis: Tickets que sobraram; Array dos compradores
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)
		

			// Mensagem de Obrigado pela compra E quantos tickets sobraram após a mesma
			fmt.Printf("Thank you %v %v for booking %v tickets. You'll receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		
			firstNames := []string{}
			for _, bookingValue := range bookings {
				var names = strings.Fields(bookingValue)
				firstNames = append(firstNames, names[0])
			}


			// Lista de pessoas que compraram o ticket
			fmt.Printf("Esses são os nomes dos que compraram nossos tickets: %v\n", firstNames)

				if remainingTickets == 0 {
				// fechar a aplicação
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