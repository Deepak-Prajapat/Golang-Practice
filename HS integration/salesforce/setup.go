package salesforce

import "fmt"

type SFConfig struct {
	ClientID      string
	ClientSecret  string
	GrantType     string
	Password      string
	SecurityToken string
	Username      string
}

var sf = SFConfig{}

func SetupSFConfig(password *string, grant *string, cID *string, cSecret *string, secToken *string, user *string) {
	fmt.Println("sf at beginnning:", sf)
	sf.Password = *password
	sf.GrantType = *grant
	sf.ClientID = *cID
	sf.ClientSecret = *cSecret
	sf.SecurityToken = *secToken
	sf.Username = *user
}

func getSfConfig() SFConfig {
	return sf
}
