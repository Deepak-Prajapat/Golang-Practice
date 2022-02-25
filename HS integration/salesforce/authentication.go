package salesforce

import (
	"encoding/json"
	"hsintegration/api"
	"io"
)

type Authentication struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func AuthorizationToken() *Authentication {
	conf := getSfConfig()
	sfapi := &api.SfAuthentication{Config: &api.SFConfig{
		ClientID:      conf.ClientID,
		ClientSecret:  conf.ClientSecret,
		GrantType:     conf.GrantType,
		Password:      conf.Password,
		SecurityToken: conf.SecurityToken,
		Username:      conf.Username,
	}}

	response, err := sfapi.Request().Send()
	if err != nil {
		return nil
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil
	}

	var auth Authentication
	json.Unmarshal(bytes, &auth)
	return &auth
}
