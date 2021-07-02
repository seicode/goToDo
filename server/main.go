package main

import (
	"fmt"
	"goToDo/server/router"
	"log"
	"net/http"
)

// Router is exported and used in main.go
func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 1323...")

	log.Fatal(http.ListenAndServe(":1323", r))
}
