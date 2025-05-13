package main

import (
	"fmt"
	"study-app/helpers"
)

func main() {
	var conferenceName = "GopherCon"
	const conferenceTickets = 100
	var remainingTickets = conferenceTickets
	var usersData = []map[string]string{}

	for remainingTickets > 0 {
		fmt.Println("Welcome to our conference booking application!")
		fmt.Println("Get your tickets here to attend the", conferenceName)
		fmt.Println("We have", remainingTickets, "tickets left!")

		var userFirstName string
		var userLastName string
		var userEmail string
		var userTickets int

		fmt.Print("Enter your name: ")
		fmt.Scanln(&userFirstName)

		fmt.Print("Enter your last name: ")
		fmt.Scanln(&userLastName)

		fmt.Print("Enter your email: ")
		fmt.Scanln(&userEmail)

		fmt.Print("Enter your number of tickets: ")
		fmt.Scanln(&userTickets)

		if !helpers.ValidateUserName(userFirstName) || !helpers.ValidateUserLastName(userLastName) || !helpers.ValidateUserEmail(userEmail) {
			fmt.Println("Invalid input, please try again!")
			continue
		}

		if userTickets < 0 || userTickets > remainingTickets {
			fmt.Printf("You can't book %d tickets, we have only %d tickets left!\n", userTickets, remainingTickets)
			continue
		}

		remainingTickets -= userTickets

		var userData = map[string]string{
			"FirstName": userFirstName,
			"LastName":  userLastName,
			"Email":     userEmail,
			"Tickets":   fmt.Sprint(userTickets),
		}
		usersData = append(usersData, userData)
	}

	fmt.Println("Sorry, all tickets are sold out!")

	for _, value := range usersData {
		fmt.Printf("Thanks %s for booking %s tickets for %s\n", value["FirstName"], value["Tickets"], conferenceName)
	}
	fmt.Println("Thank you for coming to the", conferenceName)
}
