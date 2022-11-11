package api

type SendRequest struct {
	SMS *SMS `json:"sms"`
}

type SMS struct {
	From string
	To   string
	Text string
}
