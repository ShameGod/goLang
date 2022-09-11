package main

import (
	"fmt"
)

func getUserInput() (string, string, string, int) {

	var firstName string
	var lastName string
	var email string
	var nbTickets int

	fmt.Println("enter your firstName : ")
	fmt.Scan(&firstName)

	fmt.Println("enter your lastName : ")
	fmt.Scan(&lastName)

	fmt.Println("enter your emailAddress : ")
	fmt.Scan(&email)

	fmt.Println("enter the number of tickets that you want  : ")
	fmt.Scan(&nbTickets)

	return firstName, lastName, email, nbTickets
}
