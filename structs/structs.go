package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Your first name: ")
	userLastName := getUserData("Your last name: ")
	userBirthDate := getUserData("Your birth date (DD/MM/YYYY): ")

	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	var admin user.Admin
	admin = user.NewAdmin("xx@gmsil.aom", "12345")
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
