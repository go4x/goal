package retry_test

import (
	"errors"
	"testing"
	"time"

	"github.com/go4x/goal/retry"
)

// TestBasicRetry tests the basic retry functionality
func TestBasicRetry(t *testing.T) {
	attempts := 0

	f := retry.F(func() (bool, error) {
		attempts++
		if attempts < 3 {
			return false, errors.New("temporary error")
		}
		return true, nil // Success
	})

	err := retry.Do(f, retry.Times(5))
	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}
}

// TestRetryWithCallback tests retry with callback function
func TestRetryWithCallback(t *testing.T) {
	attempts := 0
	callbackCalls := 0

	f := retry.F(func() (bool, error) {
		attempts++
		if attempts < 2 {
			return false, errors.New("temporary error")
		}
		return true, nil
	})

	callback := func(n uint, err error) {
		callbackCalls++
		t.Logf("Retry attempt %d failed: %v", n, err)
	}

	err := retry.Do(f, retry.Times(3), retry.Callback(callback))
	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
	if attempts != 2 {
		t.Errorf("Expected 2 attempts, got %d", attempts)
	}
	if callbackCalls != 1 {
		t.Errorf("Expected 1 callback call, got %d", callbackCalls)
	}
}

// TestRetryWithConstantInterval tests retry with constant interval
func TestRetryWithConstantInterval(t *testing.T) {
	attempts := 0

	f := retry.F(func() (bool, error) {
		attempts++
		if attempts < 2 {
			return false, errors.New("temporary error")
		}
		return true, nil
	})

	start := time.Now()
	err := retry.Do(f,
		retry.Times(3),
		retry.Interval(retry.ConstantInterval(100*time.Millisecond)))
	elapsed := time.Since(start)

	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
	if attempts != 2 {
		t.Errorf("Expected 2 attempts, got %d", attempts)
	}
	// Should have slept for at least 100ms
	if elapsed < 100*time.Millisecond {
		t.Errorf("Expected at least 100ms elapsed, got %v", elapsed)
	}
}

// TestRetryMaxAttemptsReached tests when max attempts are reached
func TestRetryMaxAttemptsReached(t *testing.T) {
	attempts := 0

	f := retry.F(func() (bool, error) {
		attempts++
		return false, errors.New("permanent error")
	})

	err := retry.Do(f, retry.Times(3))
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if attempts != 4 { // 1 initial + 3 retries
		t.Errorf("Expected 4 attempts, got %d", attempts)
	}
	if err.Error() != "permanent error" {
		t.Errorf("Expected 'permanent error', got: %v", err)
	}
}

// TestRetryTimesNotSet tests when times is not set
func TestRetryTimesNotSet(t *testing.T) {
	f := retry.F(func() (bool, error) {
		return true, nil
	})

	err := retry.Do(f)
	if err == nil {
		t.Error("Expected error for missing times setting")
	}
	if err.Error() != "retry times not set" {
		t.Errorf("Expected 'times not set' error, got: %v", err)
	}
}

// TestRetryImmediateSuccess tests when function succeeds immediately
func TestRetryImmediateSuccess(t *testing.T) {
	attempts := 0

	f := retry.F(func() (bool, error) {
		attempts++
		return true, nil // Immediate success
	})

	err := retry.Do(f, retry.Times(3))
	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
	if attempts != 1 {
		t.Errorf("Expected 1 attempt, got %d", attempts)
	}
}

// TestJitterInterval tests the jitter interval functionality
func TestJitterInterval(t *testing.T) {
	attempts := 0

	f := retry.F(func() (bool, error) {
		attempts++
		if attempts < 2 {
			return false, errors.New("temporary error")
		}
		return true, nil
	})

	// Test with jitter interval
	start := time.Now()
	err := retry.Do(f,
		retry.Times(3),
		retry.Interval(retry.ExponentialBackoffWithJitter(100*time.Millisecond, 0.5)),
	)
	elapsed := time.Since(start)

	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
	if attempts != 2 {
		t.Errorf("Expected 2 attempts, got %d", attempts)
	}
	// Should have slept for at least 100ms (base time)
	if elapsed < 100*time.Millisecond {
		t.Errorf("Expected at least 100ms elapsed, got %v", elapsed)
	}
}

// TestJitterIntervalNoJitter tests jitter interval with 0% jitter
func TestJitterIntervalNoJitter(t *testing.T) {
	attempts := 0

	f := retry.F(func() (bool, error) {
		attempts++
		if attempts < 2 {
			return false, errors.New("temporary error")
		}
		return true, nil
	})

	// Test with no jitter (0%)
	start := time.Now()
	err := retry.Do(f,
		retry.Times(3),
		retry.Interval(retry.ExponentialBackoffWithJitter(50*time.Millisecond, 0.0)),
	)
	elapsed := time.Since(start)

	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
	if attempts != 2 {
		t.Errorf("Expected 2 attempts, got %d", attempts)
	}
	// Should be close to 50ms (base time) with minimal jitter
	if elapsed < 50*time.Millisecond || elapsed > 60*time.Millisecond {
		t.Errorf("Expected around 50ms elapsed, got %v", elapsed)
	}
}
