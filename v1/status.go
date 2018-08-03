package smspartner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetSMSStatus returns the status of an SMS
/*
	TODO: Add example usage
*/
func (c *Client) GetSMSStatus(msgID int, phoneNumbers string) (map[string]interface{}, error) {
	fullURL := fmt.Sprintf("%s/message-status?apiKey=%s&messageId=%d&phoneNumber=%s", baseURL, c.apiKey, msgID, phoneNumbers)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(res, &m); err != nil {
		return nil, err
	}
	return m, nil
}

type SMSStatus struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
	MessageID   int    `json:"messageId,omitempty"`
}

type SMSStatusWrap struct {
	APIKey        string       `json:"apiKey,omitempty"`
	SMSStatusList []*SMSStatus `json:"SMSStatut_List,omitempty"`
}

type SMSStatusResponse struct {
	Success     bool   `json:"success,omitempty"`
	Code        int    `json:"code,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	MessageID   int    `json:"messageId,omitempty"`
	Status      string `json:"status,omitempty"`
	Date        string `json:"date,omitempty"`
	StopSMS     string `json:"stopSms,omitempty"`
	IsSpam      string `json:"isSpam,omitempty"`
}

type SMSStatusResponseWrap struct {
	Success               bool                 `json:"success,omitempty"`
	Code                  int                  `json:"code,omitempty"`
	SMSStatusResponseList []*SMSStatusResponse `json:"StatutResponse_List,omitempty"`
}

// GetMultiSMSStatus returns the status of multiple SMS
/*
	Example usage:
	--------------
		client, err := smspartner.NewClient()
		// handle err
		ss := &smspartner.SMSStatusWrap{
			SMSStatusList: []*smspartner.SMSStatus{
				{
					PhoneNumber: "+212620xxxxxx",
					MessageId: xxxx
				},
				{
					PhoneNumber: "+212621xxxxxx",
					MessageId: xxxx
				}
			}
		}
		res, err := client.GetMultiSMSStatus(ss)
		// handle err
		// display response if any

*/
func (c *Client) GetMultiSMSStatus(ss *SMSStatusWrap) (*SMSStatusResponseWrap, error) {
	ss.APIKey = c.apiKey
	blob, err := json.Marshal(ss)
	if err != nil {
		return nil, err
	}
	fullURL := fmt.Sprintf("%s/multi-status", baseURL)

	req, err := http.NewRequest("POST", fullURL, bytes.NewReader(blob))
	if err != nil {
		return nil, err
	}

	blob, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ssr := new(SMSStatusResponseWrap)
	if err := json.Unmarshal(blob, ssr); err != nil {
		return nil, err
	}
	return ssr, nil
}
