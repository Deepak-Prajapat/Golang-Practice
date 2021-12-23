package dbcon

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
var err error
var urlDSN = "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func init() {
	Database, err = gorm.Open(postgres.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err.Error())
		panic("Connection Failed")
	}
	Database.AutoMigrate(&Student{})
}
