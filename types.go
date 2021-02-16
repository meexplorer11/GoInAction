package main

import "fmt"

type user struct {
	firstname string
	lastname  string
}

type role struct {
	usr  user
	role string
}

func main() {
	var u user
	u = user{}
	u.firstname = "John"
	u.lastname = "Doe"

	admin := role{
		usr:  u,
		role: "admin",
	}

	fmt.Println(admin)

	admin.printRole()
	printUser(u)

	u.updateFirstName("John K")
	printUser(u)
}

func (r role) printRole() {
	fmt.Println(r.role)
}

func printUser(u user) {
	fmt.Println(u)
}

func (u *user) updateFirstName(newName string) {
	u.firstname = newName
}
