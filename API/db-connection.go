package main

import (
	"fmt"

	//"gorm.io/driver/mysql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//These variable are beingl used in handlers

var Database *gorm.DB

//var urlDSN = "root:root@tcp(localhost:3306)/godb?parseTime=true" //to establisyh to conlecitono
var urlDSN = "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai" //to establisyh to conlecitono
var err error

func DataMigration() {
	//Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	Database, err = gorm.Open(postgres.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err.Error())
		panic("Connection Failed")
	}

	Database.AutoMigrate(&Employee{})
}
