package main

import (
	"fmt"
)

func main() {
	nit := User{"Nithin", "nithi@dev.com", true, 25}

	nit.GetStatus()

	// updated only in the method's instance
	nit.EditMail("nithin@mail.com")
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	activeStatus := ""
	if u.Status {
		activeStatus = "active"
	} else {
		activeStatus = "inactive"
	}
	fmt.Println("User is", activeStatus)
}

func (u User) EditMail(newMail string) {
	u.Email = newMail
	fmt.Println("Email updated", u.Email)
}
