package smspartner

import (
	"fmt"
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
	if apikey := firstNonEmptyString(apikeys...); apikey != "" {
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

//func (c *Client) doRequest(req *http.Request) {}

// firstNonEmptyString iterates through its
// arguments trying to find the first string
// that is not blank or consists entirely  of spaces.
func firstNonEmptyString(args ...string) string {
	for _, arg := range args {
		if arg == "" {
			continue
		}
		if strings.TrimSpace(arg) != "" {
			return arg
		}
	}
	return ""
}
