package types

type SMSID string

type Sms struct {
	ID   SMSID
	From string
	To   string
	Text string
}
