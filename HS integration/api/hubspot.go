package api

import (
	"fmt"
	"net/http"
)

// AllContacts stores HubSpot Contacts
type AllContacts struct {
	APIKey  *string
	request *Request
}

// Request ...
func (contacts *AllContacts) Request() *AllContacts {
	contacts.request = &Request{
		ResourcePath: fmt.Sprintf("lists/all/contacts/all"),
		Method:       "GET",
		For:          "HS",
		ContentType:  "application/json",
		ApiKey:       map[string]string{"hapikey": *contacts.APIKey},
	}
	return contacts
}

// Send ...
func (contacts *AllContacts) Send() (*http.Response, error) {
	return contacts.request.Send()
}

// HSContactCreate ...
type HSContactCreate struct {
	APIKey  string
	Body    []byte
	request *Request
}

// Request ...
func (contact *HSContactCreate) Request() *HSContactCreate {
	contact.request = &Request{
		ResourcePath: fmt.Sprintf("contact"),
		Method:       "POST",
		For:          "HS",
		ContentType:  "application/json",
		ApiKey:       map[string]string{"hapikey": contact.APIKey},
		Body:         contact.Body,
	}
	return contact
}

// Send ...
func (contact *HSContactCreate) Send() (*http.Response, error) {
	return contact.request.Send()
}
