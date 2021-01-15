package main

import "fmt"

// User is a user type
type User struct {
ID					       int
FirstName, LastName, Email string
}

// Group is a group type
type Group struct {
	role            string
	users           []User
	newestUser      User
	spaceAvailable  bool
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
	return desc
}

func describeGroup(g Group) string {
	if len(g.users) >= 2 {
		g.spaceAvailable = false
	}
	desc := fmt.Sprintf("The %s user group has %d users. Newest user:  %s, Accepting New Users:  %t", g.role, len(g.users), g.newestUser.FirstName, g.spaceAvailable)
	return desc
}

func main() {
	u1 := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 2, FirstName: "Humphrey", LastName: "Bogart", Email: "humphrey.bogart@gmail.com"}

	g := Group{role: "admin", users: []User{u1, u2}, newestUser: u2, spaceAvailable: true}

	userDescription := describeUser(u1)
	groupDescription := describeGroup(g)

	fmt.Println(userDescription)
	fmt.Println(groupDescription)

	fmt.Println(g)
}
