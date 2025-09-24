package prob

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

// Global random number generator for better performance
var globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// Percent calculates the percentage probability.
// Returns true if the random number is less than or equal to the given percentage.
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

// PercentFloat calculates the percentage probability using float64.
// Returns true if the random number is less than the given probability (0.0 to 1.0).
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

// Half returns true with 50% probability.
func Half() bool {
	return globalRand.Intn(2) == 1
}

// Select selects an index from a slice based on the weights.
// Returns the index of the selected element or an error if weights are invalid.
func Select(weights []int) (int, error) {
	if len(weights) == 0 {
		return 0, errors.New("weights slice cannot be empty")
	}

	var total int
	for _, w := range weights {
		if w < 0 {
			return 0, errors.New("weights cannot be negative")
		}
		total += w
	}

	if total == 0 {
		return 0, errors.New("total weight cannot be zero")
	}

	r := globalRand.Intn(total)
	cumulative := 0

	for i, w := range weights {
		cumulative += w
		if r < cumulative {
			return i, nil
		}
	}

	// This should never happen, but just in case
	return len(weights) - 1, nil
}

// SelectFloat selects an index from a slice based on float64 weights.
func SelectFloat(weights []float64) (int, error) {
	if len(weights) == 0 {
		return 0, errors.New("weights slice cannot be empty")
	}

	var total float64
	for _, w := range weights {
		if w < 0 {
			return 0, errors.New("weights cannot be negative")
		}
		total += w
	}

	if total == 0 {
		return 0, errors.New("total weight cannot be zero")
	}

	r := globalRand.Float64() * total
	cumulative := 0.0

	for i, w := range weights {
		cumulative += w
		if r < cumulative {
			return i, nil
		}
	}

	return len(weights) - 1, nil
}

// SelectSafe is a safe version of Select that doesn't panic.
// Returns -1 if there's an error.
func SelectSafe(weights []int) int {
	index, err := Select(weights)
	if err != nil {
		return -1
	}
	return index
}

// WeightedChoice represents a choice with a weight and value.
type WeightedChoice[T any] struct {
	Weight int
	Value  T
}

// SelectWeighted selects a value from weighted choices.
func SelectWeighted[T any](choices []WeightedChoice[T]) (T, error) {
	if len(choices) == 0 {
		var zero T
		return zero, errors.New("choices slice cannot be empty")
	}

	weights := make([]int, len(choices))
	for i, choice := range choices {
		weights[i] = choice.Weight
	}

	index, err := Select(weights)
	if err != nil {
		var zero T
		return zero, err
	}

	return choices[index].Value, nil
}

// WeightedChoiceFloat represents a choice with a float64 weight and value.
type WeightedChoiceFloat[T any] struct {
	Weight float64
	Value  T
}

// SelectWeightedFloat selects a value from float64 weighted choices.
func SelectWeightedFloat[T any](choices []WeightedChoiceFloat[T]) (T, error) {
	if len(choices) == 0 {
		var zero T
		return zero, errors.New("choices slice cannot be empty")
	}

	weights := make([]float64, len(choices))
	for i, choice := range choices {
		weights[i] = choice.Weight
	}

	index, err := SelectFloat(weights)
	if err != nil {
		var zero T
		return zero, err
	}

	return choices[index].Value, nil
}

// Binomial calculates the probability of k successes in n trials with probability p.
//
// The binomial distribution describes the number of successes in a fixed number of independent trials,
// each with the same probability of success. It's widely used in quality control, medical research,
// and statistical analysis.
//
// Mathematical formula: P(X = k) = C(n,k) × p^k × (1-p)^(n-k)
// Where C(n,k) is the binomial coefficient "n choose k"
//
// Parameters:
//   - n: number of trials (must be non-negative)
//   - k: number of successes (must be 0 ≤ k ≤ n)
//   - p: probability of success in each trial (must be 0 ≤ p ≤ 1)
//
// Returns: probability of exactly k successes in n trials
//
// Example: Binomial(10, 3, 0.5) calculates the probability of exactly 3 heads in 10 coin flips
//
// Reference: https://en.wikipedia.org/wiki/Binomial_distribution
func Binomial(n int, k int, p float64) float64 {
	if n < 0 || k < 0 || k > n || p < 0 || p > 1 {
		return 0
	}

	// Use logarithms to avoid overflow when dealing with large numbers
	// log(P) = log(C(n,k)) + k*log(p) + (n-k)*log(1-p)
	logResult := 0.0
	for i := 0; i < k; i++ {
		logResult += math.Log(float64(n-i)) - math.Log(float64(i+1))
	}
	logResult += float64(k)*math.Log(p) + float64(n-k)*math.Log(1-p)

	return math.Exp(logResult)
}

