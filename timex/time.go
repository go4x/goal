// Package timex provides a comprehensive set of time utilities and extensions to Go's standard time package.
// It includes time formatting, business day calculations, timezone handling, and various utility functions
// for common time operations.
//
// Key features:
//   - Time formatting with predefined layouts
//   - Business day calculations and operations
//   - Relative time formatting (e.g., "2 hours ago")
//   - Timezone conversion and DST handling
//   - Batch operations on time slices
//   - Performance measurement utilities
//
// Example usage:
//
//	// Time formatting
//	formatted := timex.Format(time.Now(), timex.YmdDash)
//
//	// Business day calculations
//	businessDays := timex.BusinessDaysBetween(start, end)
//	nextWorkDay := timex.AddBusinessDays(today, 5)
//
//	// Relative time formatting
//	relative := timex.FormatRelative(someTime) // "2 hours ago"
//
//	// Timezone operations
//	beijingTime := timex.ConvertTimezone(utcTime, beijingLocation)
package timex

import (
	"fmt"
	"math"
	"sort"
	"time"
)

// Layout represents a time format layout string.
type Layout string

// Common time format layouts
const (
	// Ymdhms represents compact datetime format: 20060102150405
	Ymdhms = Layout("20060102150405")
	// Ymd represents compact date format: 20060102
	Ymd = Layout("20060102")
	// Hms represents time format: 15:04:05
	Hms = Layout("15:04:05")
	// YmdhmsDash represents datetime with dashes: 2006-01-02 15:04:05
	YmdhmsDash = Layout("2006-01-02 15:04:05")
	// YmdDash represents date with dashes: 2006-01-02
	YmdDash = Layout("2006-01-02")
	// YmdhmsSlash represents datetime with slashes: 2006/01/02 15:04:05
	YmdhmsSlash = Layout("2006/01/02 15:04:05")
	// YmdSlash represents date with slashes: 2006/01/02
	YmdSlash = Layout("2006/01/02")
	// YmdhmsZh represents Chinese datetime format: 2006年01月02日 15时04分05秒
	YmdhmsZh = Layout("2006年01月02日 15时04分05秒")
	// YmdZh represents Chinese date format: 2006年01月02日
	YmdZh = Layout("2006年01月02日")
	// HmsZh represents Chinese time format: 15时04分05秒
	HmsZh = Layout("15时04分05秒")
)

// String returns the string representation of the layout.
func (ly Layout) String() string {
	return string(ly)
}

// From parses a time string using the specified layout.
// Returns an error if the string cannot be parsed.
//
// Example:
//
//	t, err := timex.From("2023-01-01", timex.YmdDash)
func From(s string, layout Layout) (time.Time, error) {
	return time.Parse(layout.String(), s)
}

// MustFrom parses a time string using the specified layout.
// Panics if the string cannot be parsed. Use with caution.
//
// Example:
//
//	t := timex.MustFrom("2023-01-01", timex.YmdDash)
func MustFrom(s string, layout Layout) time.Time {
	t, err := time.Parse(layout.String(), s)
	if err != nil {
		panic(err)
	}
	return t
}

// Format formats a time using the specified layout.
//
// Example:
//
//	formatted := timex.Format(time.Now(), timex.YmdDash)
func Format(date time.Time, layout Layout) string {
	return date.Format(layout.String())
}

// DiffDay calculates the number of days between two times.
// It normalizes both times to UTC and calculates the difference in days,
// ignoring the time component. This ensures consistent results regardless
// of the original timezones.
// Returns a positive integer representing the absolute difference in days.
//
// Example:
//
//	days := timex.DiffDay(time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
//	                     time.Date(2023, 1, 5, 8, 0, 0, 0, time.UTC))
//	// days = 4
func DiffDay(src time.Time, dst time.Time) int {
	// Convert both times to UTC to ensure consistent comparison
	srcUTC := src.UTC()
	dstUTC := dst.UTC()

	// Normalize both times to start of day in UTC
	srcStart := time.Date(srcUTC.Year(), srcUTC.Month(), srcUTC.Day(), 0, 0, 0, 0, time.UTC)
	dstStart := time.Date(dstUTC.Year(), dstUTC.Month(), dstUTC.Day(), 0, 0, 0, 0, time.UTC)

	// Calculate the difference in nanoseconds
	diff := srcStart.Sub(dstStart)

	// Convert to days using integer arithmetic to avoid floating point precision issues
	days := int(diff.Hours() / 24)

	// Return absolute value
	if days < 0 {
		return -days
	}
	return days
}

