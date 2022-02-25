package dbaccess

func SetupMigrations() {
	err := s.db.AutoMigrate(&Client{})
	if err != nil {
		return
	}
}
