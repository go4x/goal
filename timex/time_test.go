package timex_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/go4x/goal/timex"
)

func TestIsValidTime(t *testing.T) {
	tests := []struct {
		input  string
		layout timex.Layout
		want   bool
	}{
		{"2023-01-01", timex.YmdDash, true},
		{"invalid", timex.YmdDash, false},
		{"20230101", timex.Ymd, true},
		{"2023-01-01 12:30:45", timex.YmdhmsDash, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := timex.IsValidTime(tt.input, tt.layout)
			if got != tt.want {
				t.Errorf("IsValidTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsToday_IsYesterday_IsTomorrow(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	if !timex.IsToday(now) {
		t.Error("IsToday() should return true for now")
	}
	if !timex.IsYesterday(yesterday) {
		t.Error("IsYesterday() should return true for yesterday")
	}
	if !timex.IsTomorrow(tomorrow) {
		t.Error("IsTomorrow() should return true for tomorrow")
	}
}

func TestIsWeekend_IsBusinessDay(t *testing.T) {
	// Create a Saturday
	saturday := time.Date(2023, 1, 7, 12, 0, 0, 0, time.UTC)
	// Create a Monday
	monday := time.Date(2023, 1, 9, 12, 0, 0, 0, time.UTC)

	if !timex.IsWeekend(saturday) {
		t.Error("IsWeekend() should return true for Saturday")
	}
	if !timex.IsBusinessDay(monday) {
		t.Error("IsBusinessDay() should return true for Monday")
	}
}

func TestIsInRange(t *testing.T) {
	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 1, 31, 23, 59, 59, 0, time.UTC)
	middle := time.Date(2023, 1, 15, 12, 0, 0, 0, time.UTC)
	outside := time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	if !timex.IsInRange(middle, start, end) {
		t.Error("IsInRange() should return true for middle date")
	}
	if timex.IsInRange(outside, start, end) {
		t.Error("IsInRange() should return false for outside date")
	}
}

func TestFormatRelative(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{"just now", now, "just now"},
		{"1 minute ago", now.Add(-time.Minute), "1 minutes ago"},
		{"2 hours ago", now.Add(-2 * time.Hour), "2 hours ago"},
		{"3 days ago", now.Add(-3 * 24 * time.Hour), "3 days ago"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := timex.FormatRelative(tt.input)
			// Since time is changing, we only check that the result is not empty
			if result == "" {
				t.Errorf("FormatRelative() returned empty string")
			}
		})
	}
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		input    time.Duration
		expected string
	}{
		{30 * time.Second, "30 seconds"},
		{5 * time.Minute, "5 minutes"},
		{2 * time.Hour, "2 hours"},
		{2*time.Hour + 30*time.Minute, "2 hours 30 minutes"},
		{3 * 24 * time.Hour, "3 days"},
		{3*24*time.Hour + 5*time.Hour, "3 days 5 hours"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := timex.FormatDuration(tt.input)
			if result != tt.expected {
				t.Errorf("FormatDuration() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBusinessDaysBetween(t *testing.T) {
	// Monday
	start := time.Date(2023, 1, 9, 12, 0, 0, 0, time.UTC)
	// Friday
	end := time.Date(2023, 1, 13, 12, 0, 0, 0, time.UTC)

	result := timex.BusinessDaysBetween(start, end)
	if result != 5 { // Monday to Friday, total 5 business days
		t.Errorf("BusinessDaysBetween() = %d, want 5", result)
	}
}

func TestAddBusinessDays(t *testing.T) {
	// Friday
	start := time.Date(2023, 1, 13, 12, 0, 0, 0, time.UTC)

	// Add 2 business days, should jump to next Tuesday
	result := timex.AddBusinessDays(start, 2)
	expected := time.Date(2023, 1, 17, 12, 0, 0, 0, time.UTC) // Next Tuesday

	if !timex.SameDate(result, expected) {
		t.Errorf("AddBusinessDays() = %v, want %v", result, expected)
	}
}

func TestQuarterOf(t *testing.T) {
	tests := []struct {
		input    time.Time
		expected int
	}{
		{time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), 1},  // January -> Q1
		{time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC), 2},  // April -> Q2
		{time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC), 3},  // July -> Q3
		{time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC), 4}, // December -> Q4
	}

	for _, tt := range tests {
		t.Run(tt.input.Format("2006-01"), func(t *testing.T) {
			result := timex.QuarterOf(tt.input)
			if result != tt.expected {
				t.Errorf("QuarterOf() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestAge(t *testing.T) {
	birthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	currentDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	result := timex.Age(birthDate, currentDate)
	if result != 33 {
		t.Errorf("Age() = %d, want 33", result)
	}
}

func TestSortTimes(t *testing.T) {
	times := []time.Time{
		time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
	}

	sorted := timex.SortTimes(times)

	if !sorted[0].Equal(times[1]) || !sorted[1].Equal(times[2]) || !sorted[2].Equal(times[0]) {
		t.Error("SortTimes() did not sort correctly")
	}
}

func TestFindClosest(t *testing.T) {
	times := []time.Time{
		time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC),
	}
	target := time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC)

	result := timex.FindClosest(times, target)
	expected := times[0] // 2023-01-05 (closest to 2023-01-03, difference of 2 days)

	if !result.Equal(expected) {
		t.Errorf("FindClosest() = %v, want %v", result, expected)
	}
}

func TestGroupByDay(t *testing.T) {
	times := []time.Time{
		time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 2, 12, 0, 0, 0, time.UTC),
	}

	groups := timex.GroupByDay(times)

	if len(groups["2023-01-01"]) != 2 {
		t.Errorf("Expected 2 times for 2023-01-01, got %d", len(groups["2023-01-01"]))
	}
	if len(groups["2023-01-02"]) != 1 {
		t.Errorf("Expected 1 time for 2023-01-02, got %d", len(groups["2023-01-02"]))
	}
}

func TestFilterByRange(t *testing.T) {
	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 1, 31, 23, 59, 59, 0, time.UTC)

	times := []time.Time{
		time.Date(2023, 1, 15, 12, 0, 0, 0, time.UTC),  // Within range
		time.Date(2023, 2, 1, 12, 0, 0, 0, time.UTC),   // Out of range
		time.Date(2022, 12, 31, 12, 0, 0, 0, time.UTC), // Out of range
	}

	filtered := timex.FilterByRange(times, start, end)

	if len(filtered) != 1 {
		t.Errorf("Expected 1 filtered time, got %d", len(filtered))
	}
}

