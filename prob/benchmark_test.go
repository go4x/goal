package prob

import (
	"testing"
)

func BenchmarkPercent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Percent(50)
	}
}

func BenchmarkPercentFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PercentFloat(0.5)
	}
}

func BenchmarkHalf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Half()
	}
}

func BenchmarkSelect(b *testing.B) {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Select(weights)
	}
}

func BenchmarkSelectFloat(b *testing.B) {
	weights := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectFloat(weights)
	}
}

func BenchmarkSelectSafe(b *testing.B) {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectSafe(weights)
	}
}

func BenchmarkSelectWeighted(b *testing.B) {
	choices := []WeightedChoice[string]{
		{Weight: 1, Value: "A"},
		{Weight: 2, Value: "B"},
		{Weight: 3, Value: "C"},
		{Weight: 4, Value: "D"},
		{Weight: 5, Value: "E"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectWeighted(choices)
	}
}

func BenchmarkSelectWeightedFloat(b *testing.B) {
	choices := []WeightedChoiceFloat[string]{
		{Weight: 1.0, Value: "A"},
		{Weight: 2.0, Value: "B"},
		{Weight: 3.0, Value: "C"},
		{Weight: 4.0, Value: "D"},
		{Weight: 5.0, Value: "E"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectWeightedFloat(choices)
	}
}

func BenchmarkBinomial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Binomial(20, 5, 0.3)
	}
}

func BenchmarkPoisson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Poisson(3, 2.0)
	}
}

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Normal(0, 0, 1)
	}
}

func BenchmarkUniform(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uniform(0, 10)
	}
}

func BenchmarkNormalRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalRandom(0, 1)
	}
}

func BenchmarkExponential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Exponential(1.0)
	}
}

func BenchmarkGeometric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Geometric(3, 0.3)
	}
}

func BenchmarkHypergeometric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hypergeometric(10, 3, 20, 100)
	}
}

func BenchmarkShuffle(b *testing.B) {
	slice := make([]int, 100)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Shuffle(slice)
	}
}

func BenchmarkSample(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sample(slice, 10)
	}
}

func BenchmarkWeightedSample(b *testing.B) {
	slice := make([]int, 100)
	weights := make([]int, 100)
	for i := range slice {
		slice[i] = i
		weights[i] = i + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WeightedSample(slice, weights, 10)
	}
}

func BenchmarkLogCombination(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logCombination(20, 5)
	}
}

// Memory allocation benchmarks
func BenchmarkPercentAlloc(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Percent(50)
	}
}

func BenchmarkSelectAlloc(b *testing.B) {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Select(weights)
	}
}

func BenchmarkSelectWeightedAlloc(b *testing.B) {
	choices := []WeightedChoice[string]{
		{Weight: 1, Value: "A"},
		{Weight: 2, Value: "B"},
		{Weight: 3, Value: "C"},
		{Weight: 4, Value: "D"},
		{Weight: 5, Value: "E"},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectWeighted(choices)
	}
}

func BenchmarkShuffleAlloc(b *testing.B) {
	slice := make([]int, 100)
	for i := range slice {
		slice[i] = i
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Shuffle(slice)
	}
}

func BenchmarkSampleAlloc(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sample(slice, 10)
	}
}

// Large dataset benchmarks
func BenchmarkSelectLarge(b *testing.B) {
	weights := make([]int, 1000)
	for i := range weights {
		weights[i] = i + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Select(weights)
	}
}

func BenchmarkSelectFloatLarge(b *testing.B) {
	weights := make([]float64, 1000)
	for i := range weights {
		weights[i] = float64(i + 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectFloat(weights)
	}
}

func BenchmarkShuffleLarge(b *testing.B) {
	slice := make([]int, 10000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Shuffle(slice)
	}
}

func BenchmarkSampleLarge(b *testing.B) {
	slice := make([]int, 10000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sample(slice, 100)
	}
}

// Complex probability calculations
func BenchmarkBinomialLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Binomial(100, 50, 0.5)
	}
}

func BenchmarkPoissonLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Poisson(50, 25.0)
	}
}

func BenchmarkHypergeometricLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hypergeometric(50, 25, 100, 200)
	}
}

// Distribution generation benchmarks
func BenchmarkNormalRandomLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalRandom(0, 1)
	}
}

func BenchmarkExponentialLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Exponential(1.0)
	}
}

func BenchmarkUniformLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uniform(0, 100)
	}
}
