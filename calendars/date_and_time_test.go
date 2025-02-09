package calendars

import (
	"testing"
	"time"
)

func Test_DaysInPeriod_July2024(t *testing.T) {
	start := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC)

	totalDays, weekDays, mondayToThursday, fridays, saturdayPlusSunday := DaysInPeriod(start, end)

	if totalDays != 31 {
		t.Errorf("Expected 31 days total, got %d", totalDays)
	}
	if weekDays != 23 {
		t.Errorf("Expected 23 week days, got %d", weekDays)
	}
	if mondayToThursday != 19 {
		t.Errorf("Expected 19 monday to thursdays, got %d", mondayToThursday)
	}
	if fridays != 4 {
		t.Errorf("Expected 4 fridays, got %d", fridays)
	}
	if saturdayPlusSunday != 8 {
		t.Errorf("Expected 8 saturday + sundays, got %d", saturdayPlusSunday)
	}
}

func Test_DaysInPeriod_JulyToSeptember2024(t *testing.T) {
	start := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC)

	totalDays, weekDays, mondayToThursday, fridays, saturdayPlusSunday := DaysInPeriod(start, end)

	if totalDays != 92 {
		t.Errorf("Expected 92 days total, got %d", totalDays)
	}
	if weekDays != 66 {
		t.Errorf("Expected 66 week days, got %d", weekDays)
	}
	if mondayToThursday != 53 {
		t.Errorf("Expected 53 monday to thursdays, got %d", mondayToThursday)
	}
	if fridays != 13 {
		t.Errorf("Expected 13 fridays, got %d", fridays)
	}
	if saturdayPlusSunday != 26 {
		t.Errorf("Expected 26 saturday + sundays, got %d", saturdayPlusSunday)
	}
}

func Test_GetMonthName(t *testing.T) {
	month1 := GetMonthName(1)
	expectedMonth1 := "January"
	if month1 != expectedMonth1 {
		t.Errorf("Expected %s, got %s", expectedMonth1, month1)
	}

	month13 := GetMonthName(13)
	if month13 != "" {
		t.Errorf("Expected 'empty string', got %s", month13)
	}
}

func Test_IsDateEqual(t *testing.T) {
	start1 := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)
	end1 := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)
	res1 := IsDateEqual(start1, end1)
	if res1 != true {
		t.Errorf("Expected '%s' to be equal to '%s'", start1, end1)
	}

	start2 := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)
	end2 := time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC)
	res2 := IsDateEqual(start2, end2)
	if res2 != false {
		t.Errorf("Expected '%s' not to be equal to '%s'", start2, end2)
	}
}

func Test_AsZone(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Copenhagen")
	expectedRes := time.Date(2024, 7, 1, 0, 0, 0, 0, loc)
	date := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)
	res := AsZone(date, loc)
	if res != expectedRes {
		t.Errorf("Expected '%s' to be '%s'", res, expectedRes)
	}
}
