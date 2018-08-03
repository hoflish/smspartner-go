package smspartner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SMS struct {
	ApiKey       string
	PhoneNumbers string
	Message      string
	Gamme        int
	Sender       string

	// FIXME: Define optional params
	// ScheduledDeliveryDate Date
	// Time
	// Minute
	// IsStopSms
	// Sandbox
}

type SMSResponse struct {
	*ResponseState
	MessageID   int     `json:"message_id,omitempty"`
	NumberSMS   int     `json:"nbSMS,omitempty"`
	Cost        float64 `json:"cost,omitempty"`
	Currency    string  `json:"currency,omitempty"`
	PhoneNumber string  `json:"phoneNumber,omitempty"`
}

// SendSMS send SMS, either immediately or at a set time.
func (c *Client) SendSMS(sms *SMS) (*SMSResponse, error) {
	// TODO: Validate SMS fields
	s := &SMS{
		ApiKey:       c.apiKey,
		PhoneNumbers: sms.PhoneNumbers,
		Message:      sms.Message,
	}

	blob, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	fullURL := fmt.Sprintf("%s/send", baseURL)

	req, err := http.NewRequest("POST", fullURL, bytes.NewReader(blob))
	if err != nil {
		return nil, err
	}

	blob, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	smsr := new(SMSResponse)
	if err := json.Unmarshal(blob, smsr); err != nil {
		return nil, err
	}
	return smsr, nil
}
