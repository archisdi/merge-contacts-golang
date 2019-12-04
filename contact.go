package main

// StringMap exported
type StringMap map[string]interface{}

// Contact exported
type Contact struct {
	fullname   string
	email      map[string]string
	phone      map[string]string
	address    map[string]string
	otherField map[string]string
}

func joinmap(m1 map[string]interface{}, m2 map[string]string) map[string]string {
	for ia, va := range m1 {
		m2[ia] = va.(string)
	}
	return m2
}

func (c *Contact) merge(rawContact StringMap) {
	for k, v := range rawContact {
		contact := (*c)

		if k == "fullname" {
			continue
		} else if k == "email" {
			c.email = joinmap(v.(map[string]interface{}), contact.email)
		} else if k == "address" {
			c.address = joinmap(v.(map[string]interface{}), contact.address)
		} else if k == "phone" {
			c.phone = joinmap(v.(map[string]interface{}), contact.phone)
		} else {
			c.otherField = joinmap(v.(map[string]interface{}), contact.otherField)
		}

	}
}

// MergeContact exported
func MergeContact(rawContacts []StringMap) []Contact {
	mergedContacts := map[string]Contact{}

	for i := 0; i < len(rawContacts); i++ {
		currentContact := rawContacts[i]
		_, keyExist := currentContact["fullname"]
		if keyExist {
			key := currentContact["fullname"].(string)
			_, alreadyExist := mergedContacts[key]

			if alreadyExist {
				existingContact := mergedContacts[key]
				existingContact.merge(currentContact)
			} else {
				newContact := Contact{
					fullname:   key,
					email:      map[string]string{},
					address:    map[string]string{},
					phone:      map[string]string{},
					otherField: map[string]string{},
				}
				newContact.merge(currentContact)
				mergedContacts[key] = newContact
			}
		}
	}

	// transform into array
	var contacts []Contact
	for _, v := range mergedContacts {
		contacts = append(contacts, v)
	}

	return contacts
}
