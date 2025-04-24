package booking

import (
	"fmt"
	"os"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
// Mon Jan 2 15:04:05 -0700 MST 2006
func Schedule(date string) time.Time {
	fmt.Println(date)
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	return t
}

// HasPassed returns whether a date has passed.
// Mon Jan 2 15:04:05 -0700 MST 2006
func HasPassed(date string) bool {
	apptTime, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	return apptTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
// Mon Jan 2 15:04:05 -0700 MST 2006
func IsAfternoonAppointment(date string) bool {
	apptTime, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	hour := apptTime.Hour()
	return 12 <= hour && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := Schedule(date)
	y, _, d := t.Date()
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
		t.Weekday().String(), t.Month().String(), d, y, t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
