package timex_test

import (
	"fmt"
	"time"

	"github.com/go4x/goal/timex"
)

// ExampleFormatRelative demonstrates relative time formatting
func ExampleFormatRelative() {
	now := time.Now()

	// Format relative times
	times := []time.Time{
		now.Add(-30 * time.Second),
		now.Add(-5 * time.Minute),
		now.Add(-2 * time.Hour),
		now.Add(-3 * 24 * time.Hour),
	}

	for _, t := range times {
		fmt.Printf("%s -> %s\n", t.Format("15:04:05"), timex.FormatRelative(t))
	}
}

// ExampleBusinessDaysBetween demonstrates business day calculations
func ExampleBusinessDaysBetween() {
	// Calculate business days between two dates
	start := time.Date(2023, 1, 9, 0, 0, 0, 0, time.UTC) // Monday
	end := time.Date(2023, 1, 13, 0, 0, 0, 0, time.UTC)  // Friday

	businessDays := timex.BusinessDaysBetween(start, end)
	fmt.Printf("Business days between %s and %s: %d\n",
		start.Format("2006-01-02"), end.Format("2006-01-02"), businessDays)

	// Add business days to a date
	newDate := timex.AddBusinessDays(start, 3)
	fmt.Printf("3 business days after %s: %s\n",
		start.Format("2006-01-02"), newDate.Format("2006-01-02"))
	// Output:
	// Business days between 2023-01-09 and 2023-01-13: 5
	// 3 business days after 2023-01-09: 2023-01-12
}

// ExampleQuarterOf demonstrates quarter operations
func ExampleQuarterOf() {
	testDate := time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC)

	quarter := timex.QuarterOf(testDate)
	quarterStart := timex.StartOfQuarter(testDate)
	quarterEnd := timex.EndOfQuarter(testDate)

	fmt.Printf("Date: %s\n", testDate.Format("2006-01-02"))
	fmt.Printf("Quarter: %d\n", quarter)
	fmt.Printf("Quarter start: %s\n", quarterStart.Format("2006-01-02"))
	fmt.Printf("Quarter end: %s\n", quarterEnd.Format("2006-01-02"))
	// Output:
	// Date: 2023-05-15
	// Quarter: 2
	// Quarter start: 2023-04-01
	// Quarter end: 2023-06-30
}

// ExampleSortTimes demonstrates time slice operations
func ExampleSortTimes() {
	// Sort a slice of times
	times := []time.Time{
		time.Date(2023, 3, 15, 10, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 10, 14, 0, 0, 0, time.UTC),
		time.Date(2023, 2, 20, 9, 0, 0, 0, time.UTC),
	}

	sorted := timex.SortTimes(times)
	fmt.Print("Sorted times: ")
	for _, t := range sorted {
		fmt.Printf("%s ", t.Format("01-02"))
	}
	fmt.Println()

	// Group by day
	groups := timex.GroupByDay(times)
	fmt.Println("Grouped by day:")
	for day, dayTimes := range groups {
		fmt.Printf("  %s: %d times\n", day, len(dayTimes))
	}
}

// ExampleConvertTimezone demonstrates timezone operations
func ExampleConvertTimezone() {
	utcTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	beijing, _ := time.LoadLocation("Asia/Shanghai")

	// Convert to different timezone
	beijingTime := timex.ConvertTimezone(utcTime, beijing)

	fmt.Printf("UTC time: %s\n", utcTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Beijing time: %s\n", beijingTime.Format("2006-01-02 15:04:05"))

	// Get timezone info
	info := timex.GetTimezoneInfo(beijingTime)
	fmt.Printf("Timezone info: %s, offset: +%.0f hours\n",
		info["name"], info["offset_hours"])
	// Output:
	// UTC time: 2023-01-01 12:00:00
	// Beijing time: 2023-01-01 20:00:00
	// Timezone info: CST, offset: +8 hours
}

// ExampleAge demonstrates age calculation
func ExampleAge() {
	birthDate := time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC)
	currentDate := time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)

	age := timex.Age(birthDate, currentDate)
	fmt.Printf("Born on: %s\n", birthDate.Format("2006-01-02"))
	fmt.Printf("Current date: %s\n", currentDate.Format("2006-01-02"))
	fmt.Printf("Age: %d years\n", age)
	// Output:
	// Born on: 1990-06-15
	// Current date: 2023-06-15
	// Age: 33 years
}

// ExampleFormatDuration demonstrates duration formatting
func ExampleFormatDuration() {
	durations := []time.Duration{
		30 * time.Second,
		5 * time.Minute,
		2*time.Hour + 30*time.Minute,
		3*24*time.Hour + 5*time.Hour,
	}

	for _, d := range durations {
		fmt.Printf("%v -> %s\n", d, timex.FormatDuration(d))
	}
	// Output:
	// 30s -> 30 seconds
	// 5m0s -> 5 minutes
	// 2h30m0s -> 2 hours 30 minutes
	// 77h0m0s -> 3 days 5 hours
}

// ExampleIsToday demonstrates time checking functions
func ExampleIsToday() {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	fmt.Printf("Now is today: %t\n", timex.IsToday(now))
	fmt.Printf("Yesterday is today: %t\n", timex.IsToday(yesterday))
	fmt.Printf("Tomorrow is today: %t\n", timex.IsToday(tomorrow))
	fmt.Printf("Now is business day: %t\n", timex.IsBusinessDay(now))
	fmt.Printf("Now is weekend: %t\n", timex.IsWeekend(now))
}

// ExampleMeasureExecution demonstrates performance measurement
func ExampleMeasureExecution() {
	// Measure execution time of a function
	duration := timex.MeasureExecution(func() {
		// Simulate some work
		sum := 0
		for i := 0; i < 100000; i++ {
			sum += i
		}
	})

	fmt.Printf("Execution time: %v\n", duration)
}
