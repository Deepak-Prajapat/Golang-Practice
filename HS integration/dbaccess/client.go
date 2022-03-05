package dbaccess

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ID     int    `gorm:"primaryKey" json:"id"`
	APIKey string `json:"api_key"`
}
