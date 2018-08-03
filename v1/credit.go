package smspartner

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Profile struct {
	Username  string `json:"username,omitempty"`
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
}

type Credits struct {
	CreditSMS int `json:"creditSms,omitempty"`
	ToSend    int `json:"toSend,omitempty"`
	Solde     int `json:"solde,omitempty"`
}

type CreditsResponse struct {
	Success bool     `json:"success,omitempty"`
	Code    int      `json:"code,omitempty"`
	User    *Profile `json:"user,omitempty"`
	Credits *Credits `json:"credits,omitempty"`
}

// CheckCredits return your SMS credit (number of SMS available, based on your own purchases and usage),
// as well as the number of SMS that are about to be sent.
// See: https://my.smspartner.fr/documentation-fr/api/v1
/*
	Example usage:
	--------------
		client, err := smspartner.NewClient()
		// handle err
		res, err := client.CheckCredits()
		// handle err
		// display response if any

*/
func (c *Client) CheckCredits() (*CreditsResponse, error) {
	fullURL := fmt.Sprintf("%s/me?apiKey=%s", baseURL, c.apiKey)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	// TODO: handle success response
	credits := new(CreditsResponse)
	if err := json.Unmarshal(res, credits); err != nil {
		return nil, err
	}
	return credits, nil
}
