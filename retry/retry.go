package retry

import (
	"errors"
	"math/rand"
	"time"
)

// defInterval is the default interval strategy that uses exponential backoff
var defInterval = &defaultInterval{}

// F is the function type to be retried.
// It returns two values:
//   - bool: true to stop retrying, false to continue retrying
//   - error: nil for success, non-nil for failure
//
// Retry behavior:
//   - If error is nil, retry stops immediately regardless of the bool value
//   - If error is not nil, the bool value determines whether to continue retrying
type F func() (bool, error)

// Intervaler defines the interface for retry interval strategies.
// Implementations determine how long to wait between retry attempts.
type Intervaler interface {
	// Interval is called between retry attempts to determine wait time.
	// Parameter n is the current retry attempt number (1-based).
	Interval(n uint)
}

// defaultInterval implements exponential backoff strategy with jitter.
// It sleeps for 2^n seconds plus random jitter between retry attempts.
type defaultInterval struct {
}

// Interval implements exponential backoff with jitter: sleep for 2^n seconds with random jitter
func (s *defaultInterval) Interval(n uint) {
	base := time.Second * (1 << n)
	// Add jitter: random value between 0 and base/2
	jitter := time.Duration(rand.Int63n(int64(base / 2)))
	time.Sleep(base + jitter)
}

// constantInterval implements a constant interval strategy.
// It sleeps for the same duration between all retry attempts.
type constantInterval struct {
	interval time.Duration
}

// Interval sleeps for the constant interval duration
func (s *constantInterval) Interval(n uint) {
	time.Sleep(s.interval)
}

// DefaultInterval returns the default interval strategy.
// It implements exponential backoff strategy with jitter.
// It sleeps for 2^n seconds plus random jitter between retry attempts.
//
// If no interval strategy is set, it will be used as default.
func DefaultInterval() Intervaler {
	return defInterval
}

// ConstantInterval creates a constant interval strategy.
// It returns an Intervaler that sleeps for the same duration between all retry attempts.
//
// Parameters:
//   - interval: The duration to wait between retry attempts
//
// Returns:
//   - Intervaler: A strategy that uses constant intervals
//
// Example:
//
//	retry.Interval(retry.ConstantInterval(100 * time.Millisecond))
func ConstantInterval(interval time.Duration) Intervaler {
	return &constantInterval{interval: interval}
}

// jitterInterval implements exponential backoff with configurable jitter.
// It provides more control over jitter behavior than the default strategy.
type jitterInterval struct {
	base   time.Duration
	jitter float64 // jitter factor (0.0 to 1.0)
}

// Interval implements exponential backoff with configurable jitter
func (j *jitterInterval) Interval(n uint) {
	base := j.base * time.Duration(1<<n)
	jitterAmount := float64(base) * j.jitter
	jitter := time.Duration(rand.Float64() * jitterAmount)
	time.Sleep(base + jitter)
}

// ExponentialBackoffWithJitter creates an exponential backoff strategy with configurable jitter.
// This provides more control over jitter behavior than the default strategy.
//
// Parameters:
//   - base: Base duration for the first retry
//   - jitter: Jitter factor (0.0 to 1.0), where 0.0 = no jitter, 1.0 = 100% jitter
//
// Returns:
//   - Intervaler: A strategy that uses exponential backoff with jitter
//
// Example:
//
//	// 10% jitter
//	retry.Interval(retry.ExponentialBackoffWithJitter(time.Second, 0.1))
//
//	// 50% jitter
//	retry.Interval(retry.ExponentialBackoffWithJitter(time.Second, 0.5))
func ExponentialBackoffWithJitter(base time.Duration, jitter float64) Intervaler {
	if jitter < 0 {
		jitter = 0
	}
	if jitter > 1 {
		jitter = 1
	}
	return &jitterInterval{
		base:   base,
		jitter: jitter,
	}
}

// setting holds the configuration for retry behavior
type setting struct {
	times    uint                    // Maximum number of retries
	interval Intervaler              // Strategy for intervals between retries
	callback func(n uint, err error) // Callback function called after each retry
}

// settings is a function type used to configure retry behavior
// It follows the functional options pattern for flexible configuration
type settings func(p *setting)

// Times sets the maximum number of retry attempts.
// This is a required setting - if not provided, Do will return an error.
//
// Parameters:
//   - n: Maximum number of retries (must be > 0)
//
// Returns:
//   - settings: A configuration function for the retry behavior
func Times(n uint) settings {
	return func(p *setting) {
		p.times = n
	}
}

// Interval sets the interval strategy for determining wait time between retries.
// If not set, the default exponential backoff strategy will be used.
//
// Parameters:
//   - s: An Intervaler implementation that determines wait time between retries
//
// Returns:
//   - settings: A configuration function for the retry behavior
//
// Example:
//
//	retry.Interval(retry.ConstantInterval(time.Second))
func Interval(s Intervaler) settings {
	return func(p *setting) {
		p.interval = s
	}
}

// Callback sets a callback function that will be called after each retry attempt.
// This is useful for logging, monitoring, or implementing custom retry logic.
//
// Parameters:
//   - c: Callback function that receives:
//   - n: Current retry attempt number (1-based)
//   - err: The error returned by the retried function
//
// Returns:
//   - settings: A configuration function for the retry behavior
//
// Example:
//
//	retry.Callback(func(n uint, err error) {
//	    log.Printf("Retry attempt %d failed: %v", n, err)
//	})
func Callback(c func(n uint, err error)) settings {
	return func(p *setting) {
		p.callback = c
	}
}

// Do executes the given function with retry logic.
// It will retry the function until it succeeds, stops explicitly, or reaches the maximum number of retries.
//
// Parameters:
//   - f: The function to be retried. It should return (bool, error) where:
//   - pf: Optional settings to configure retry behavior:
//
// Returns:
//   - error: nil if the function succeeded, or the last error if all retries failed
//
// Example:
//
//	err := retry.Do(func() (bool, error) {
//	    // Your operation here
//	    return false, someError
//	}, retry.Times(3), retry.Interval(retry.ConstantInterval(time.Second)))
func Do(f F, pf ...settings) error {
	var p setting
	var n uint
	var err error
	var stop bool
	p.interval = defInterval
	for _, fn := range pf {
		fn(&p)
	}
	if p.times == 0 {
		return errors.New("retry times not set")
	}
	for {
		if stop, err = f(); err == nil || stop {
			return err
		}
		n++
		if p.callback != nil {
			p.callback(n, err)
		}
		if n > p.times {
			return err
		}
		p.interval.Interval(uint(n))
	}
}
