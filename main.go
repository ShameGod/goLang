package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const theaterName = "cinemaKejji"
const theaterTicekets = 50

var bookings = make([]UserData, 0)
var remainingTickets uint = theaterTicekets

type UserData struct {
	firstName string
	lastName  string
	email     string
	nbTickets int
}

func main() {

	greetUser()
	var firstName string
	var lastName string
	var email string
	var nbTickets int
	for {
		firstNames := []string{}

		firstName, lastName, email, nbTickets = getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUser(firstName, lastName, email, nbTickets)

		if !isValidTicketNumber {
			fmt.Printf("Not enough available tickets, you can book a maximum of %v", remainingTickets)
			continue
		}

		if !isValidName {
			fmt.Printf("the first or last name you entered is in invalide")
			continue
		}

		if !isValidEmail {
			fmt.Printf("the email you entre %v is invalid", email)
			continue
		}
		var userData = UserData{
			firstName: firstName,
			lastName:  lastName,
			email:     email,
			nbTickets: nbTickets,
		}
		bookTicket(userData)
		printFirstNames(bookings, firstNames)
		go sendEmail(userData)
		fmt.Println("Thank you for booking a ticket at our conferance")
		if remainingTickets == 0 {
			fmt.Println("Our Conference is booked out, Sowwy !")
			break
		}
	}

}

func greetUser() {
	fmt.Println("Welcome to ", theaterName, " !")
	fmt.Printf("we have total of %v places and %v avaiable\n", theaterTicekets, remainingTickets)
}

func printFirstNames(bookings []UserData, firstNames []string) {
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("the first names are %v \n", firstNames)
}

func validateUser(firstName string, lastName string, email string, nbTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := nbTickets < int(remainingTickets) && nbTickets > 0
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(userData UserData) {
	remainingTickets = remainingTickets - uint(userData.nbTickets)
	bookings = append(bookings, userData)
	fmt.Printf("thank you %v for booking %v tickets you ll receive an email at %v \n",
		userData.firstName, userData.nbTickets, userData.email)
	fmt.Printf("%v are remaining for the conference \n", remainingTickets)
}

func sendEmail(user UserData) {
	time.Sleep(1 * time.Second)
	fileName := fmt.Sprintf("bill of %v %v", user.firstName, user.lastName)
	err := ioutil.WriteFile(fileName, []byte(user.email), 0777)
	fmt.Print(err)
	fmt.Println("#####################################")
	fmt.Printf("Sending the email to the user %v %v", user.firstName, user.lastName)
	fmt.Println("#####################################")
}
