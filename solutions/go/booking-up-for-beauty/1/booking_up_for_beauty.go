package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	parse, _ := time.Parse(layout, date)
	return parse
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	parse, _ := time.Parse(layout, date)
	return parse.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	parse, _ := time.Parse(layout, date)
	hour := parse.Hour()
	if hour >= 12 && hour < 18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	parse, _ := time.Parse(layout, date)
	weekday := parse.Weekday()
	month := parse.Month()
	day := parse.Day()
	year := parse.Year()
	hour := parse.Hour()
	minute := parse.Minute()
	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", weekday, month, day, year, hour, minute)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
