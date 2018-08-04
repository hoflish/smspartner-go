package smspartner_test

import (
	"testing"

	"github.com/hoflish/smspartner-go/v1"
)

func TestStatusOK(t *testing.T) {
	got := smspartner.StatusOK(200)
	want := true

	if got != want {
		t.Errorf("got: %t, want: %t", got, want)
	}
}

func TestFirstNonEmptyString(t *testing.T) {
	tests := [...]struct {
		input []string
		want  string
	}{
		0: {input: []string{"foo", "bar"}, want: "foo"},
		1: {input: []string{"", "bar"}, want: "bar"},
		2: {input: []string{""}, want: ""},
	}

	for i, tt := range tests {
		got := smspartner.FirstNonEmptyString(tt.input...)
		if got != tt.want {
			t.Errorf("tc: #%d, got: '%s', want: '%s'", i, got, tt.want)
		}
	}

}
