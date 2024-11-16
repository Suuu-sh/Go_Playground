package user

import (
	"errors"
	"fmt"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	// createdAt time.Time
}

func NewUser(userFirstName, userLastName, userBirthDate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		return nil, errors.New("required")
	}
	return &User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
	}, nil
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func (u *User) OutPutUserDetails2() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
