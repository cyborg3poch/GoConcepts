package main

import (
	"fmt"
)

type roles map[string]bool

type User struct {
	Name     string
	Password string
	Roles    roles
}

func main() {

	user := User{
		Name: "Shivank", Password: "12345", Roles: roles{"admin": true},
	}

	// passing pointer to struct becuase when passing value type (i.e. array structs)  copy of variable are sent
	updater(&user)

	fmt.Printf("%#v", user)

}

func updater(user *User) {
	user.Name = "Pandey"
	user.Password = "Chaubey"
	user.Roles = roles{"admin": false}

}
