package smspartner

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type NumFormat struct {
	E164          string
	International string
	National      string
	RFC3966       string
}

type LookupItem struct {
	Request     string
	Success     bool
	CountryCode string
	PrefixCode  int
	PhoneNumber string
	Type        string
	Network     string
	Format      *NumFormat
}

type LookupResponse struct {
	Success bool
	Code    int
	Lookup  *[]LookupItem
}

// VerifyNumberFormat checks the format of a phone number
/*
	Example usage:
	--------------
	client, err := smspartner.NewClient()
	// handle err
	phoneNumbers := []string{"+212620xxxxxx", "+212621xxxxxx"}
	res, err := client.VerifyNumberFormat(phoneNumbers...)
	// handle err
	// diplay response if any
*/
func (c *Client) VerifyNumberFormat(phoneNumbers ...string) (*LookupResponse, error) {
	if len(phoneNumbers) == 0 {
		return nil, errors.New("phoneNumber is required")
	}
	p := strings.Join(phoneNumbers, ",")

	var payload struct {
		APIKey       string `json:"apiKey"`
		PhoneNumbers string `json:"phoneNumbers"`
	}
	payload.APIKey = c.apiKey
	payload.PhoneNumbers = p

	blob, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	fullURL := fmt.Sprintf("%s/lookup", baseURL)

	req, err := http.NewRequest("POST", fullURL, bytes.NewReader(blob))
	if err != nil {
		return nil, err
	}

	blob, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	lr := new(LookupResponse)
	if err := json.Unmarshal(blob, lr); err != nil {
		return nil, err
	}
	return lr, nil

}