// DiffSec calculates the absolute difference in seconds between two times.
func DiffSec(src time.Time, dst time.Time) int {
	return int(math.Abs(src.Sub(dst).Seconds()))
}

// DiffMin calculates the absolute difference in minutes between two times.
func DiffMin(src time.Time, dst time.Time) int {
	return int(math.Abs(src.Sub(dst).Minutes()))
}

// DiffHour calculates the absolute difference in hours between two times.
func DiffHour(src time.Time, dst time.Time) int {
	return int(math.Abs(src.Sub(dst).Hours()))
}

// LastDay returns the previous day for the given time.
func LastDay(date time.Time) time.Time {
	return date.AddDate(0, 0, -1)
}

// NextDay returns the next day for the given time.
func NextDay(date time.Time) time.Time {
	return date.AddDate(0, 0, 1)
}

// AddDate adds years, months, and days to a time.
// This is a wrapper around time.Time.AddDate for convenience.
func AddDate(src time.Time, year int, month int, day int) time.Time {
	return src.AddDate(year, month, day)
}

// Add adds a duration to a time.
// This is a wrapper around time.Time.Add for convenience.
func Add(src time.Time, d time.Duration) time.Time {
	return src.Add(d)
}

// SameDate checks if two times represent the same date.
// It normalizes both times to UTC before comparison to ensure consistent results.
func SameDate(src time.Time, dst time.Time) bool {
	srcUTC := src.UTC()
	dstUTC := dst.UTC()

	if srcUTC.Year() != dstUTC.Year() {
		return false
	}
	if srcUTC.Month() != dstUTC.Month() {
		return false
	}
	if srcUTC.Day() != dstUTC.Day() {
		return false
	}
	return true
}

// StartTime returns the start of the day (00:00:00.000000000) for the given time.
func StartTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndTime returns the end of the day (23:59:59.999999999) for the given time.
func EndTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// StartOfMonth returns the start of the month (first day at 00:00:00) for the given time.
func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth returns the end of the month (last day at 23:59:59.999999999) for the given time.
func EndOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, -1, t.Location())
}

// StartOfYear returns the start of the year (January 1st at 00:00:00) for the given time.
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear returns the end of the year (December 31st at 23:59:59.999999999) for the given time.
func EndOfYear(t time.Time) time.Time {
	return time.Date(t.Year()+1, 1, 1, 0, 0, 0, -1, t.Location())
}

