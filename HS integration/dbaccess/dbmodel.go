package dbaccess

import "gorm.io/gorm"
import log "github.com/sirupsen/logrus"

var s *Service

// Service ...
type Service struct {
	db  *gorm.DB
	log *log.Entry
}

func SetupService(db *gorm.DB, log *log.Entry) {
	s = &Service{
		db:  db,
		log: log,
	}
}