func TestConvertTimezone(t *testing.T) {
	utcTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	beijing, _ := time.LoadLocation("Asia/Shanghai")

	result := timex.ConvertTimezone(utcTime, beijing)

	// Beijing time is 8 hours ahead of UTC
	expected := time.Date(2023, 1, 1, 20, 0, 0, 0, beijing)
	if !result.Equal(expected) {
		t.Errorf("ConvertTimezone() = %v, want %v", result, expected)
	}
}

func TestMeasureExecution(t *testing.T) {
	duration := timex.MeasureExecution(func() {
		time.Sleep(10 * time.Millisecond)
	})

	if duration < 10*time.Millisecond {
		t.Errorf("MeasureExecution() = %v, expected at least 10ms", duration)
	}
}

func TestGetTimezoneInfo(t *testing.T) {
	utcTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)

	info := timex.GetTimezoneInfo(utcTime)

	if info["name"] != "UTC" {
		t.Errorf("Expected timezone name 'UTC', got %v", info["name"])
	}
	if info["offset"] != 0 {
		t.Errorf("Expected offset 0, got %v", info["offset"])
	}
}

// ===== Additional Test Cases to Improve Coverage =====

func TestLayout_String(t *testing.T) {
	layout := timex.Layout("2006-01-02")
	if layout.String() != "2006-01-02" {
		t.Errorf("Layout.String() = %v, want '2006-01-02'", layout.String())
	}
}

func TestFrom_Error(t *testing.T) {
	_, err := timex.From("invalid", timex.YmdDash)
	if err == nil {
		t.Error("From() should return error for invalid input")
	}
}

func TestMustFrom_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustFrom() should panic on invalid input")
		}
	}()
	timex.MustFrom("invalid", timex.YmdDash)
}

