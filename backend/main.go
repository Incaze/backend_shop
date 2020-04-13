package main

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"shop/backend/store"
	"shop/backend/utils"
)

func main() {
	router := store.CreateRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	handlerCRUD := handlers.CORS(allowedOrigins, allowedMethods)(router)

	log.Fatal(http.ListenAndServe(":"+utils.PORT, handlerCRUD))
}
