package main

import (
	"SmsPilot2/example-project/internal/app"
	"flag"
	"net/http"
)

var smsProviderPath = flag.String("sms-provider", "http://sms-provider", "path to sms provider")
var dbConnString = flag.String("db-conn-string", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable", "connection string to db")

func main() {
	flag.Parse()
	httpClient := http.DefaultClient
	dbClient, err := NewDBClient(*dbConnString)
	if err != nil {
		panic(err)
	}
	smsProviderClient := NewSMSProviderClient(*smsProviderPath, httpClient)
	a := app.NewApp(smsProviderClient, dbClient)
	http.ListenAndServe(":8080", NewRouter(a))
}
