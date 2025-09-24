package prob_test

import (
	"fmt"

	"github.com/go4x/goal/prob"
)

func ExampleBinomial() {
	// Calculate probability of exactly 3 successes in 10 trials with 0.5 probability
	n, k, p := 10, 3, 0.5
	probability := prob.Binomial(n, k, p)
	fmt.Printf("Probability of %d successes in %d trials: %.4f\n", k, n, probability)
	// Output: Probability of 3 successes in 10 trials: 0.1172
}

func ExamplePoisson() {
	// Calculate Poisson probability
	k, lambda := 3, 2.0
	probability := prob.Poisson(k, lambda)
	fmt.Printf("Poisson probability for k=%d, λ=%.1f: %.4f\n", k, lambda, probability)
	// Output: Poisson probability for k=3, λ=2.0: 0.1804
}

func ExampleNormal() {
	// Calculate normal distribution probability density
	x, mean, stdDev := 0.0, 0.0, 1.0
	density := prob.Normal(x, mean, stdDev)
	fmt.Printf("Normal density at x=%.1f: %.4f\n", x, density)
	// Output: Normal density at x=0.0: 0.3989
}

func ExampleGeometric() {
	// Calculate geometric probability
	k, p := 3, 0.3
	probability := prob.Geometric(k, p)
	fmt.Printf("Geometric probability for first success on trial %d: %.4f\n", k, probability)
	// Output: Geometric probability for first success on trial 3: 0.1470
}

func ExampleHypergeometric() {
	// Calculate hypergeometric probability
	N, K, n, k := 100, 20, 10, 3
	probability := prob.Hypergeometric(n, k, K, N)
	fmt.Printf("Hypergeometric probability: %.4f\n", probability)
	// Output: Hypergeometric probability: 0.2092
}
