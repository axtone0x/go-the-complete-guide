package main

import (
	"fmt"
	"example.com/go-structs-and-custom-types/user"
)

func main() {
	// userFirstName := getUserData("Please enter your first name: ")
	// userLastName := getUserData("Please enter your last name: ")
	// userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser *user.User

	// appUser = &user.User{
	// 	FirstName: "Isela",
	// 	LastName: "Miranda",
	// 	Birthdate: "",
	// } //-> The return type is a pointer

	//appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)

	// appUser, err := user.NewUser(
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// )

	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }

	admin := user.NewAdmin("oceballos@ucol.mx","spoons123")

	//Shorter notation
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }

	// appUser.OutputUserDetails()
	// appUser.ClearUserName()
	// appUser.OutputUserDetails()

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

}

func getUserData(promptText string) (value string) {
	fmt.Print(promptText)
	fmt.Scanln(&value)
	return
}
