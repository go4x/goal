package prob

import (
	"math"
	"testing"
)

func TestPercent(t *testing.T) {
	// Test edge cases
	if Percent(0) {
		t.Error("Percent(0) should return false")
	}
	if !Percent(100) {
		t.Error("Percent(100) should return true")
	}

	// Test panic cases
	defer func() {
		if r := recover(); r == nil {
			t.Error("Percent(-1) should panic")
		}
	}()
	Percent(-1)
}

func TestPercentFloat(t *testing.T) {
	// Test edge cases
	if PercentFloat(0.0) {
		t.Error("PercentFloat(0.0) should return false")
	}
	if !PercentFloat(1.0) {
		t.Error("PercentFloat(1.0) should return true")
	}

	// Test panic cases
	defer func() {
		if r := recover(); r == nil {
			t.Error("PercentFloat(-0.1) should panic")
		}
	}()
	PercentFloat(-0.1)
}

func TestHalf(t *testing.T) {
	// Test that Half returns both true and false
	trueCount := 0
	totalTests := 1000

	for i := 0; i < totalTests; i++ {
		if Half() {
			trueCount++
		}
	}

	// Should be roughly 50%
	if trueCount < 400 || trueCount > 600 {
		t.Errorf("Half() should return roughly 50%% true, got %d/%d", trueCount, totalTests)
	}
}

func TestSelect(t *testing.T) {
	// Test valid selection
	weights := []int{1, 2, 3, 4}
	index, err := Select(weights)
	if err != nil {
		t.Errorf("Select() error = %v, want nil", err)
	}
	if index < 0 || index >= len(weights) {
		t.Errorf("Select() index = %d, want 0 <= index < %d", index, len(weights))
	}

	// Test empty slice
	_, err = Select([]int{})
	if err == nil {
		t.Error("Select([]int{}) should return error")
	}

	// Test negative weights
	_, err = Select([]int{1, -1, 2})
	if err == nil {
		t.Error("Select() with negative weights should return error")
	}

	// Test zero total weight
	_, err = Select([]int{0, 0, 0})
	if err == nil {
		t.Error("Select() with zero total weight should return error")
	}
}

func TestSelectFloat(t *testing.T) {
	// Test valid selection
	weights := []float64{1.0, 2.0, 3.0, 4.0}
	index, err := SelectFloat(weights)
	if err != nil {
		t.Errorf("SelectFloat() error = %v, want nil", err)
	}
	if index < 0 || index >= len(weights) {
		t.Errorf("SelectFloat() index = %d, want 0 <= index < %d", index, len(weights))
	}

	// Test empty slice
	_, err = SelectFloat([]float64{})
	if err == nil {
		t.Error("SelectFloat([]float64{}) should return error")
	}
}

func TestSelectSafe(t *testing.T) {
	// Test valid selection
	weights := []int{1, 2, 3, 4}
	index := SelectSafe(weights)
	if index < 0 || index >= len(weights) {
		t.Errorf("SelectSafe() index = %d, want 0 <= index < %d", index, len(weights))
	}

	// Test invalid selection
	index = SelectSafe([]int{})
	if index != -1 {
		t.Errorf("SelectSafe([]int{}) = %d, want -1", index)
	}
}

func TestSelectWeighted(t *testing.T) {
	choices := []WeightedChoice[string]{
		{Weight: 1, Value: "A"},
		{Weight: 2, Value: "B"},
		{Weight: 3, Value: "C"},
	}

	value, err := SelectWeighted(choices)
	if err != nil {
		t.Errorf("SelectWeighted() error = %v, want nil", err)
	}
	if value != "A" && value != "B" && value != "C" {
		t.Errorf("SelectWeighted() value = %s, want A, B, or C", value)
	}

	// Test empty choices
	_, err = SelectWeighted([]WeightedChoice[string]{})
	if err == nil {
		t.Error("SelectWeighted() with empty choices should return error")
	}
}

