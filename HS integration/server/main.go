package server

import (
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	Config *Config
	DB     *gorm.DB
	log    *log.Entry
	r      *render.Render
}

type Config struct {

	// database
	DbHost           *string
	DbName           *string
	DbUser           *string
	DbPass           *string
	DbPort           *int
	DbMaxConnections *int
	DbVerboseLogs    *bool

	// Localhost PORT
	WebPort *int

	//HubSpot
	HubSpotURL *string
	APIKey     *string

	//Salesforce Authentication
	SfTokenURL      *string
	SfGrantType     *string
	SfClientID      *string
	SfClientSecret  *string
	SfUsername      *string
	SfPassword      *string
	SfSecurityToken *string
	SfInstanceURL   *string

	//Salesforce API
	InstanceURL *string
	CalloutsURL *string
	BearerToken *string
}

// NewServer returns a new instance of Server
func NewServer() *Server {
	s := &Server{
		Config: &Config{},
	}

	return s
}
func (s *Server) setupRenderer() {
	s.r = render.New()
}
func (s *Server) setupAccess() {
	s.setupDBAccess()
	s.setupMigrations()
	s.setupRenderer()
	s.setupUtility()
	s.setupSF()
}

func (s *Server) setupCloseHandler() {
	c := make(chan os.Signal, 2)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Warn("process has been interrupted")
		s.Close()
		os.Exit(0)
	}()
}

func (s *Server) Close() {
	log.Print("shutting down")
}

// Run starts the server
func (s *Server) Run() error {
	defer s.Close()

	//To stop serer
	s.setupCloseHandler()

	// configure database connections
	if err := s.dbConnect(); err != nil {
		log.Print(err.Error())
		return err
	}

	s.setupAccess()
	s.RunWebServer()

	return nil
}
