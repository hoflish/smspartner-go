package smspartner_test

import (
	"net/http"
	"testing"

	smspartner "github.com/hoflish/smspartner-go/v1"
)

func TestCredits(t *testing.T) {
	client, err := smspartner.NewClient(&http.Client{})
	if err != nil {
		t.Fatal(err)
	}
	res, err := client.CheckCredits()
	if err != nil {
		t.Fatal(err)
	}

	got := res.Credits.Currency
	want := "EUR"
	if got != want {
		t.Errorf("got: '%s', want: '%s'", got, want)
	}
}
