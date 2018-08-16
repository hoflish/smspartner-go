package smspartner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Gamme int

const (
	Premium Gamme = 1
	LowCost Gamme = 2
)

type SMSPayload struct {
	PhoneNumbers string `json:"phoneNumbers,omitempty"`
	Message      string `json:"message,omitempty"`
}

type SMS struct {
	APIKey                string `json:"apiKey,omitempty"`
	PhoneNumbers          string `json:"phoneNumbers,omitempty"`
	Message               string `json:"message,omitempty"`
	Gamme                 Gamme  `json:"gamme,omitempty"`
	Sender                string `json:"sender,omitempty"`
	ScheduledDeliveryDate string `json:"scheduledDeliveryDate,omitempty"`
	Time                  int    `json:"time,omitempty"`
	Minute                int    `json:"minute,omitempty"`
	// IsStopSms
	// Sandbox
}

type BulkSMS struct {
	APIKey                string        `json:"apiKey,omitempty"`
	SMSList               *[]SMSPayload `json:"SMSList,omitempty"`
	Gamme                 Gamme         `json:"gamme,omitempty"`
	Sender                string        `json:"sender,omitempty"`
	ScheduledDeliveryDate string        `json:"scheduledDeliveryDate,omitempty"`
	Time                  int           `json:"time,omitempty"`
	Minute                int           `json:"minute,omitempty"`
	// IsStopSms
	// Sandbox
}

type SMSResponse struct {
	Success               bool    `json:"success,omitempty"`
	Code                  int     `json:"code,omitempty"`
	MessageID             int     `json:"message_id,omitempty"`
	NumberOfSMS           int     `json:"nb_sms,omitempty"`
	Cost                  float64 `json:"cost,omitempty"`
	Currency              string  `json:"currency,omitempty"`
	ScheduledDeliveryDate string  `json:"scheduledDeliveryDate,omitempty"`
}

/*type SMSResponseItem struct {
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
}*/

// SendSMS sends SMS, either immediately or at a set time.
// See: https://my.smspartner.fr/documentation-fr/api/v1/send-sms
/*
	Example usage:
	--------------
		client, err := smspartner.NewClient(&http.Client{})
		// handle err
		// d := smspartner.NewDate(2018, 8, 16, 17, 45)
		// minute, err = d.MinuteToSendSMS()
		// handle err
		sms := &smspartner.SMS{
					PhoneNumbers:    "+212620xxxxxx, +212621xxxxxx",
					Message: "This is your message",
					Gamme: LowCost,
					// ScheduledDeliveryDate: d.ScheduledDeliveryDate(),
					// Time: d.Time.Hour(),
					// Minute: minute
			},
		}
		res, err = client.SendSMS(sms)
		// handle err
		// handle response

*/
func (c *Client) SendSMS(sms *SMS) (*SMSResponse, error) {
	sms.APIKey = c.apiKey

	blob, err := json.Marshal(sms)
	if err != nil {
		return nil, err
	}
	fullURL := fmt.Sprintf("%s/send", c.basePath)

	req, err := http.NewRequest("POST", fullURL, bytes.NewReader(blob))
	if err != nil {
		return nil, err
	}

	blob, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	smsr := new(SMSResponse)
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
					Message: "This is your message",
				},
				{
					PhoneNumber:    "+212620xxxxxx",
					Message: "This is your message",
				},
			},
		}
		res, err := client.SendBulkSMS(blksms)
		// handle err
		// diplay response if any

*/
/*func (c *Client) SendBulkSMS(blksms *BulkSMS) (*BulkSMSResponse, error) {
	blksms.APIKey = c.apiKey

	blob, err := json.Marshal(blksms)
	if err != nil {
		return nil, err
	}
	fullURL := fmt.Sprintf("%s/bulk-send", c.basePath)

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

type VNumber struct {
	APIKey  string
	To      string
	From    string
	Message string

	// TODO: define optional params
	// IsStopSMS
	// Sandbox
	//Format
}

// SendVirtualNumber sendw SMS, either immediately or at a set time, with a long number.
/*
	Example usage:
	--------------
		client, err := smspartner.NewClient()
		// handle err
		vn := &smspartner.VNumber{
			To: "+212620xxxxxx"
			From: "+212620xxxxxx"
			Message: "This is your message"
		}
		res, err := client.SendVirtualNumber(vn)
		// handle err
		// diplay response if any

*/
/*func (c *Client) SendVirtualNumber(vn *VNumber) (map[string]interface{}, error) {
	vn.APIKey = c.apiKey

	blob, err := json.Marshal(vn)
	if err != nil {
		return nil, err
	}
	fullURL := fmt.Sprintf("%s/vn/send", c.basePath)

	req, err := http.NewRequest("POST", fullURL, bytes.NewReader(blob))
	if err != nil {
		return nil, err
	}

	blob, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var vnr map[string]interface{}
	if err := json.Unmarshal(blob, &vnr); err != nil {
		return nil, err
	}
	return vnr, nil
}
*/
