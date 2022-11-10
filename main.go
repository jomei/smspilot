package main

import (
	"SmsPilot2/internal/loginGetToken"
	"SmsPilot2/internal/sms"
	"SmsPilot2/pkg/jwt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/login-get-token", loginGetToken.Get).Methods("POST")
	r.Handle("/v1/add-sms", jwt.IsAuthorized(sms.Add)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

}