// StartOfWeek returns the start of the week (Monday) for the given time.
func StartOfWeek(t time.Time) time.Time {
	weekday := t.Weekday()
	if weekday == time.Sunday {
		weekday = 7 // Treat Sunday as day 7
	}
	daysFromMonday := int(weekday - time.Monday)
	startOfWeek := t.AddDate(0, 0, -daysFromMonday)
	return time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfWeek returns the end of the week (Sunday) for the given time.
func EndOfWeek(t time.Time) time.Time {
	weekday := t.Weekday()
	if weekday == time.Sunday {
		weekday = 7 // Treat Sunday as day 7
	}
	daysToSunday := int(time.Sunday - weekday + 7)
	endOfWeek := t.AddDate(0, 0, daysToSunday)
	return time.Date(endOfWeek.Year(), endOfWeek.Month(), endOfWeek.Day(), 23, 59, 59, 999999999, t.Location())
}

// IsLeapYear checks if the given year is a leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// DaysInMonth returns the number of days in the month for the given year and month.
func DaysInMonth(year int, month time.Month) int {
	switch month {
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	default:
		return 31
	}
}

// ===== Time Checking Methods =====

// IsValidTime validates if a time string is valid
func IsValidTime(s string, layout Layout) bool {
	_, err := time.Parse(layout.String(), s)
	return err == nil
}

// IsToday checks if the time is today
func IsToday(t time.Time) bool {
	now := time.Now()
	return SameDate(t, now)
}

// IsYesterday checks if the time is yesterday
func IsYesterday(t time.Time) bool {
	yesterday := time.Now().AddDate(0, 0, -1)
	return SameDate(t, yesterday)
}

// IsTomorrow checks if the time is tomorrow
func IsTomorrow(t time.Time) bool {
	tomorrow := time.Now().AddDate(0, 0, 1)
	return SameDate(t, tomorrow)
}

// IsWeekend checks if the time is a weekend (Saturday or Sunday)
func IsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// IsBusinessDay checks if the time is a business day (Monday to Friday)
func IsBusinessDay(t time.Time) bool {
	return !IsWeekend(t)
}

// IsInRange checks if the time is within the specified range
func IsInRange(t, start, end time.Time) bool {
	return !t.Before(start) && !t.After(end)
}

// ===== Time Formatting and Beautification Methods =====

// FormatRelative formats time as relative text (e.g., "2 hours ago")
func FormatRelative(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	if diff < 0 {
		diff = -diff
		// Future time
		if diff < time.Minute {
			return "just now"
		} else if diff < time.Hour {
			return fmt.Sprintf("in %d minutes", int(diff.Minutes()))
		} else if diff < 24*time.Hour {
			return fmt.Sprintf("in %d hours", int(diff.Hours()))
		} else {
			return fmt.Sprintf("in %d days", int(diff.Hours()/24))
		}
	} else {
		// Past time
		if diff < time.Minute {
			return "just now"
		} else if diff < time.Hour {
			return fmt.Sprintf("%d minutes ago", int(diff.Minutes()))
		} else if diff < 24*time.Hour {
			return fmt.Sprintf("%d hours ago", int(diff.Hours()))
		} else if diff < 7*24*time.Hour {
			return fmt.Sprintf("%d days ago", int(diff.Hours()/24))
		} else {
			return t.Format("2006-01-02")
		}
	}
}

// TimeAgo calculates time difference and returns a friendly string
func TimeAgo(t time.Time) string {
	return FormatRelative(t)
}

// FormatDuration formats duration as human-readable text
func FormatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%d seconds", int(d.Seconds()))
	} else if d < time.Hour {
		return fmt.Sprintf("%d minutes", int(d.Minutes()))
	} else if d < 24*time.Hour {
		hours := int(d.Hours())
		minutes := int(d.Minutes()) % 60
		if minutes == 0 {
			return fmt.Sprintf("%d hours", hours)
		}
		return fmt.Sprintf("%d hours %d minutes", hours, minutes)
	} else {
		days := int(d.Hours() / 24)
		hours := int(d.Hours()) % 24
		if hours == 0 {
			return fmt.Sprintf("%d days", days)
		}
		return fmt.Sprintf("%d days %d hours", days, hours)
	}
}

