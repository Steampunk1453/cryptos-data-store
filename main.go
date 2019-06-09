package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cryptos-data-store/app"
	"github.com/cryptos-data-store/controllers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/crypto/new", controllers.CreateCrypto).Methods("POST")
	router.HandleFunc("/api/crypto/{id}", controllers.FindCryptoById).Methods("GET")
	router.HandleFunc("/api/crypto/user/{id}", controllers.FindCryptoByUserId).Methods("GET")
	router.HandleFunc("/api/crypto", controllers.FindAllCrypto).Methods("GET")
	router.HandleFunc("/api/crypto", controllers.UpdateCrypto).Methods("PUT")
	router.HandleFunc("/api/crypto/{id}", controllers.DeleteCrypto).Methods("DELETE")

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8080" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

}
