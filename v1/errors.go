package smspartner

import "fmt"

type SPError struct {
	Success bool     `json:"success,omitempty"`
	Code    int      `json:"code,omitempty"`
	Message string   `json:"message,omitempty"`
	Errs    []*Error `json:"error,omitempty"`
}

type Error struct {
	ElementID string `json:"elementId,omitempty"`
	Message   string `json:"message,omitempty"`
}

func (spe *SPError) Error() string {
	if spe == nil {
		return ""
	}

	if len(spe.Errs) > 0 {
		s, n := "", 0
		for _, e := range spe.Errs {
			if e != nil {
				if n == 0 {
					s = e.Message
				}
				n++
			}
		}
		switch n {
		case 0:
			return "(0 errors)"
		case 1:
			return s
		case 2:
			return s + " (and 1 other error)"
		}
		return fmt.Sprintf("%s (and %d other errors)", s, n-1)
	}

	return spe.Message
}
