package smspartner

import (
	"fmt"
	"time"
)

// layout defines the format of the reference time.
const layout = "02/01/2006"

type Date struct{ time.Time }

// NewDate creates a new Date
func NewDate(year int, month time.Month, day, hour, min int) Date {
	return Date{Time: time.Date(year, month, day, hour, min, 0, 0, time.UTC)}
}

// ScheduledDeliveryDate returns the date when to send SMS in "dd/mm/YYYY" format
func (date Date) ScheduledDeliveryDate() string {
	return date.Time.Format(layout)
}

// MinutesToSendSMS returns the minute when to send SMS
func (date Date) MinutesToSendSMS() (int, error) {
	min := date.Time.Minute()
	if min%5 != 0 {
		return 0, fmt.Errorf("Minutes must be of 5 minute interval")
	}
	return min, nil
}
