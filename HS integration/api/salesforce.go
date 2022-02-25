package api

import (
	"fmt"
	"net/http"
)

type SFConfig struct {
	URL           string
	ClientID      string
	ClientSecret  string
	GrantType     string
	Password      string
	SecurityToken string
	Username      string
}
type SfAuthentication struct {
	Config  *SFConfig
	request *Request
}

func (sfapi *SfAuthentication) Request() *SfAuthentication {
	//grant_type=password&username=golang@concret.io&password=gopractice1ISv9Qophfz2JWOu67ojDofou
	uriParams := fmt.Sprintf("?grant_type=%v&client_id=%v&client_secret=%v&username=%v&password=%v%v", sfapi.Config.GrantType, sfapi.Config.ClientID, sfapi.Config.ClientSecret, sfapi.Config.Username, sfapi.Config.Password, sfapi.Config.SecurityToken)

	sfapi.request = &Request{
		Method:    "POST",
		For:       "SF_AUTH",
		URIParams: uriParams,
	}
	return sfapi
}

func (sf *SfAuthentication) Send() (*http.Response, error) {
	return sf.request.Send()
}

type GetAccount struct {
	Authorization string
	URL           string
	AccountID     string
	request       *Request
}

func (acc *GetAccount) Request() *GetAccount {
	acc.request = &Request{
		Method:              "GET",
		For:                 "SF_INSTANCE",
		ResourcePath:        "/services/apexrest/Account",
		URIParams:           "?id=" + acc.AccountID,
		BearerAuthorization: acc.Authorization,
	}
	return acc
}

func (acc *GetAccount) Send() (*http.Response, error) {
	return acc.request.Send()
}

// UpdateAccount ...
type UpdateAccount struct {
	Authorization string
	URL           string
	AccountID     string
	Body          []byte
	request       *Request
}

func (acc *UpdateAccount) Request() *UpdateAccount {
	acc.request = &Request{
		Method:              "PUT",
		For:                 "SF_INSTANCE",
		ResourcePath:        "/services/apexrest/Account",
		BearerAuthorization: acc.Authorization,
		Body:                acc.Body,
	}
	return acc
}

func (acc *UpdateAccount) Send() (*http.Response, error) {
	return acc.request.Send()
}
