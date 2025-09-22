# TimeX - Comprehensive Time Utilities for Go

TimeX is a comprehensive time utilities package that extends Go's standard `time` package with time formatting, business day calculations, timezone handling, and various utility functions for common time operations.

## Features

- **Time Formatting**: Multiple predefined layouts and custom formatting
- **Business Day Calculations**: Calculate working days and business operations
- **Relative Time Formatting**: Human-readable time differences (e.g., "2 hours ago")
- **Timezone Handling**: DST support and timezone conversions
- **Batch Operations**: Sort, filter, and group time slices
- **Performance Utilities**: Execution time measurement and optimization

## Installation

```bash
go get github.com/go4x/goal/timex
```

## Quick Start

```go
package main

import (
    "fmt"
    "time"
    "github.com/go4x/goal/timex"
)

func main() {
    // Time formatting
    t := time.Now()
    formatted := timex.Format(t, timex.YmdDash)
    fmt.Println(formatted)
    
    // Business day calculations
    start := time.Date(2023, 1, 9, 0, 0, 0, 0, time.UTC) // Monday
    end := time.Date(2023, 1, 13, 0, 0, 0, 0, time.UTC)   // Friday
    businessDays := timex.BusinessDaysBetween(start, end)
    fmt.Printf("Business days: %d\n", businessDays)
    
    // Relative time formatting
    relative := timex.FormatRelative(time.Now().Add(-2 * time.Hour))
    fmt.Printf("Time ago: %s\n", relative) // "2 hours ago"
}
```

## API Reference

### Time Parsing

#### `From(s string, layout Layout) (time.Time, error)`
Parse a time string using a specific layout.

```go
// Parse with specific layout
t1, _ := timex.From("2023-01-01 12:30:45", timex.YmdhmsDash)
t2, _ := timex.From("2023/01/01 12:30:45", timex.YmdhmsSlash)
t3, _ := timex.From("2023-01-01", timex.YmdDash)
```

#### `MustFrom(s string, layout Layout) time.Time`
Parse a time string using a specific layout, panicking on error.

### Time Formatting

#### `Format(date time.Time, layout Layout) string`
Format a time using the specified layout.

#### `FormatRelative(t time.Time) string`
Format time as relative text (e.g., "2小时前", "3天后").

#### `FormatDuration(d time.Duration) string`
Format duration as human-readable text (e.g., "2小时30分钟").

#### `FormatChinese(t time.Time) string`
Format time in Chinese format (e.g., "2023年01月01日 12时30分45秒").

#### `FormatISO(t time.Time) string`
Format time in ISO standard format (RFC3339).

### Time Calculations

#### `DiffDay(src, dst time.Time) int`
Calculate the absolute difference in days between two times.

#### `BusinessDaysBetween(start, end time.Time) int`
Calculate the number of business days between two times.

#### `AddBusinessDays(t time.Time, days int) time.Time`
Add a specified number of business days to a time.

#### `NextBusinessDay(t time.Time) time.Time`
Get the next business day after the given time.

#### `PrevBusinessDay(t time.Time) time.Time`
Get the previous business day before the given time.

### Time Queries

#### `IsToday(t time.Time) bool`
Check if the time is today.

#### `IsYesterday(t time.Time) bool`
Check if the time is yesterday.

#### `IsTomorrow(t time.Time) bool`
Check if the time is tomorrow.

#### `IsWeekend(t time.Time) bool`
Check if the time is a weekend (Saturday or Sunday).

#### `IsBusinessDay(t time.Time) bool`
Check if the time is a business day (Monday to Friday).

#### `IsInRange(t, start, end time.Time) bool`
Check if the time is within the specified range.

### Period Boundaries

#### `StartTime(t time.Time) time.Time`
Get the start of the day (00:00:00.000000000).

#### `EndTime(t time.Time) time.Time`
Get the end of the day (23:59:59.999999999).

#### `StartOfWeek(t time.Time) time.Time`
Get the start of the week (Monday).

#### `EndOfWeek(t time.Time) time.Time`
Get the end of the week (Sunday).

#### `StartOfMonth(t time.Time) time.Time`
Get the start of the month (first day).

#### `EndOfMonth(t time.Time) time.Time`
Get the end of the month (last day).

#### `StartOfYear(t time.Time) time.Time`
Get the start of the year (January 1st).

#### `EndOfYear(t time.Time) time.Time`
Get the end of the year (December 31st).

#### `QuarterOf(t time.Time) int`
Get the quarter (1-4) for the given time.

#### `StartOfQuarter(t time.Time) time.Time`
Get the start of the quarter.

#### `EndOfQuarter(t time.Time) time.Time`
Get the end of the quarter.

### Date Utilities

#### `Age(birthDate, currentDate time.Time) int`
Calculate age in years.

#### `IsLeapYear(year int) bool`
Check if the year is a leap year.

#### `DaysInMonth(year int, month time.Month) int`
Get the number of days in a month for a given year.

#### `SameDate(src, dst time.Time) bool`
Check if two times represent the same date.

### Timezone Operations

#### `ConvertTimezone(t time.Time, location *time.Location) time.Time`
Convert time to a different timezone.

#### `GetTimezoneOffset(t time.Time) time.Duration`
Get the timezone offset.

#### `IsDST(t time.Time) bool`
Check if the time is in daylight saving time.

#### `GetTimezoneInfo(t time.Time) map[string]interface{}`
Get comprehensive timezone information.

### Batch Operations

#### `SortTimes(times []time.Time) []time.Time`
Sort a slice of times in ascending order.

#### `FindClosest(times []time.Time, target time.Time) time.Time`
Find the closest time to the target in a slice.

#### `GroupByDay(times []time.Time) map[string][]time.Time`
Group times by day.

#### `FilterByRange(times []time.Time, start, end time.Time) []time.Time`
Filter times that fall within a range.

### Utility Functions

#### `RoundToNearest(t time.Time, unit time.Duration) time.Time`
Round time to the nearest unit.

#### `TruncateTo(t time.Time, unit time.Duration) time.Time`
Truncate time to the specified unit.

#### `MeasureExecution(fn func()) time.Duration`
Measure the execution time of a function.

## Predefined Layouts

TimeX provides several predefined layout constants:

```go
timex.Ymdhms      // "20060102150405"
timex.Ymd         // "20060102"
timex.Hms         // "15:04:05"
timex.YmdhmsDash  // "2006-01-02 15:04:05"
timex.YmdDash     // "2006-01-02"
timex.YmdhmsSlash // "2006/01/02 15:04:05"
timex.YmdSlash    // "2006/01/02"
timex.YmdhmsZh    // "2006年01月02日 15时04分05秒"
timex.YmdZh       // "2006年01月02日"
timex.HmsZh       // "15时04分05秒"
```

## Performance

TimeX is optimized for performance with many functions achieving zero allocations:

```
BenchmarkParseSmart-12             	19308361	        62.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkDiffDay-12                	28509064	        42.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBusinessDay-12          	398281634	         3.024 ns/op	       0 B/op	       0 allocs/op
BenchmarkConvertTimezone-12        	1000000000	         0.3156 ns/op	       0 B/op	       0 allocs/op
```

## Testing

The package includes comprehensive tests with 86.6% code coverage:

```bash
go test -cover ./timex
```

Run benchmarks:

```bash
go test -bench=. -benchmem ./timex
```

## Examples

See the [example_test.go](example_test.go) file for comprehensive usage examples.

## License

This project is part of the Goal library and follows the same license terms.

## Contributing

Contributions are welcome! Please ensure all tests pass and maintain the high code coverage.
