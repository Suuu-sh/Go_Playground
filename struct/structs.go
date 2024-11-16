package main

import (
	"fmt"

	"example.com/practice-struct/user"
)

func main() {
	userFirstName := user.GetUserData("please")
	userLastName := user.GetUserData("please")
	userBirthDate := user.GetUserData("please")

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthDate)
	admin := user.NewAdmin("test", "test")

	admin.User.OutPutUserDetails2()

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutPutUserDetails2()
	appUser.ClearUserName()
	appUser.OutPutUserDetails2()
}