func TestSelectWeightedFloat(t *testing.T) {
	choices := []WeightedChoiceFloat[string]{
		{Weight: 1.0, Value: "A"},
		{Weight: 2.0, Value: "B"},
		{Weight: 3.0, Value: "C"},
	}

	value, err := SelectWeightedFloat(choices)
	if err != nil {
		t.Errorf("SelectWeightedFloat() error = %v, want nil", err)
	}
	if value != "A" && value != "B" && value != "C" {
		t.Errorf("SelectWeightedFloat() value = %s, want A, B, or C", value)
	}
}

func TestBinomial(t *testing.T) {
	// Test edge cases
	if Binomial(0, 0, 0.5) != 1.0 {
		t.Error("Binomial(0, 0, 0.5) should return 1.0")
	}

	if Binomial(5, 0, 0.5) != math.Pow(0.5, 5) {
		t.Error("Binomial(5, 0, 0.5) should return 0.5^5")
	}

	// Test invalid parameters
	if Binomial(-1, 1, 0.5) != 0 {
		t.Error("Binomial(-1, 1, 0.5) should return 0")
	}

	if Binomial(5, 6, 0.5) != 0 {
		t.Error("Binomial(5, 6, 0.5) should return 0")
	}
}

func TestPoisson(t *testing.T) {
	// Test edge cases
	if Poisson(0, 0) != 1.0 {
		t.Error("Poisson(0, 0) should return 1.0")
	}

	if Poisson(1, 0) != 0.0 {
		t.Error("Poisson(1, 0) should return 0.0")
	}

	// Test invalid parameters
	if Poisson(-1, 1.0) != 0 {
		t.Error("Poisson(-1, 1.0) should return 0")
	}

	if Poisson(1, -1.0) != 0 {
		t.Error("Poisson(1, -1.0) should return 0")
	}
}

func TestNormal(t *testing.T) {
	// Test normal distribution properties
	mean := 0.0
	stdDev := 1.0

	// At the mean, should be highest
	atMean := Normal(mean, mean, stdDev)
	awayFromMean := Normal(mean+1, mean, stdDev)

	if atMean <= awayFromMean {
		t.Error("Normal distribution should be highest at the mean")
	}

	// Test invalid parameters
	if Normal(0, 0, 0) != 0 {
		t.Error("Normal(0, 0, 0) should return 0")
	}

	if Normal(0, 0, -1) != 0 {
		t.Error("Normal(0, 0, -1) should return 0")
	}
}

func TestUniform(t *testing.T) {
	// Test uniform distribution
	min := 0.0
	max := 10.0

	for i := 0; i < 1000; i++ {
		value := Uniform(min, max)
		if value < min || value >= max {
			t.Errorf("Uniform(%f, %f) = %f, want %f <= value < %f", min, max, value, min, max)
		}
	}

	// Test panic case
	defer func() {
		if r := recover(); r == nil {
			t.Error("Uniform(1, 1) should panic")
		}
	}()
	Uniform(1, 1)
}

func TestNormalRandom(t *testing.T) {
	// Test normal random generation
	mean := 0.0
	stdDev := 1.0

	// Generate many samples and check they're roughly normal
	samples := make([]float64, 1000)
	for i := range samples {
		samples[i] = NormalRandom(mean, stdDev)
	}

	// Calculate sample mean (should be close to 0)
	sum := 0.0
	for _, sample := range samples {
		sum += sample
	}
	sampleMean := sum / float64(len(samples))

	if math.Abs(sampleMean) > 0.2 {
		t.Errorf("NormalRandom() sample mean = %f, want close to 0", sampleMean)
	}

	// Test panic case
	defer func() {
		if r := recover(); r == nil {
			t.Error("NormalRandom(0, 0) should panic")
		}
	}()
	NormalRandom(0, 0)
}

func TestExponential(t *testing.T) {
	// Test exponential distribution
	lambda := 1.0

	for i := 0; i < 1000; i++ {
		value := Exponential(lambda)
		if value < 0 {
			t.Errorf("Exponential(%f) = %f, want >= 0", lambda, value)
		}
	}

	// Test panic case
	defer func() {
		if r := recover(); r == nil {
			t.Error("Exponential(0) should panic")
		}
	}()
	Exponential(0)
}

