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
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With"})

	handlerCRUD := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)

	log.Fatal(http.ListenAndServe(":"+utils.PORT, handlerCRUD))
}
