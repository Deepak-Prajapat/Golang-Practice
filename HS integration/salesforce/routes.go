package salesforce

import (
	"encoding/json"
	"hsintegration/api"
	u "hsintegration/utility"
	"io"
	"net/http"
	"strings"
)

func GetAccountInfo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		u.R.Text(w, http.StatusBadRequest, "error in parsing the request")
		return
	}

	accID := r.Form.Get("id")
	if u.IsBlank(accID) {
		u.R.Text(w, http.StatusBadGateway, "account id is missing")
		return
	}

	if len(accID) < 15 || len(accID) > 18 {
		u.R.Text(w, http.StatusBadGateway, "please provide a valid ID for account\nmsg: ID's length should between 15 to 18")
		return
	}

	auth := AuthorizationToken()

	account := &api.GetAccount{
		Authorization: auth.AccessToken,
		AccountID:     accID,
	}

	response, err := account.Request().Send()
	if err != nil {
		u.R.Text(w, http.StatusBadGateway, err.Error())
		return
	}

	bytes, _ := io.ReadAll(response.Body)

	//Check if Error occurs from api
	var errResponse []ErrorResponse
	json.Unmarshal(bytes, &errResponse)
	if len(errResponse) > 0 {
		u.R.Text(w, http.StatusConflict, errResponse[0].ErrMessage[:strings.Index(errResponse[0].ErrMessage, "\n")])
		return
	}

	var resp GetAccountResponse
	json.Unmarshal(bytes, &resp)
	u.R.JSON(w, http.StatusOK, resp)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		u.R.Text(w, http.StatusBadRequest, "error in parsing the request")
		return
	}

	auth := AuthorizationToken()
	bytes, _ := io.ReadAll(r.Body)

	account := &api.UpdateAccount{
		Authorization: auth.AccessToken,
		Body:          bytes,
	}

	response, _ := account.Request().Send()
	bytes, _ = io.ReadAll(response.Body)

	//respResult := []byte(string(bytes)[1 : len(string(bytes))-1])

	//Check if Error occurs
	var errResponse []ErrorResponse
	json.Unmarshal(bytes, &errResponse)
	if len(errResponse) > 0 {
		u.R.Text(w, http.StatusConflict, errResponse[0].ErrMessage[:strings.Index(errResponse[0].ErrMessage, "\n")])
		return
	}

	var resp GetAccountResponse
	json.Unmarshal(bytes, &resp)
	u.R.JSON(w, http.StatusOK, resp)
}
