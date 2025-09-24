package random

// Global random number generator for better performance
// Using crypto/rand for better seed generation
var globalRand = NewSecure()

// Int generates a random integer in the range [0, max).
//
// This function provides a convenient way to generate random integers with better performance
// than creating a new random number generator each time.
//
// Parameters:
//   - max: upper bound (exclusive), must be > 0
//
// Returns: random integer in range [0, max)
//
// Example: Int(10) returns a random integer from 0 to 9
func Int(max int) int {
	if max <= 0 {
		panic("max must be greater than 0")
	}
	return globalRand.Intn(max)
}

// Between generates a random integer in the range [min, max).
//
// This function generates a random integer between min (inclusive) and max (exclusive).
// It's more efficient than creating a new random number generator for each call.
//
// Parameters:
//   - min: lower bound (inclusive), must be >= 0
//   - max: upper bound (exclusive), must be >= min
//
// Returns: random integer in range [min, max)
//
// Example: Between(5, 15) returns a random integer from 5 to 14
func Between(min, max int) int {
	if min < 0 {
		panic("min must be >= 0")
	}
	if min > max {
		panic("min must be <= max")
	}
	if min == max {
		return min
	}
	return min + globalRand.Intn(max-min)
}

// Float64 generates a random float64 in the range [0.0, 1.0).
//
// Returns: random float64 in range [0.0, 1.0)
//
// Example: Float64() returns a random float like 0.123456789
func Float64() float64 {
	return globalRand.Float64()
}

// Float64Between generates a random float64 in the range [min, max).
//
// Parameters:
//   - min: lower bound (inclusive)
//   - max: upper bound (exclusive), must be > min
//
// Returns: random float64 in range [min, max)
//
// Example: Float64Between(1.5, 3.7) returns a random float between 1.5 and 3.7
func Float64Between(min, max float64) float64 {
	if min >= max {
		panic("min must be less than max")
	}
	return min + globalRand.Float64()*(max-min)
}

// Bool generates a random boolean value.
//
// Returns: random boolean (true or false)
//
// Example: Bool() returns true or false with equal probability
func Bool() bool {
	return globalRand.Intn(2) == 1
}

// Choice selects a random element from a slice.
//
// Parameters:
//   - slice: the slice to choose from (must not be empty)
//
// Returns: random element from the slice
//
// Example: Choice([]string{"A", "B", "C"}) returns "A", "B", or "C"
func Choice[T any](slice []T) T {
	if len(slice) == 0 {
		panic("slice cannot be empty")
	}
	return slice[globalRand.Intn(len(slice))]
}

// Shuffle shuffles a slice in place using Fisher-Yates algorithm.
//
// This function randomly reorders the elements of the slice in place.
// It uses the Fisher-Yates shuffle algorithm for unbiased results.
//
// Parameters:
//   - slice: the slice to shuffle (modified in place)
//
// Returns: none (slice is modified in place)
//
// Example: Shuffle([]int{1,2,3,4,5}) randomly reorders the elements
func Shuffle[T any](slice []T) {
	for i := len(slice) - 1; i > 0; i-- {
		j := globalRand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Sample samples k elements from a slice without replacement.
//
// This function randomly selects k elements from the slice without replacement.
// It uses reservoir sampling for efficient selection.
//
// Parameters:
//   - slice: the slice to sample from
//   - k: number of elements to sample (must be 0 < k <= len(slice))
//
// Returns: slice containing k randomly selected elements, or nil if invalid parameters
//
// Example: Sample([]string{"A","B","C","D","E"}, 3) returns 3 random elements
func Sample[T any](slice []T, k int) []T {
	if k <= 0 || k > len(slice) {
		return nil
	}

	result := make([]T, k)
	copy(result, slice[:k])

	for i := k; i < len(slice); i++ {
		j := globalRand.Intn(i + 1)
		if j < k {
			result[j] = slice[i]
		}
	}

	return result
}

// WeightedChoice represents a choice with a weight and value.
type WeightedChoice[T any] struct {
	Weight int
	Value  T
}

// SelectWeighted selects a random element based on weights.
//
// Parameters:
//   - choices: slice of choices with weights
//
// Returns: randomly selected choice based on weights
//
// Example: SelectWeighted([]WeightedChoice[string]{{Weight: 1, Value: "A"}, {Weight: 2, Value: "B"}})
func SelectWeighted[T any](choices []WeightedChoice[T]) T {
	if len(choices) == 0 {
		panic("choices cannot be empty")
	}

	var total int
	for _, choice := range choices {
		if choice.Weight < 0 {
			panic("weights cannot be negative")
		}
		total += choice.Weight
	}

	if total == 0 {
		panic("total weight cannot be zero")
	}

	r := globalRand.Intn(total)
	cumulative := 0

	for _, choice := range choices {
		cumulative += choice.Weight
		if r < cumulative {
			return choice.Value
		}
	}

	// This should never happen, but just in case
	return choices[len(choices)-1].Value
}

// Percent returns true with the given percentage probability.
//
// Parameters:
//   - percentage: probability percentage (0-100)
//
// Returns: true with the given probability
//
// Example: Percent(30) returns true with 30% probability
func Percent(percentage int) bool {
	if percentage < 0 || percentage > 100 {
		panic("percentage must be between 0 and 100")
	}
	if percentage == 0 {
		return false
	}
	if percentage == 100 {
		return true
	}
	return globalRand.Intn(100) < percentage
}

// PercentFloat returns true with the given probability.
//
// Parameters:
//   - probability: probability (0.0-1.0)
//
// Returns: true with the given probability
//
// Example: PercentFloat(0.3) returns true with 30% probability
func PercentFloat(probability float64) bool {
	if probability < 0.0 || probability > 1.0 {
		panic("probability must be between 0.0 and 1.0")
	}
	if probability == 0.0 {
		return false
	}
	if probability == 1.0 {
		return true
	}
	return globalRand.Float64() < probability
}
