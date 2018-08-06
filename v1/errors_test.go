package smspartner_test

import (
	"testing"

	"github.com/hoflish/smspartner-go/v1"
)

func TestError(t *testing.T) {
	var tests = []struct {
		remAPIErrResp *smspartner.RemoteAPIError
		wantErr       string
	}{
		{&smspartner.RemoteAPIError{}, "(0 errors)"},
		{&smspartner.RemoteAPIError{
			VError: []*smspartner.ValidationError{
				{ElementID: "children[foo].data", Message: "foo error"},
			},
		}, "foo error"},
		{&smspartner.RemoteAPIError{
			VError: []*smspartner.ValidationError{
				{ElementID: "children[foo].data", Message: "foo error"},
				{ElementID: "children[bar].data", Message: "bar error"},
			},
		}, "foo error (and 1 other error)"},
		{&smspartner.RemoteAPIError{
			Message: "foobar error",
			VError: []*smspartner.ValidationError{
				{ElementID: "children[foo].data", Message: "foo error"},
				{ElementID: "children[bar].data", Message: "bar error"},
				{ElementID: "children[zoo].data", Message: "zoo error"},
			},
		}, "foo error (and 2 other errors)"},
	}

	for _, tt := range tests {
		want := tt.remAPIErrResp.Error()
		got := tt.wantErr
		if want != got {
			t.Errorf("got: '%s', want: '%s'", want, got)
		}
	}
}
