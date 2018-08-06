package smspartner

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const envSMSPartnerAPIKey = "SMSPARTNER_API_KEY"
const apiBasePath = "http://api.smspartner.fr/v1"
const clientDefaultTimeout time.Duration = 10 * time.Second

var errUnsetAPIKey = fmt.Errorf("could not find %q in your environment", envSMSPartnerAPIKey)

type Client struct {
	hc       *http.Client
	basePath string
	apiKey   string
}

// NewClient returns an HTTP client.
func NewClient(client *http.Client) (*Client, error) {
	wrapClient := new(http.Client)
	*wrapClient = *client

	t := client.Timeout
	if t == 0 {
		t = clientDefaultTimeout
	}
	tr := client.Transport
	if tr == nil {
		tr = http.DefaultTransport
	}

	wrapClient.Timeout = t
	wrapClient.Transport = tr

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

// REVIEW:(hoflish) add context to client methods
// IMPORTANT: on client side, do we really need context ??
func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.hc.Do(req)
	// REVIEW:(hoflish) handle timeout error ?
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
		remAPIErr := &RemoteAPIError{}
		if err := json.Unmarshal(body, remAPIErr); err != nil {
			return nil, fmt.Errorf("error unmarshalling response: %v", err)
		}

		if !remAPIErr.Success && remAPIErr.Code != 200 {
			if remAPIErr.Message != "" {
				return nil, errors.New(remAPIErr.Message)
			}
			return nil, errors.New(remAPIErr.Error())
		}

		if remAPIErr == nil {
			return nil, fmt.Errorf("unexpected response: %s", string(body))
		}
	}
	// Note: each client method handle its expected response data
	return body, nil
}
