package main

import "gorm.io/gorm"

type Employee struct {
	gorm.Model         //To get by default fields like creareddAta
	EmpName    string  `json:"empname"`
	EmpSalary  float64 `json:"salary"`
	Email      string  `json:"email"`
}
