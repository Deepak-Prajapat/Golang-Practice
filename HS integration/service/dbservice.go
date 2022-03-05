package service

import (
	"hsintegration/dbaccess"
	u "hsintegration/utility"
)

func VerifyAPI(key string) bool {
	client, err := dbaccess.CheckForAPI(key)
	if err != nil {
		return false
	}
	if u.IsBlank(client.APIKey) {
		return false
	}
	return true
}
