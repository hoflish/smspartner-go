package smspartner

import "fmt"

type SPError struct {
	*ResponseState `json:"-"`
	Message        string    `json:"message,omitempty"`
	Errors         []*VError `json:"error,omitempty"`
}

type ResponseState struct {
	Success bool `json:"success,omitempty"`
	Code    int  `json:"code,omitempty"`
}

type VError struct {
	ElementID string `json:"elementId,omitempty"`
	Message   string `json:"message,omitempty"`
}

func (spe *SPError) Error() string {
	if spe == nil {
		return ""
	}

	if len(spe.Errors) > 0 {
		msg, n := "", 0
		for _, e := range spe.Errors {
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

	return spe.Message
}
