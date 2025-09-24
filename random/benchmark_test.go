package random_test

import (
	"testing"

	"github.com/go4x/goal/random"
)

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Int(1000)
	}
}

func BenchmarkBetween(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Between(10, 1000)
	}
}

func BenchmarkFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Float64()
	}
}

func BenchmarkFloat64Between(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Float64Between(1.0, 100.0)
	}
}

func BenchmarkBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Bool()
	}
}

func BenchmarkChoice(b *testing.B) {
	slice := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		random.Choice(slice)
	}
}

func BenchmarkShuffle(b *testing.B) {
	slice := make([]int, 100)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		random.Shuffle(slice)
	}
}

func BenchmarkSample(b *testing.B) {
	slice := make([]string, 1000)
	for i := range slice {
		slice[i] = "item"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		random.Sample(slice, 10)
	}
}

func BenchmarkSelectWeighted(b *testing.B) {
	choices := []random.WeightedChoice[string]{
		{Weight: 1, Value: "A"},
		{Weight: 2, Value: "B"},
		{Weight: 3, Value: "C"},
		{Weight: 4, Value: "D"},
		{Weight: 5, Value: "E"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		random.SelectWeighted(choices)
	}
}

func BenchmarkPercent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Percent(50)
	}
}

func BenchmarkPercentFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.PercentFloat(0.5)
	}
}

// Memory allocation benchmarks
func BenchmarkIntAlloc(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		random.Int(1000)
	}
}

func BenchmarkChoiceAlloc(b *testing.B) {
	slice := []string{"A", "B", "C", "D", "E"}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		random.Choice(slice)
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
		random.Shuffle(slice)
	}
}

func BenchmarkSampleAlloc(b *testing.B) {
	slice := make([]string, 1000)
	for i := range slice {
		slice[i] = "item"
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		random.Sample(slice, 10)
	}
}

// Large dataset benchmarks
func BenchmarkShuffleLarge(b *testing.B) {
	slice := make([]int, 10000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		random.Shuffle(slice)
	}
}

func BenchmarkSampleLarge(b *testing.B) {
	slice := make([]string, 10000)
	for i := range slice {
		slice[i] = "item"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		random.Sample(slice, 100)
	}
}

func BenchmarkSelectWeightedLarge(b *testing.B) {
	choices := make([]random.WeightedChoice[string], 1000)
	for i := range choices {
		choices[i] = random.WeightedChoice[string]{
			Weight: i + 1,
			Value:  "item",
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		random.SelectWeighted(choices)
	}
}

// String generation benchmarks
func BenchmarkLowercase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Lowercase(10)
	}
}

func BenchmarkUppercase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Uppercase(10)
	}
}

func BenchmarkDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Digits(10)
	}
}

func BenchmarkSymbols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Symbols(10)
	}
}

func BenchmarkHexUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.HexUpper(10)
	}
}

func BenchmarkAlphanumericSymbols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.AlphanumericSymbols(10)
	}
}

func BenchmarkStrongPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.StrongPassword(12)
	}
}

func BenchmarkReadable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Readable(10)
	}
}

func BenchmarkShortID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.ShortID(8)
	}
}

func BenchmarkPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Password(12, true)
	}
}

func BenchmarkUsername(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Username(8)
	}
}

func BenchmarkEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Email(8)
	}
}

func BenchmarkToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Token(32)
	}
}

func BenchmarkColorHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.ColorHex()
	}
}

func BenchmarkColorRGB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.ColorRGB()
	}
}

func BenchmarkMACAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.MACAddress()
	}
}

func BenchmarkIPAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.IPAddress()
	}
}

func BenchmarkWeightedString(b *testing.B) {
	chars := []random.WeightedChar{
		{Char: 'a', Weight: 1},
		{Char: 'b', Weight: 2},
		{Char: 'c', Weight: 3},
	}
	for i := 0; i < b.N; i++ {
		random.WeightedString(chars, 10)
	}
}

func BenchmarkPatternString(b *testing.B) {
	pattern := "aAn"
	for i := 0; i < b.N; i++ {
		random.PatternString(pattern)
	}
}

// Large string generation benchmarks
func BenchmarkLargeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Alphanumeric(1000)
	}
}

func BenchmarkLargePassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Password(100, true)
	}
}

func BenchmarkLargeToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Token(1000)
	}
}

// Memory allocation benchmarks
func BenchmarkStringAlloc(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		random.Alphanumeric(100)
	}
}

func BenchmarkPasswordAlloc(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		random.Password(50, true)
	}
}

func BenchmarkTokenAlloc(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		random.Token(100)
	}
}
