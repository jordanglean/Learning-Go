package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

// OutputUserDetails reciver
func (u *User) OutputUserDetails() {
	fmt.Println("First Name: ", u.firstName)
	fmt.Println("Last Name: ", u.lastName)
	fmt.Println("BirthDate: ", u.birthDate)
}

func (a *Admin) OutputAdminDetails() {
	fmt.Println("Email", a.email)
	fmt.Println("Password", a.password)
	fmt.Println("Admin User Details: ")
	fmt.Println("Admin First Name: ", a.firstName)
	fmt.Println("Admin Last Name: ", a.lastName)
	fmt.Println("Admin BirthDate: ", a.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstName, last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) (*Admin, error) {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "----",
			createdAt: time.Now(),
		},
	}, nil
}

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	_, _ = fmt.Scanln(&value)
	return value
}
