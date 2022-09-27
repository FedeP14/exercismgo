package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour, minutes, seconds := t.Hour(), t.Minute(), t.Second()

	if hour > 12 && hour < 18 {
		return true
	}

	if hour == 12 && minutes > 0 {
		return true
	}

	if hour == 12 && seconds > 0 {
		return true
	}

	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	return Schedule(date).Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return Schedule("9/15/2022 00:00:00")
}
