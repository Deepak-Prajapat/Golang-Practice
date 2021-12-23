package main

import (
	"log"
	"net/http"
	"school/pkg/crud"
	_ "school/pkg/dbcon"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", crud.Routing()))
}