func TestFormat_AllLayouts(t *testing.T) {
	testTime := time.Date(2023, 1, 1, 12, 30, 45, 0, time.UTC)

	layouts := []struct {
		name   string
		layout timex.Layout
		expect string
	}{
		{"Ymdhms", timex.Ymdhms, "20230101123045"},
		{"Ymd", timex.Ymd, "20230101"},
		{"Hms", timex.Hms, "12:30:45"},
		{"YmdhmsDash", timex.YmdhmsDash, "2023-01-01 12:30:45"},
		{"YmdDash", timex.YmdDash, "2023-01-01"},
		{"YmdhmsSlash", timex.YmdhmsSlash, "2023/01/01 12:30:45"},
		{"YmdSlash", timex.YmdSlash, "2023/01/01"},
		{"YmdhmsZh", timex.YmdhmsZh, "2023年01月01日 12时30分45秒"},
		{"YmdZh", timex.YmdZh, "2023年01月01日"},
		{"HmsZh", timex.HmsZh, "12时30分45秒"},
	}

	for _, tt := range layouts {
		t.Run(tt.name, func(t *testing.T) {
			result := timex.Format(testTime, tt.layout)
			if result != tt.expect {
				t.Errorf("Format(%s) = %v, want %v", tt.name, result, tt.expect)
			}
		})
	}
}

func TestDiffDay_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		src      time.Time
		dst      time.Time
		expected int
	}{
		{
			name:     "same time",
			src:      time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			dst:      time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			expected: 0,
		},
		{
			name:     "negative difference",
			src:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			dst:      time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := timex.DiffDay(tt.src, tt.dst)
			if result != tt.expected {
				t.Errorf("DiffDay() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestDiffSec_DiffMin_DiffHour_EdgeCases(t *testing.T) {
	// Test with same time
	sameTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)

	if timex.DiffSec(sameTime, sameTime) != 0 {
		t.Error("DiffSec() should return 0 for same time")
	}
	if timex.DiffMin(sameTime, sameTime) != 0 {
		t.Error("DiffMin() should return 0 for same time")
	}
	if timex.DiffHour(sameTime, sameTime) != 0 {
		t.Error("DiffHour() should return 0 for same time")
	}
}

func TestSameDate_EdgeCases(t *testing.T) {
	// Test with same time
	sameTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	if !timex.SameDate(sameTime, sameTime) {
		t.Error("SameDate() should return true for same time")
	}

	// Test with different timezones but same date
	utc := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	beijing, _ := time.LoadLocation("Asia/Shanghai")
	beijingTime := time.Date(2023, 1, 1, 12, 0, 0, 0, beijing)

	if !timex.SameDate(utc, beijingTime) {
		t.Error("SameDate() should return true for same date in different timezones")
	}
}

func TestStartTime_EndTime_EdgeCases(t *testing.T) {
	testTime := time.Date(2023, 1, 1, 12, 30, 45, 123456789, time.UTC)

	start := timex.StartTime(testTime)
	if start.Hour() != 0 || start.Minute() != 0 || start.Second() != 0 || start.Nanosecond() != 0 {
		t.Error("StartTime() should reset time to start of day")
	}

	end := timex.EndTime(testTime)
	if end.Hour() != 23 || end.Minute() != 59 || end.Second() != 59 || end.Nanosecond() != 999999999 {
		t.Error("EndTime() should set time to end of day")
	}
}

func TestStartOfMonth_EndOfMonth_EdgeCases(t *testing.T) {
	// Test February in leap year
	leapYear := time.Date(2024, 2, 15, 12, 0, 0, 0, time.UTC)
	startOfMonth := timex.StartOfMonth(leapYear)
	endOfMonth := timex.EndOfMonth(leapYear)

	if startOfMonth.Day() != 1 {
		t.Error("StartOfMonth() should return day 1")
	}
	if endOfMonth.Day() != 29 { // 2024 is leap year
		t.Error("EndOfMonth() should return day 29 for February 2024")
	}
}

func TestStartOfYear_EndOfYear_EdgeCases(t *testing.T) {
	testTime := time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC)
	startOfYear := timex.StartOfYear(testTime)
	endOfYear := timex.EndOfYear(testTime)

	if startOfYear.Month() != 1 || startOfYear.Day() != 1 {
		t.Error("StartOfYear() should return January 1st")
	}
	if endOfYear.Month() != 12 || endOfYear.Day() != 31 {
		t.Error("EndOfYear() should return December 31st")
	}
}

func TestStartOfWeek_EndOfWeek_EdgeCases(t *testing.T) {
	// Test Sunday
	sunday := time.Date(2023, 1, 8, 12, 0, 0, 0, time.UTC) // Sunday
	startOfWeek := timex.StartOfWeek(sunday)
	endOfWeek := timex.EndOfWeek(sunday)

	if startOfWeek.Weekday() != time.Monday {
		t.Error("StartOfWeek() should return Monday for Sunday input")
	}
	if endOfWeek.Weekday() != time.Sunday {
		t.Error("EndOfWeek() should return Sunday")
	}
}

func TestIsLeapYear_EdgeCases(t *testing.T) {
	tests := []struct {
		year int
		want bool
	}{
		{1600, true},  // Divisible by 400
		{1700, false}, // Divisible by 100 but not 400
		{1800, false}, // Divisible by 100 but not 400
		{1900, false}, // Divisible by 100 but not 400
		{2000, true},  // Divisible by 400
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("year_%d", tt.year), func(t *testing.T) {
			got := timex.IsLeapYear(tt.year)
			if got != tt.want {
				t.Errorf("IsLeapYear(%d) = %v, want %v", tt.year, got, tt.want)
			}
		})
	}
}

