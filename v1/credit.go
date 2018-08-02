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
	*SPError
	User    *Profile `json:"user,omitempty"`
	Credits *Credits `json:"credits,omitempty"`
}

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
