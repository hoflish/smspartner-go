package smspartner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
 // TODO: Define optional params
 type OptionalParams struct {
	// Gamme        int
	// Sender       string
	// ScheduledDeliveryDate Date
	// Time
	// Minute
	// IsStopSms
	// Sandbox
 }
*/

type SMSPayload struct {
	PhoneNumber string `json:"phone"`
	Message     string
}

type SMS struct {
	APIKey       string `json:"apiKey,omitempty"`
	PhoneNumbers string `json:"phoneNumbers,omitempty"`
	Message      string `json:"message,omitempty"`
	//opts OptionalParams
}

type BulkSMS struct {
	APIKey  string        `json:"apiKey,omitempty"`
	SMSList *[]SMSPayload `json:"SMSList,omitempty"`
	//opts OptionalParams
}

type SMSResponseItem struct {
	Success     bool    `json:"success,omitempty"`
	Code        int     `json:"code,omitempty"`
	NumberSMS   int     `json:"nbSms,omitempty"`
	Cost        float64 `json:"cost,omitempty"`
	PhoneNumber string  `json:"phoneNumber,omitempty"`
}

type BulkSMSResponse struct {
	Success         bool               `json:"success,omitempty"`
	Code            int                `json:"code,omitempty"`
	MessageID       int                `json:"message_id,omitempty"`
	Currency        string             `json:"currency,omitempty"`
	Cost            float64            `json:"cost,omitempty"`
	NumberSMS       int                `json:"nbSMS,omitempty"`
	SMSResponseList *[]SMSResponseItem `json:"SMSResponse_List,omiyempty"`
}

// SendSMS sends SMS, either immediately or at a set time.
// See: https://my.smspartner.fr/documentation-fr/api/v1/send-sms
/*
	Example usage:
	--------------
		client, err := smspartner.NewClient()
		// handle err
		sms := &smspartner.SMS{
					PhoneNumbers:    "+212620xxxxxx, +212621xxxxxx",
					Message: "your message",
			},
		}
		res, err := client.SendSMS(sms)
		// handle err
		// diplay response if any

*/
func (c *Client) SendSMS(sms *SMS) (map[string]interface{}, error) {
	sms.APIKey = c.apiKey

	blob, err := json.Marshal(sms)
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

	var smsr map[string]interface{}
	if err := json.Unmarshal(blob, &smsr); err != nil {
		return nil, err
	}
	return smsr, nil
}

// SendBulkSMS sends SMS in batch of 500 either immediately or at a set time.
// See: https://my.smspartner.fr/documentation-fr/api/v1/bulk-send
/*
	Example usage:
	--------------
		client, err := smspartner.NewClient()
		// handle err
		blksms := &smspartner.BulkSMS{
			SMSList: []*smspartner.SMSPayload{
				{
					PhoneNumber:    "+212620xxxxxx",
					Message: "foo",
				},
				{
					PhoneNumber:    "+212620xxxxxx",
					Message: "foobar",
				},
			},
		}
		res, err := client.SendBulkSMS(blksms)
		// handle err
		// diplay response if any

*/
func (c *Client) SendBulkSMS(blksms *BulkSMS) (*BulkSMSResponse, error) {
	blksms.APIKey = c.apiKey

	blob, err := json.Marshal(blksms)
	if err != nil {
		return nil, err
	}
	fullURL := fmt.Sprintf("%s/bulk-send", baseURL)

	req, err := http.NewRequest("POST", fullURL, bytes.NewReader(blob))
	if err != nil {
		return nil, err
	}

	blob, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	blksmsr := new(BulkSMSResponse)
	if err := json.Unmarshal(blob, blksmsr); err != nil {
		return nil, err
	}
	return blksmsr, nil
}
