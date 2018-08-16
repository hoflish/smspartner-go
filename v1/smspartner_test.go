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

func TestSendSMS(t *testing.T) {
	client, err := smspartner.NewClient(&http.Client{})
	if err != nil {
		t.Fatal(err)
	}
	d := smspartner.NewDate(2018, 8, 16, 17, 45)
	minute, err := d.MinuteToSendSMS()
	if err != nil {
		t.Error(err)
	}
	sms := &smspartner.SMS{
		PhoneNumbers: "0620429957",
		Message:      "Your message goes here",
		Gamme:        smspartner.LowCost,
		ScheduledDeliveryDate: d.ScheduledDeliveryDate(),
		Time:   d.Time.Hour(),
		Minute: minute,
	}

	res, err := client.SendSMS(sms)
	if err != nil {
		t.Fatal(err)
	}

	tests := [...]struct {
		res          *smspartner.SMSResponse
		wantNbSMS    int
		wantCost     float64
		wantCurrency string
	}{
		{res, 1, 0.024, "EUR"},
	}

	for _, tt := range tests {
		if tt.res.NumberOfSMS != tt.wantNbSMS {
			t.Errorf("got: %d, want: %d", tt.res.NumberOfSMS, tt.wantNbSMS)
		}

		if tt.res.Cost != tt.wantCost {
			t.Errorf("got: %f, want: %f", tt.res.Cost, tt.wantCost)
		}

		if tt.res.Currency != tt.wantCurrency {
			t.Errorf("got: %s, want: %s", tt.res.Currency, tt.wantCurrency)
		}
	}
}
