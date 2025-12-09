package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// reciver
func (u user) outputUserDetails() {
	fmt.Println("First Name", u.firstName)
	fmt.Println("Last Name", u.lastName)
	fmt.Println("Birthdata", u.birthDate)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}

	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	_, _ = fmt.Scan(&value)
	return value
}