// Poisson calculates the Poisson probability.
//
// The Poisson distribution describes the number of events occurring in a fixed interval of time or space,
// given a constant average rate of occurrence. It's commonly used for modeling rare events like
// phone calls to a call center, defects in manufacturing, or radioactive decay.
//
// Mathematical formula: P(X = k) = (λ^k × e^(-λ)) / k!
// Where λ (lambda) is the average rate of occurrence
//
// Parameters:
//   - k: number of events (must be non-negative)
//   - lambda: average rate of occurrence (must be non-negative)
//
// Returns: probability of exactly k events occurring
//
// Example: Poisson(3, 2.0) calculates the probability of exactly 3 events when the average is 2
//
// Reference: https://en.wikipedia.org/wiki/Poisson_distribution
func Poisson(k int, lambda float64) float64 {
	if k < 0 || lambda < 0 {
		return 0
	}

	if lambda == 0 {
		if k == 0 {
			return 1
		}
		return 0
	}

	// Use logarithms to avoid overflow when dealing with large numbers
	// log(P) = k*log(λ) - λ - log(k!)
	logResult := float64(k)*math.Log(lambda) - lambda
	for i := 1; i <= k; i++ {
		logResult -= math.Log(float64(i))
	}

	return math.Exp(logResult)
}

// Normal calculates the normal distribution probability density.
//
// The normal (Gaussian) distribution is the most important probability distribution in statistics.
// It describes many natural phenomena and is characterized by its bell-shaped curve.
// The distribution is symmetric about its mean and follows the 68-95-99.7 rule.
//
// Mathematical formula: f(x) = (1/(σ√(2π))) × e^(-0.5×((x-μ)/σ)²)
// Where μ is the mean and σ is the standard deviation
//
// Parameters:
//   - x: value at which to evaluate the density
//   - mean: mean of the distribution (μ)
//   - stdDev: standard deviation of the distribution (σ, must be positive)
//
// Returns: probability density at point x (NOT a probability - density values can be > 1)
//
// Example: Normal(0, 0, 1) calculates the density of standard normal distribution at 0
//
// Reference: https://en.wikipedia.org/wiki/Normal_distribution
func Normal(x, mean, stdDev float64) float64 {
	if stdDev <= 0 {
		return 0
	}

	coefficient := 1.0 / (stdDev * math.Sqrt(2*math.Pi))
	z := (x - mean) / stdDev
	exponent := -0.5 * z * z

	return coefficient * math.Exp(exponent)
}

// Uniform generates a random number from a uniform distribution.
func Uniform(min, max float64) float64 {
	if min >= max {
		panic("min must be less than max")
	}
	return min + globalRand.Float64()*(max-min)
}

// NormalRandom generates a random number from a normal distribution using Box-Muller transform.
//
// The Box-Muller transform is a method for generating pairs of independent standard
// normally distributed random numbers from uniform random numbers. This implementation
// generates a single normal random number with specified mean and standard deviation.
//
// Mathematical method: Box-Muller transform
// Z = sqrt(-2*ln(U1)) * cos(2π*U2) where U1, U2 are uniform [0,1]
//
// Parameters:
//   - mean: mean of the normal distribution (μ)
//   - stdDev: standard deviation of the normal distribution (σ, must be positive)
//
// Returns: random number from normal distribution with specified mean and stdDev
//
// Example: NormalRandom(100, 15) generates a random number with mean 100, stdDev 15
//
// Reference: https://en.wikipedia.org/wiki/Box%E2%80%93Muller_transform
func NormalRandom(mean, stdDev float64) float64 {
	if stdDev <= 0 {
		panic("standard deviation must be positive")
	}

	// Box-Muller transform: generates standard normal Z, then scales to N(μ,σ²)
	u1 := globalRand.Float64()
	u2 := globalRand.Float64()

	z0 := math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2)

	return mean + stdDev*z0
}

// Exponential generates a random number from an exponential distribution.
//
// The exponential distribution describes the time between events in a Poisson process.
// It has the memoryless property, meaning the probability of an event occurring
// in the next time interval is independent of how much time has already elapsed.
//
// Mathematical formula: f(x) = λ × e^(-λx) for x ≥ 0
// Where λ (lambda) is the rate parameter
//
// Parameters:
//   - lambda: rate parameter (must be positive)
//
// Returns: random number from exponential distribution
//
// Example: Exponential(2.0) generates a random waiting time with rate 2
//
// Reference: https://en.wikipedia.org/wiki/Exponential_distribution
func Exponential(lambda float64) float64 {
	if lambda <= 0 {
		panic("lambda must be positive")
	}
	return -math.Log(globalRand.Float64()) / lambda
}

// Geometric calculates the probability of the first success on the k-th trial.
//
// The geometric distribution describes the number of trials needed to get the first success
// in a sequence of independent Bernoulli trials. It has the memoryless property.
//
// Mathematical formula: P(X = k) = p × (1-p)^(k-1)
// Where p is the probability of success on each trial
//
// Parameters:
//   - k: trial number of first success (must be ≥ 1)
//   - p: probability of success on each trial (must be 0 ≤ p ≤ 1)
//
// Returns: probability that the first success occurs on the k-th trial
//
// Example: Geometric(3, 0.5) calculates the probability of first head on the 3rd coin flip
//
// Reference: https://en.wikipedia.org/wiki/Geometric_distribution
func Geometric(k int, p float64) float64 {
	if k < 1 || p < 0 || p > 1 {
		return 0
	}
	return p * math.Pow(1-p, float64(k-1))
}