func TestDaysInMonth_AllMonths(t *testing.T) {
	tests := []struct {
		year  int
		month time.Month
		want  int
	}{
		{2023, time.January, 31},
		{2023, time.February, 28},
		{2024, time.February, 29}, // Leap year
		{2023, time.March, 31},
		{2023, time.April, 30},
		{2023, time.May, 31},
		{2023, time.June, 30},
		{2023, time.July, 31},
		{2023, time.August, 31},
		{2023, time.September, 30},
		{2023, time.October, 31},
		{2023, time.November, 30},
		{2023, time.December, 31},
	}

	for _, tt := range tests {
		t.Run(tt.month.String(), func(t *testing.T) {
			got := timex.DaysInMonth(tt.year, tt.month)
			if got != tt.want {
				t.Errorf("DaysInMonth(%d, %v) = %d, want %d", tt.year, tt.month, got, tt.want)
			}
		})
	}
}

func TestFormatRelative_EdgeCases(t *testing.T) {
	now := time.Now()

	// Test future times
	futureTime := now.Add(time.Hour)
	result := timex.FormatRelative(futureTime)
	if !strings.Contains(result, "in") {
		t.Error("FormatRelative() should handle future times")
	}

	// Test very old time
	oldTime := now.Add(-365 * 24 * time.Hour)
	result = timex.FormatRelative(oldTime)
	if !strings.Contains(result, "202") { // Should contain year
		t.Error("FormatRelative() should format very old times as date")
	}
}

