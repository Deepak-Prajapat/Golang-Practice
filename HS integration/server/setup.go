package server

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"hsintegration/dbaccess"
	"hsintegration/resourcehandler"
	"hsintegration/salesforce"
	"hsintegration/utility"
)

func (s *Server) setupDBAccess() {
	dbaccess.SetupService(s.DB, s.log)
	//hubspot.SetupDB(s)
}

func (s *Server) setupMigrations() {
	dbaccess.SetupMigrations()
}

func (s *Server) setupUtility() {
	utility.SetupService(s.log, s.r)
}

func (s *Server) setupSF() {
	salesforce.SetupSFConfig(s.Config.SfPassword, s.Config.SfGrantType, s.Config.SfClientID, s.Config.SfClientSecret, s.Config.SfSecurityToken, s.Config.SfUsername)
}

func (s *Server) setupPublicRoutes(r *mux.Router) {
	//HubSpot
	hubspotRoute := r.PathPrefix("/hubspot").Subrouter()
	resourcehandler.AddHubSpotRoutes(hubspotRoute)

	//Salesforce
	sfRoutes := r.PathPrefix("/sf").Subrouter()
	resourcehandler.AddSfRoutes(sfRoutes)
}

// RunWebServer configures and starts the API server
func (s *Server) RunWebServer() {
	apiServer := negroni.New()
	router := mux.NewRouter()

	//Setup Public Routes
	s.setupPublicRoutes(router)

	apiServer.UseHandler(router)
	log.WithField("port", *s.Config.WebPort).Info("starting web server")
	apiServer.Run(fmt.Sprintf(":%d", *s.Config.WebPort))
}
