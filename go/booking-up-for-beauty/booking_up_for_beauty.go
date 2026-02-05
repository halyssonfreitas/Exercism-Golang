package booking

import (
	"log"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	time, err := time.Parse(layout, date)
	if err != nil {
		log.Fatalln(err)
	}
	return time
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:4:5"
	schedule, err := time.Parse(layout, date)
	if err != nil {
		log.Fatalln(err)
	}
	return time.Now().After(schedule)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:4:5"
	time, err := time.Parse(layout, date)
	if err != nil {
		log.Fatalln(err)
	}
	hour := time.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layoutIn := "1/2/2006 15:4:5"
	layoutOut := "You have an appointment on Monday, January 2, 2006, at 15:4."
	time, err := time.Parse(layoutIn, date)
	if err != nil {
		log.Fatalln(err)
	}
	return time.Format(layoutOut)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
