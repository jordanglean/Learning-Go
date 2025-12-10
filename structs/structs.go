package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := user.GetUserData("Please enter your first name: ")
	userLastName := user.GetUserData("Please enter your last name: ")
	userBirthDate := user.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Print("Error creating new user!")
		return
	}

	admin, err := user.NewAdmin("jordan@test.com", "password")
	if err != nil {
		fmt.Println("Error creating new admin!")
	}

	admin.OutputAdminDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}
