package toolkit

import "time"

// RangeToList generates a list of dates between startDate and endDate inclusive.
func RangeToList(startDate, endDate time.Time) []time.Time {
	var dates []time.Time
	currentDate := startDate

	for !currentDate.After(endDate) {
		dates = append(dates, currentDate)
		currentDate = currentDate.AddDate(0, 0, 1) // Increment by one day
	}

	return dates
}
