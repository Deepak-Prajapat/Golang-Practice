package hubspot

import (
	"encoding/json"
	"errors"
	"hsintegration/api"
	"hsintegration/service"
	u "hsintegration/utility"
	"io"
	"net/http"
)

type DB interface {
	VerifyAPIKey(key string) error
}
type db struct{}

// DBService will help to mock methods
var DBService DB = db{}

func (d db) VerifyAPIKey(key string) error {
	if service.VerifyAPI(key) {
		return nil
	}
	return errors.New("api key is not available in database")
}

func GetAllContacts(w http.ResponseWriter, r *http.Request) {
	//When api key is not provided
	r.ParseForm()
	apiKey := r.Form.Get("hapikey")
	if u.IsBlank(apiKey) {
		u.R.Text(w, http.StatusBadRequest, "Please provide an api key")
		return
	}

	// If api is not found in database
	err := DBService.VerifyAPIKey(apiKey)
	if err != nil {
		u.R.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	contacts := api.AllContacts{APIKey: &apiKey}
	response, err := contacts.Request().Send()
	if err != nil {
		u.R.Text(w, http.StatusBadGateway, "error while getting data from api")

		return
	}
	js, _ := io.ReadAll(response.Body)
	u.R.Data(w, http.StatusOK, js)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	apiKey := r.Form.Get("hapikey")
	if u.IsBlank(apiKey) {
		//w.Write([]byte("please provide an api"))
		u.R.Text(w, http.StatusBadGateway, "api key is missing")

		return
	}

	// If api is not found in database
	err := DBService.VerifyAPIKey(apiKey)
	if err != nil {
		u.R.Text(w, http.StatusBadGateway, err.Error())
		return
	}

	body, err := io.ReadAll(r.Body)
	contact := &api.HSContactCreate{
		APIKey: apiKey,
		Body:   body,
	}

	response, err := contact.Request().Send()
	defer response.Body.Close()

	js, _ := io.ReadAll(response.Body)

	if !isResponseOK(response.StatusCode) {
		var mapp map[string]interface{}
		_ = json.Unmarshal(js, &mapp)

		er := mapp["errors"]
		iSlice := u.InterfaceSlice(er)[0]
		u.R.JSON(w, http.StatusAlreadyReported, iSlice)
		return
	}

	if err != nil {
		u.R.Text(w, http.StatusBadRequest, "error while getting response")
	}
	//Write JSON Response
	var jsonMap map[string]interface{}
	json.Unmarshal(js, &jsonMap)
	u.R.JSON(w, http.StatusOK, jsonMap)
}

func isResponseOK(resCode interface{}) bool {
	if resCode == 200 {
		return true
	} else if resCode == 201 {
		return true
	}
	return false
}
