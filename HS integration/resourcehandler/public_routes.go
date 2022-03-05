package resourcehandler

import (
	"github.com/gorilla/mux"
	"hsintegration/hubspot"
	"hsintegration/salesforce"
	"net/http"
)

func AddHubSpotRoutes(router *mux.Router) {
	ContactRoutes(router)
}

func AddSfRoutes(r *mux.Router) {
	AccountRoutes(r)
}

func ContactRoutes(router *mux.Router) {
	//https://api.hubapi.com/contacts/v1/lists/all/contacts/all?hapikey=demo&count=2
	pathPrefix := "/contacts"
	router.HandleFunc(pathPrefix, hubspot.GetAllContacts).Methods(http.MethodGet)

	pathPrefix = "/contact"
	router.HandleFunc(pathPrefix, hubspot.CreateContact).Methods(http.MethodPost)

}

func AccountRoutes(r *mux.Router) {
	pathPrefix := "/account"
	r.HandleFunc(pathPrefix, salesforce.GetAccountInfo).Methods(http.MethodGet)
	r.HandleFunc(pathPrefix, salesforce.UpdateAccount).Methods(http.MethodPut)
}
