package db

import "SmsPilot2/example-project/internal/types"

type Client struct {
	driver SomePSQLDriver
}

func NewClient(connstring string) (*Client, error) {
	driver, err := NewPSQLDriver(connstring)
	if err != nil {
		return nil, err
	}
	return &Client{driver: driver}, nil
}

func (c *Client) SaveSMS(sms types.Sms) error {
	return c.driver.Exec("INSERT INTO sms (from, to, text) VALUES ($1, $2, $3)", sms.From, sms.To, sms.Text)
}

func (c *Client) GetSMS(id types.SMSID) (types.Sms, error) {
	var sms types.Sms
	err := c.driver.QueryRow("SELECT from, to, text FROM sms WHERE id = $1", id).Scan(&sms.From, &sms.To, &sms.Text)
	return sms, err
}
