package smspartner_test

import (
	"testing"

	"github.com/hoflish/smspartner-go/v1"
)

func TestSummaryError(t *testing.T) {
	var tests = []struct {
		remAPIErrResp  *smspartner.RemoteAPIError
		wantSummaryErr string
	}{
		{&smspartner.RemoteAPIError{}, ""},
		{&smspartner.RemoteAPIError{Message: "msg error"}, "msg error"},
		{&smspartner.RemoteAPIError{
			Message: "foobarzoo error",
			VError: []*smspartner.ValidationError{
				{ElementID: "children[foo].data", Message: "foo error"},
				{ElementID: "children[bar].data", Message: "bar error"},
				{ElementID: "children[zoo].data", Message: "zoo error"},
			},
		}, "foo error (and 2 other errors)"},
	}

	for _, tt := range tests {
		want := tt.remAPIErrResp.ErrorSummary()
		got := tt.wantSummaryErr
		if want != got {
			t.Errorf("got: '%s', want: '%s'", want, got)
		}
	}
}