func TestGeometric(t *testing.T) {
	// Test geometric distribution
	p := 0.5

	// Sum of probabilities should be 1
	sum := 0.0
	for k := 1; k <= 100; k++ {
		sum += Geometric(k, p)
	}

	if math.Abs(sum-1.0) > 0.01 {
		t.Errorf("Geometric probabilities sum = %f, want close to 1.0", sum)
	}

	// Test invalid parameters
	if Geometric(0, 0.5) != 0 {
		t.Error("Geometric(0, 0.5) should return 0")
	}

	if Geometric(1, -0.1) != 0 {
		t.Error("Geometric(1, -0.1) should return 0")
	}
}

func TestHypergeometric(t *testing.T) {
	// Test hypergeometric distribution
	N, K, n, k := 100, 20, 10, 3

	prob := Hypergeometric(n, k, K, N)
	if prob < 0 || prob > 1 {
		t.Errorf("Hypergeometric(%d, %d, %d, %d) = %f, want 0 <= prob <= 1", n, k, K, N, prob)
	}

	// Test invalid parameters
	if Hypergeometric(-1, 1, 1, 1) != 0 {
		t.Error("Hypergeometric(-1, 1, 1, 1) should return 0")
	}
}

func TestShuffle(t *testing.T) {
	// Test shuffle preserves length
	slice := []int{1, 2, 3, 4, 5}
	original := make([]int, len(slice))
	copy(original, slice)

	Shuffle(slice)

	if len(slice) != len(original) {
		t.Error("Shuffle() should preserve slice length")
	}

	// Test that all elements are still present
	for _, elem := range original {
		found := false
		for _, shuffled := range slice {
			if elem == shuffled {
				found = true
				break
			}
		}
		if !found {
			t.Error("Shuffle() should preserve all elements")
		}
	}
}

func TestSample(t *testing.T) {
	// Test sample
	slice := []int{1, 2, 3, 4, 5}
	k := 3

	sampled := Sample(slice, k)

	if len(sampled) != k {
		t.Errorf("Sample() length = %d, want %d", len(sampled), k)
	}

	// Test invalid k
	if Sample(slice, 0) != nil {
		t.Error("Sample() with k=0 should return nil")
	}

	if Sample(slice, 10) != nil {
		t.Error("Sample() with k > len(slice) should return nil")
	}
}

func TestWeightedSample(t *testing.T) {
	// Test weighted sample
	slice := []string{"A", "B", "C", "D"}
	weights := []int{1, 2, 3, 4}
	k := 2

	sampled, err := WeightedSample(slice, weights, k)
	if err != nil {
		t.Errorf("WeightedSample() error = %v, want nil", err)
	}

	if len(sampled) != k {
		t.Errorf("WeightedSample() length = %d, want %d", len(sampled), k)
	}

	// Test invalid parameters
	_, err = WeightedSample(slice, weights, 0)
	if err == nil {
		t.Error("WeightedSample() with k=0 should return error")
	}

	_, err = WeightedSample(slice, weights, 10)
	if err == nil {
		t.Error("WeightedSample() with k > len(slice) should return error")
	}

	_, err = WeightedSample(slice, []int{1, 2}, 2)
	if err == nil {
		t.Error("WeightedSample() with mismatched lengths should return error")
	}
}

func TestLogCombination(t *testing.T) {
	// Test log combination
	n, k := 5, 2
	result := logCombination(n, k)

	// Should be positive
	if result <= 0 {
		t.Errorf("logCombination(%d, %d) = %f, want > 0", n, k, result)
	}

	// Test edge cases
	if logCombination(5, 0) != 0 {
		t.Error("logCombination(5, 0) should return 0")
	}

	if logCombination(5, 5) != 0 {
		t.Error("logCombination(5, 5) should return 0")
	}

	if !math.IsInf(logCombination(5, 6), -1) {
		t.Error("logCombination(5, 6) should return -Inf")
	}
}
