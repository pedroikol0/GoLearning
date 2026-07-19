package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u User) outputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) User {
	return User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser User

	appUser = newUser(userFirstName, userLastName, userBirthdate)

	appUser = User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	appUser.outputUserData()
	appUser.clearUserName()
	appUser.outputUserData()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