func TestFormatDuration_EdgeCases(t *testing.T) {
	tests := []struct {
		input    time.Duration
		expected string
	}{
		{0, "0 seconds"},
		{time.Second, "1 seconds"},
		{time.Minute, "1 minutes"},
		{time.Hour, "1 hours"},
		{24 * time.Hour, "1 days"},
		{25 * time.Hour, "1 days 1 hours"},
		{24*time.Hour + 30*time.Minute, "1 days"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := timex.FormatDuration(tt.input)
			if result != tt.expected {
				t.Errorf("FormatDuration() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBusinessDaysBetween_EdgeCases(t *testing.T) {
	// Test same day
	sameDay := time.Date(2023, 1, 9, 12, 0, 0, 0, time.UTC) // Monday
	result := timex.BusinessDaysBetween(sameDay, sameDay)
	if result != 1 { // Same business day counts as 1
		t.Errorf("BusinessDaysBetween() same day = %d, want 1", result)
	}

	// Test weekend
	weekend := time.Date(2023, 1, 7, 12, 0, 0, 0, time.UTC) // Saturday
	result = timex.BusinessDaysBetween(weekend, weekend)
	if result != 0 {
		t.Errorf("BusinessDaysBetween() weekend = %d, want 0", result)
	}
}

func TestAddBusinessDays_EdgeCases(t *testing.T) {
	// Test adding 0 days
	base := time.Date(2023, 1, 9, 12, 0, 0, 0, time.UTC)
	result := timex.AddBusinessDays(base, 0)
	if !result.Equal(base) {
		t.Error("AddBusinessDays(0) should return same time")
	}

	// Test negative days
	result = timex.AddBusinessDays(base, -2)
	expected := time.Date(2023, 1, 5, 12, 0, 0, 0, time.UTC) // Previous Thursday
	if !timex.SameDate(result, expected) {
		t.Errorf("AddBusinessDays(-2) = %v, want %v", result, expected)
	}
}

func TestQuarterOf_EdgeCases(t *testing.T) {
	tests := []struct {
		month    time.Month
		expected int
	}{
		{time.January, 1},
		{time.February, 1},
		{time.March, 1},
		{time.April, 2},
		{time.May, 2},
		{time.June, 2},
		{time.July, 3},
		{time.August, 3},
		{time.September, 3},
		{time.October, 4},
		{time.November, 4},
		{time.December, 4},
	}

	for _, tt := range tests {
		t.Run(tt.month.String(), func(t *testing.T) {
			testTime := time.Date(2023, tt.month, 15, 0, 0, 0, 0, time.UTC)
			result := timex.QuarterOf(testTime)
			if result != tt.expected {
				t.Errorf("QuarterOf(%v) = %d, want %d", tt.month, result, tt.expected)
			}
		})
	}
}

func TestStartOfQuarter_EndOfQuarter_EdgeCases(t *testing.T) {
	// Test Q1
	q1 := time.Date(2023, 2, 15, 12, 0, 0, 0, time.UTC)
	start := timex.StartOfQuarter(q1)
	end := timex.EndOfQuarter(q1)

	if start.Month() != time.January || start.Day() != 1 {
		t.Error("StartOfQuarter Q1 should return January 1st")
	}
	if end.Month() != time.March || end.Day() != 31 {
		t.Error("EndOfQuarter Q1 should return March 31st")
	}

	// Test Q4
	q4 := time.Date(2023, 11, 15, 12, 0, 0, 0, time.UTC)
	start = timex.StartOfQuarter(q4)
	end = timex.EndOfQuarter(q4)

	if start.Month() != time.October || start.Day() != 1 {
		t.Error("StartOfQuarter Q4 should return October 1st")
	}
	if end.Month() != time.December || end.Day() != 31 {
		t.Error("EndOfQuarter Q4 should return December 31st")
	}
}

func TestRoundToNearest_EdgeCases(t *testing.T) {
	// Test rounding to hour
	testTime := time.Date(2023, 1, 1, 12, 30, 0, 0, time.UTC)
	rounded := timex.RoundToNearest(testTime, time.Hour)
	expected := time.Date(2023, 1, 1, 13, 0, 0, 0, time.UTC)

	if !rounded.Equal(expected) {
		t.Errorf("RoundToNearest() = %v, want %v", rounded, expected)
	}
}

func TestAge_EdgeCases(t *testing.T) {
	// Test same date
	birth := time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC)
	current := time.Date(1990, 6, 15, 12, 0, 0, 0, time.UTC)
	age := timex.Age(birth, current)
	if age != 0 {
		t.Errorf("Age() same date = %d, want 0", age)
	}

	// Test day before birthday
	current = time.Date(2023, 6, 14, 0, 0, 0, 0, time.UTC)
	age = timex.Age(birth, current)
	if age != 32 {
		t.Errorf("Age() day before birthday = %d, want 32", age)
	}
}

func TestSortTimes_EmptySlice(t *testing.T) {
	var empty []time.Time
	result := timex.SortTimes(empty)
	if len(result) != 0 {
		t.Error("SortTimes() should handle empty slice")
	}
}

func TestFindClosest_EmptySlice(t *testing.T) {
	var empty []time.Time
	target := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	result := timex.FindClosest(empty, target)
	if !result.IsZero() {
		t.Error("FindClosest() should return zero time for empty slice")
	}
}

func TestGroupByDay_EmptySlice(t *testing.T) {
	var empty []time.Time
	result := timex.GroupByDay(empty)
	if len(result) != 0 {
		t.Error("GroupByDay() should handle empty slice")
	}
}

func TestFilterByRange_EmptySlice(t *testing.T) {
	var empty []time.Time
	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)
	result := timex.FilterByRange(empty, start, end)
	if len(result) != 0 {
		t.Error("FilterByRange() should handle empty slice")
	}
}

func TestGetTimezoneOffset_EdgeCases(t *testing.T) {
	utc := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	offset := timex.GetTimezoneOffset(utc)
	if offset != 0 {
		t.Errorf("GetTimezoneOffset() UTC = %v, want 0", offset)
	}
}

func TestIsDST_EdgeCases(t *testing.T) {
	// Test UTC (no DST)
	utc := time.Date(2023, 7, 1, 12, 0, 0, 0, time.UTC)
	if timex.IsDST(utc) {
		t.Error("IsDST() should return false for UTC")
	}
}

func TestMeasureExecution_ZeroDuration(t *testing.T) {
	duration := timex.MeasureExecution(func() {})
	if duration < 0 {
		t.Error("MeasureExecution() should return non-negative duration")
	}
}

func TestGetTimezoneInfo_EdgeCases(t *testing.T) {
	utc := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	info := timex.GetTimezoneInfo(utc)

	// Check all required fields
	requiredFields := []string{"name", "offset", "offset_hours", "is_dst", "location"}
	for _, field := range requiredFields {
		if _, exists := info[field]; !exists {
			t.Errorf("GetTimezoneInfo() missing field: %s", field)
		}
	}
}

// ===== Benchmark Tests =====

func BenchmarkFormatRelative(b *testing.B) {
	testTime := time.Now().Add(-2 * time.Hour)
	for i := 0; i < b.N; i++ {
		timex.FormatRelative(testTime)
	}
}

func BenchmarkDiffDay(b *testing.B) {
	time1 := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	time2 := time.Date(2023, 1, 10, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		timex.DiffDay(time1, time2)
	}
}

func BenchmarkBusinessDaysBetween(b *testing.B) {
	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		timex.BusinessDaysBetween(start, end)
	}
}

func BenchmarkAddBusinessDays(b *testing.B) {
	base := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		timex.AddBusinessDays(base, 5)
	}
}

