package main

import (
	"log"
	"net/http"

	"github.com/Evedel/MYOBTestService/handlers"
)

func main() {

	handlers.HandleRequests()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
