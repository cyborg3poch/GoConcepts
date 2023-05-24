package main

import (
	"encoding/json"
	"fmt"
)

type roles map[string]bool

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Roles    roles  `json:role,omitempty`
}

func main() {

	user := []User{
		{Name: "Shivank", Password: "12345", Roles: roles{"admin": true}},
	}

	op, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(op))

}
