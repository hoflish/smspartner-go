package smspartner

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const envSMSPartnerAPIKey = "SMSPARTNER_API_KEY"
const apiBasePath = "http://api.smspartner.fr/v1"

var errUnsetAPIKey = fmt.Errorf("could not find %q in your environment", envSMSPartnerAPIKey)

type Client struct {
	hc       *http.Client
	basePath string
	apiKey   string
}

func NewClient(client *http.Client) (*Client, error) {
	wrapClient := new(http.Client)
	*wrapClient = *client

	t := client.Transport
	if t == nil {
		t = http.DefaultTransport
	}

	apiKey, err := getAPIKeyFromEnv()
	if err != nil {
		return nil, err
	}

	return &Client{
		hc:       wrapClient,
		apiKey:   apiKey,
		basePath: apiBasePath,
	}, nil
}

func getAPIKeyFromEnv() (string, error) {
	apikey := strings.TrimSpace(os.Getenv(envSMSPartnerAPIKey))
	if apikey == "" {
		return "", errUnsetAPIKey
	}
	return apikey, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response: %v", err)
	}

	// handle non-200 status code
	if resp.StatusCode != http.StatusOK {
		remResp := &Response{}
		if err := json.Unmarshal(body, remResp); err != nil {
			return nil, fmt.Errorf("error unmarshalling response: %v", err)
		}

		if !remResp.Success && remResp.Code != 200 {
			return nil, errors.New(remResp.errorSummary())
		}

		if remResp == nil {
			return nil, fmt.Errorf("unexpected response: %s", string(body))
		}
	}
	// Each client method handle its expected response data
	return body, nil
}

// Response ,anonymous struct type, has minimal struct fields to check a server response
// other object keys are ignored
type Response struct {
	Success bool               `json:"success"`
	Code    int                `json:"code"`
	Message string             `json:"message,omitempty"`
	VError  []*ValidationError `json:"error,omitempty"`
}

type ValidationError struct {
	ElementID string `json:"elementId,omitempty"`
	Message   string `json:"message,omitempty"`
}

func (r *Response) hasVError() bool {
	if r == nil {
		return false
	}
	return len(r.VError) > 0
}

func (r *Response) errorSummary() string {
	if r.hasVError() {
		msg, n := "", 0
		for _, e := range r.VError {
			if e != nil {
				if n == 0 {
					msg = e.Message
				}
				n++
			}
		}
		switch n {
		case 0:
			return "(0 errors)"
		case 1:
			return msg
		case 2:
			return msg + " (and 1 other error)"
		}
		return fmt.Sprintf("%s (and %d other errors)", msg, n-1)
	}
	return r.Message
}
