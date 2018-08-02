package smspartner

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const envSMSPartnerAPIKey = "SMSPARTNER_API_KEY"
const baseURL = "http://api.smspartner.fr/v1/"

var errUnsetAPIKey = fmt.Errorf("could not find %q in your environment", envSMSPartnerAPIKey)

type Client struct {
	apiKey string
}

func NewClient(apikeys ...string) (*Client, error) {
	if apikey := FirstNonEmptyString(apikeys...); apikey != "" {
		return &Client{apiKey: apikey}, nil
	}

	// Otherwise fallback to retrieving it from the environment
	return NewClientFromEnv()
}

func NewClientFromEnv() (*Client, error) {
	apikey := strings.TrimSpace(os.Getenv(envSMSPartnerAPIKey))
	if apikey == "" {
		return nil, errUnsetAPIKey
	}

	return &Client{apiKey: apikey}, nil
}

func (c *Client) httpClient() *http.Client {
	return &http.Client{}
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.httpClient().Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// REVIEW: Handle non 2xx errors

	blob, err := ioutil.ReadAll(res.Body)
	return blob, err
}