// Hypergeometric calculates the hypergeometric probability.
//
// The hypergeometric distribution describes the probability of k successes in n draws
// from a finite population of size N containing exactly K successes, without replacement.
// It's used in quality control, sampling surveys, and statistical analysis.
//
// Mathematical formula: P(X = k) = C(K,k) × C(N-K,n-k) / C(N,n)
// Where C(a,b) is the binomial coefficient "a choose b"
//
// Parameters:
//   - n: number of draws (sample size)
//   - k: number of successes in the sample
//   - K: number of successes in the population
//   - N: total population size
//
// Returns: probability of exactly k successes in n draws without replacement
//
// Example: Hypergeometric(10, 3, 20, 100) calculates the probability of 3 successes
//
//	in 10 draws from a population of 100 with 20 successes
//
// Reference: https://en.wikipedia.org/wiki/Hypergeometric_distribution
func Hypergeometric(n, k, K, N int) float64 {
	if n < 0 || k < 0 || K < 0 || N < 0 || k > K || n > N || n-k > N-K {
		return 0
	}

	// Use logarithms to avoid overflow when dealing with large numbers
	// log(P) = log(C(K,k)) + log(C(N-K,n-k)) - log(C(N,n))
	logResult := 0.0

	// Calculate log(C(K,k) * C(N-K,n-k) / C(N,n))
	logResult += logCombination(K, k)
	logResult += logCombination(N-K, n-k)
	logResult -= logCombination(N, n)

	return math.Exp(logResult)
}

// logCombination calculates log(n choose k) using logarithms.
//
// This function computes the natural logarithm of the binomial coefficient C(n,k) = n!/(k!(n-k)!)
// using the identity: log(C(n,k)) = log(n!) - log(k!) - log((n-k)!)
// This approach avoids overflow issues when dealing with large factorials.
//
// Parameters:
//   - n: total number of items
//   - k: number of items to choose
//
// Returns: log(C(n,k)) or -∞ if k > n or k < 0
//
// Example: logCombination(10, 3) calculates log(C(10,3)) = log(120)
//
// Reference: https://en.wikipedia.org/wiki/Binomial_coefficient
func logCombination(n, k int) float64 {
	if k > n || k < 0 {
		return math.Inf(-1)
	}
	if k == 0 || k == n {
		return 0
	}

	// Use the identity: log(n choose k) = log(n!) - log(k!) - log((n-k)!)
	result := 0.0
	for i := 1; i <= n; i++ {
		result += math.Log(float64(i))
	}
	for i := 1; i <= k; i++ {
		result -= math.Log(float64(i))
	}
	for i := 1; i <= n-k; i++ {
		result -= math.Log(float64(i))
	}

	return result
}

// Shuffle shuffles a slice in place using Fisher-Yates algorithm.
//
// The Fisher-Yates shuffle (also known as the Knuth shuffle) is an algorithm for generating
// a random permutation of a finite sequence. It produces an unbiased permutation where
// every permutation is equally likely.
//
// Algorithm: For each position i from n-1 down to 1:
//  1. Generate a random integer j such that 0 ≤ j ≤ i
//  2. Swap the elements at positions i and j
//
// Parameters:
//   - slice: the slice to shuffle (modified in place)
//
// Returns: none (slice is modified in place)
//
// Example: Shuffle([]int{1,2,3,4,5}) randomly reorders the elements
//
// Reference: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
func Shuffle[T any](slice []T) {
	for i := len(slice) - 1; i > 0; i-- {
		j := globalRand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Sample samples k elements from a slice without replacement.
//
// This function implements reservoir sampling, an algorithm for randomly selecting
// k elements from a population of unknown size. It ensures each element has an
// equal probability of being selected.
//
// Algorithm: Reservoir sampling
// 1. Initialize reservoir with first k elements
// 2. For each element i from k to n-1:
//   - Generate random j in [0, i]
//   - If j < k, replace reservoir[j] with element i
//
// Parameters:
//   - slice: the slice to sample from
//   - k: number of elements to sample (must be 0 < k ≤ len(slice))
//
// Returns: slice containing k randomly selected elements, or nil if invalid parameters
//
// Example: Sample([]string{"A","B","C","D","E"}, 3) returns 3 random elements
//
// Reference: https://en.wikipedia.org/wiki/Reservoir_sampling
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

// WeightedSample samples k elements from a slice based on weights.
func WeightedSample[T any](slice []T, weights []int, k int) ([]T, error) {
	if k <= 0 || k > len(slice) || len(slice) != len(weights) {
		return nil, errors.New("invalid parameters")
	}

	result := make([]T, 0, k)
	remaining := make([]int, len(slice))
	copy(remaining, weights)
	indices := make([]int, len(slice))
	for i := range indices {
		indices[i] = i
	}

	for i := 0; i < k; i++ {
		index, err := Select(remaining)
		if err != nil {
			return nil, err
		}

		result = append(result, slice[indices[index]])

		// Remove the selected element
		remaining = append(remaining[:index], remaining[index+1:]...)
		indices = append(indices[:index], indices[index+1:]...)
	}

	return result, nil
}
