package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	JSONFile, err := os.Open("sample/contacts.json")
	if err != nil {
		fmt.Println(err)
	}
	defer JSONFile.Close()

	byteValue, _ := ioutil.ReadAll(JSONFile)

	var rawContacts []StringMap
	json.Unmarshal([]byte(byteValue), &rawContacts)

	contacts := MergeContact(rawContacts)

	for _, contact := range contacts {
		fmt.Println(contact)
	}
}
