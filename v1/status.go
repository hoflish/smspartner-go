package smspartner

import (
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
