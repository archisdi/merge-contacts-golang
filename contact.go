package main

// StringMap exported
type StringMap map[string]interface{}

// Contact exported
type Contact struct {
	fullname   string
	email      StringMap
	phone      StringMap
	address    StringMap
	otherField StringMap
}

// MergeContact exported
func MergeContact(contacts []StringMap) []Contact {
	mergedContacts := []Contact{}
	return mergedContacts
}
