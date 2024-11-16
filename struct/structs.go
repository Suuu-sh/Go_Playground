package main

import (
	"fmt"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	// createdAt time.Time
}

func newUser(userFirstName, userLastName, userBirthDate string) User {
	return User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
	}
}
func main() {
	userFirstName := getUserData("please")
	userLastName := getUserData("please")
	userBirthDate := getUserData("please")

	appUser := newUser(userFirstName, userLastName, userBirthDate)

	appUser.outPutUserDetails2()
	appUser.clearUserName()
	appUser.outPutUserDetails2()

}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func (u User) outPutUserDetails2() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
