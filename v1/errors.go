package smspartner

import "fmt"

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

func (r *Response) hasVError() bool {
	if r == nil {
		return false
	}
	return len(r.VError) > 0
}
