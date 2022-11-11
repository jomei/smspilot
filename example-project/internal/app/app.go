package app

import (
	"SmsPilot2/example-project/api"
	"SmsPilot2/example-project/internal/types"
)

type SMSProvider interface {
	Send(sms types.Sms) error
}

type Storage interface {
	SaveSMS(sms types.Sms) error
	GetSMS(id types.SMSID) (types.Sms, error)
}

type App struct {
	providerClient SMSProvider
	storage        Storage
}

func NewApp(providerClient SMSProvider, storage Storage) *App {
	return &App{providerClient: providerClient, storage: storage}
}

func (a *App) SendSMS(smsReq api.SendRequest) error {
	sms := types.Sms{From: smsReq.SMS.From, To: smsReq.SMS.To, Text: smsReq.SMS.Text}
	if err := a.providerClient.Send(sms); err != nil {
		return err
	}
	return a.storage.SaveSMS(sms)
}
