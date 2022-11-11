package smsprovider

import (
	"SmsPilot2/example-project/internal/types"
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	url        string
	httpClient *http.Client
}

func NewClient(url string, httpClient *http.Client) *Client {
	return &Client{url: url, httpClient: httpClient}
}

type sendRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

type sendResponse struct {
	Ok     bool `json:"ok"`
	Reason string
}

func (c *Client) Send(sms types.Sms) error {
	req := sendRequest{From: sms.From, To: sms.To, Text: sms.Text}
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	res, err := c.httpClient.Post(c.url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	var resp sendResponse
	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return err
	}
	if !resp.Ok {
		return errors.New("sms sending failed: " + resp.Reason)
	}
}
