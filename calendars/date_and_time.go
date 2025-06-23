package calendars

import "time"

// DateOnly returns the date with 00:00:00.0 time from the provided time instance in the same location
func DateOnly(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GetMonthName get the name of the month specified or an empty string
//
//goland:noinspection GoUnusedExportedFunction
func GetMonthName(month int) string {
	switch month {
	case 1:
		return "January"
	case 2:
		return "February"
	case 3:
		return "March"
	case 4:
		return "April"
	case 5:
		return "May"
	case 6:
		return "June"
	case 7:
		return "July"
	case 8:
		return "August"
	case 9:
		return "September"
	case 10:
		return "October"
	case 11:
		return "November"
	case 12:
		return "December"
	}
	return ""
}

// DaysInPeriod calculates the total number of days in the specified period, calculating week days, monday to thursdays, fridays and saturday to sundays
// returns total days | weekdays | monday-thursday | fridays | saturday-sunday
//
//goland:noinspection GoUnusedExportedFunction
func DaysInPeriod(start time.Time, end time.Time) (int, int, int, int, int) {
	totalDays := 0
	weekDays := 0
	mondayToThursday := 0
	fridays := 0
	saturdayPlusSunday := 0
	for {
		if start.Weekday() != time.Saturday && start.Weekday() != time.Sunday {
			if start.Weekday() == time.Monday || start.Weekday() == time.Tuesday || start.Weekday() == time.Wednesday || start.Weekday() == time.Thursday {
				mondayToThursday++
			}
			if start.Weekday() == time.Friday {
				fridays++
			}
			weekDays++
		}
		if start.Weekday() == time.Saturday || start.Weekday() == time.Sunday {
			saturdayPlusSunday++
		}
		totalDays++
		start = start.Add(time.Hour * 24)
		
		if start.Equal(end) {
			return totalDays, weekDays, mondayToThursday, fridays, saturdayPlusSunday
		}
	}
}

// IsDateEqual checks if the year and day of year are equal
//
//goland:noinspection GoUnusedExportedFunction
func IsDateEqual(date1 time.Time, date2 time.Time) bool {
	return date1.Year() == date2.Year() && date1.YearDay() == date2.YearDay()
}

// AsZone returns 'date' as a new time.Time with the specified 'loc'
//
//goland:noinspection GoUnusedExportedFunction
func AsZone(date time.Time, loc *time.Location) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), loc)
}