// FormatISO formats time in ISO standard format
func FormatISO(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ===== Time Calculation and Operation Methods =====

// BusinessDaysBetween calculates the number of business days between two times
func BusinessDaysBetween(start, end time.Time) int {
	if start.After(end) {
		start, end = end, start
	}

	days := 0
	current := StartTime(start)
	endDay := StartTime(end)

	for current.Before(endDay) || current.Equal(endDay) {
		if IsBusinessDay(current) {
			days++
		}
		current = current.AddDate(0, 0, 1)
	}

	return days
}

// AddBusinessDays adds a specified number of business days
func AddBusinessDays(t time.Time, days int) time.Time {
	if days == 0 {
		return t
	}

	direction := 1
	if days < 0 {
		direction = -1
		days = -days
	}

	result := t
	addedDays := 0

	for addedDays < days {
		result = result.AddDate(0, 0, direction)
		if IsBusinessDay(result) {
			addedDays++
		}
	}

	return result
}

// NextBusinessDay gets the next business day
func NextBusinessDay(t time.Time) time.Time {
	result := NextDay(t)
	for !IsBusinessDay(result) {
		result = NextDay(result)
	}
	return result
}

// PrevBusinessDay gets the previous business day
func PrevBusinessDay(t time.Time) time.Time {
	result := LastDay(t)
	for !IsBusinessDay(result) {
		result = LastDay(result)
	}
	return result
}

// QuarterOf gets the quarter (1-4) for the given time
func QuarterOf(t time.Time) int {
	month := int(t.Month())
	return (month-1)/3 + 1
}

// StartOfQuarter gets the start of the quarter
func StartOfQuarter(t time.Time) time.Time {
	quarter := QuarterOf(t)
	startMonth := time.Month((quarter-1)*3 + 1)
	return time.Date(t.Year(), startMonth, 1, 0, 0, 0, 0, t.Location())
}

// EndOfQuarter gets the end of the quarter
func EndOfQuarter(t time.Time) time.Time {
	quarter := QuarterOf(t)
	endMonth := time.Month(quarter * 3)
	lastDay := DaysInMonth(t.Year(), endMonth)
	return time.Date(t.Year(), endMonth, lastDay, 23, 59, 59, 999999999, t.Location())
}

// ===== Time Validation and Utility Methods =====

// RoundToNearest rounds time to the nearest time unit
func RoundToNearest(t time.Time, unit time.Duration) time.Time {
	nanos := t.UnixNano()
	unitNanos := int64(unit)
	rounded := (nanos + unitNanos/2) / unitNanos * unitNanos
	return time.Unix(0, rounded).In(t.Location())
}

// TruncateTo truncates time to the specified time unit
func TruncateTo(t time.Time, unit time.Duration) time.Time {
	return t.Truncate(unit)
}

// Age calculates age in years
func Age(birthDate, currentDate time.Time) int {
	age := currentDate.Year() - birthDate.Year()
	if currentDate.Month() < birthDate.Month() ||
		(currentDate.Month() == birthDate.Month() && currentDate.Day() < birthDate.Day()) {
		age--
	}
	return age
}

// ===== Batch Operation Methods =====

// SortTimes sorts a slice of times in ascending order
func SortTimes(times []time.Time) []time.Time {
	result := make([]time.Time, len(times))
	copy(result, times)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Before(result[j])
	})
	return result
}

// FindClosest finds the closest time to the target in a slice
func FindClosest(times []time.Time, target time.Time) time.Time {
	if len(times) == 0 {
		return time.Time{}
	}

	closest := times[0]
	minDiff := target.Sub(times[0])
	if minDiff < 0 {
		minDiff = -minDiff
	}

	for _, t := range times[1:] {
		diff := target.Sub(t)
		if diff < 0 {
			diff = -diff
		}
		if diff < minDiff {
			minDiff = diff
			closest = t
		}
	}

	return closest
}

// GroupByDay groups times by day
func GroupByDay(times []time.Time) map[string][]time.Time {
	groups := make(map[string][]time.Time)
	for _, t := range times {
		day := t.Format("2006-01-02")
		groups[day] = append(groups[day], t)
	}
	return groups
}

// FilterByRange filters times that fall within a range
func FilterByRange(times []time.Time, start, end time.Time) []time.Time {
	var result []time.Time
	for _, t := range times {
		if IsInRange(t, start, end) {
			result = append(result, t)
		}
	}
	return result
}

// ===== Timezone Handling Methods =====

// ConvertTimezone converts time to a different timezone
func ConvertTimezone(t time.Time, location *time.Location) time.Time {
	return t.In(location)
}

// GetTimezoneOffset gets the timezone offset
func GetTimezoneOffset(t time.Time) time.Duration {
	_, offset := t.Zone()
	return time.Duration(offset) * time.Second
}

// IsDST checks if the time is in daylight saving time
func IsDST(t time.Time) bool {
	_, winterOffset := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location()).Zone()
	_, summerOffset := time.Date(t.Year(), 7, 1, 0, 0, 0, 0, t.Location()).Zone()
	_, currentOffset := t.Zone()
	return currentOffset != winterOffset && currentOffset == summerOffset
}

// ===== Performance and Statistics Methods =====

// MeasureExecution measures the execution time of a function
func MeasureExecution(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

// GetTimezoneInfo gets comprehensive timezone information
func GetTimezoneInfo(t time.Time) map[string]interface{} {
	name, offset := t.Zone()
	return map[string]interface{}{
		"name":         name,
		"offset":       offset,
		"offset_hours": float64(offset) / 3600,
		"is_dst":       IsDST(t),
		"location":     t.Location().String(),
	}
}
