package main

import "fmt"

func Foo() {
	fmt.Println("hello duniya")
}

type StaffMember struct {
	Name string
	ID   string
}

type Developer struct {
	memberDetails  StaffMember
	Skills  []string
	Package float32
}
