package main

import (
	"SmsPilot2/example-project/api"
	"SmsPilot2/example-project/internal/app"
	"encoding/json"
	"net/http"
)

func NewRouter(a *app.App) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/send", NewSendHandler(a))

	return mux
}

func NewSendHandler(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		var smsReq api.SendRequest
		if err := json.NewDecoder(r.Body).Decode(&smsReq); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := a.SendSMS(smsReq)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}
