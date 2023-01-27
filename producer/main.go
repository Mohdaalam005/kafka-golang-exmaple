package main

import (
	"github.com/gorilla/mux"
	produce_service "github.com/mohdaalam005/producing-msg/produce-service"
	"log"
	"net/http"
)

func main() {
	log.Println("application has been started ")
	route := mux.NewRouter()
	route.HandleFunc("/", produce_service.SendMsg).Methods("GET")
	http.ListenAndServe(":8080", route)
}
