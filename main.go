package main

import "fmt"

// StringMap exported
type StringMap map[string]string

// Contact exported
type Contact struct {
	fullname   string
	email      StringMap
	phone      StringMap
	address    StringMap
	otherField StringMap
}

func main() {
	contact := Contact{
		fullname: "archie isdiningrat",
		email: StringMap{
			"personal": "archie.isdiningrat@gmail.com",
		},
	}

	fmt.Println(contact)
}