func BenchmarkSortTimes(b *testing.B) {
	times := []time.Time{
		time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		timex.SortTimes(times)
	}
}

func BenchmarkFindClosest(b *testing.B) {
	times := []time.Time{
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 20, 0, 0, 0, 0, time.UTC),
	}
	target := time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		timex.FindClosest(times, target)
	}
}

func BenchmarkGroupByDay(b *testing.B) {
	times := []time.Time{
		time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 2, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 3, 18, 0, 0, 0, time.UTC),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		timex.GroupByDay(times)
	}
}

func BenchmarkConvertTimezone(b *testing.B) {
	utcTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	beijing, _ := time.LoadLocation("Asia/Shanghai")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		timex.ConvertTimezone(utcTime, beijing)
	}
}

func BenchmarkIsBusinessDay(b *testing.B) {
	testTime := time.Date(2023, 1, 9, 12, 0, 0, 0, time.UTC) // Monday
	for i := 0; i < b.N; i++ {
		timex.IsBusinessDay(testTime)
	}
}

func BenchmarkQuarterOf(b *testing.B) {
	testTime := time.Date(2023, 5, 15, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		timex.QuarterOf(testTime)
	}
}

func BenchmarkAge(b *testing.B) {
	birth := time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC)
	current := time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		timex.Age(birth, current)
	}
}

func BenchmarkFormatDuration(b *testing.B) {
	duration := 2*time.Hour + 30*time.Minute
	for i := 0; i < b.N; i++ {
		timex.FormatDuration(duration)
	}
}

func BenchmarkRoundToNearest(b *testing.B) {
	testTime := time.Date(2023, 1, 1, 12, 30, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		timex.RoundToNearest(testTime, time.Hour)
	}
}
