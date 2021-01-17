package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

var u = User{
	ID:        45,
	FirstName: "Katie",
	LastName:  "Jellyfish",
	Email:     "katie@jelly.com",
}

func updateEmail(user *User) {
	user.Email = "k@jelly.com"
}

func main() {
	fmt.Println(u)
	updateEmail(&u)
	fmt.Println(u)
}