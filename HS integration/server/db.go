package server

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// dbConnect attempts to connect to the provided database and stores the connection in DB and DBNA
func (s *Server) dbConnect() error {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", *s.Config.DbHost, *s.Config.DbPort, *s.Config.DbUser, *s.Config.DbPass, *s.Config.DbName)
	config = fmt.Sprintf("%s sslmode=disable", config)

	db, err := gorm.Open(postgres.Open(config))
	if err != nil {
		return err
	}
	s.DB = db

	if err != nil {
		return err
	}
	return nil
}
