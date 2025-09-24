package mathx

import (
	"math"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(0.1, 0.2)
	}
}

func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sub(0.3, 0.1)
	}
}

func BenchmarkMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mul(0.1, 0.2)
	}
}

func BenchmarkDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Div(1.0, 3.0, 2)
	}
}

func BenchmarkRound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Round(3.14159, 2)
	}
}

func BenchmarkTruncate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Truncate(3.14159, 2)
	}
}

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(-3.14)
	}
}

func BenchmarkCeil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ceil(3.14)
	}
}

func BenchmarkFloor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Floor(3.14)
	}
}

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pow(2.0, 3.0)
	}
}

func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sqrt(16.0)
	}
}

func BenchmarkAverage(b *testing.B) {
	values := []float64{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Average(values...)
	}
}

func BenchmarkMedian(b *testing.B) {
	values := []float64{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Median(values...)
	}
}

func BenchmarkStandardDeviation(b *testing.B) {
	values := []float64{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		StandardDeviation(values...)
	}
}

func BenchmarkMax(b *testing.B) {
	values := []float64{1, 5, 3, 9, 2}
	for i := 0; i < b.N; i++ {
		Max(values...)
	}
}

func BenchmarkMin(b *testing.B) {
	values := []float64{1, 5, 3, 9, 2}
	for i := 0; i < b.N; i++ {
		Min(values...)
	}
}

func BenchmarkSum(b *testing.B) {
	values := []float64{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(values...)
	}
}

func BenchmarkPercentage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Percentage(100.0, 15.0)
	}
}

func BenchmarkCompoundInterest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompoundInterest(1000.0, 0.1, 2)
	}
}

func BenchmarkSafeDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SafeDiv(10.0, 2.0, 2)
	}
}

func BenchmarkToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToString(3.14159)
	}
}

func BenchmarkToStringFixed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToStringFixed(3.14159, 2)
	}
}

func BenchmarkFormatMoney(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FormatMoney(1234567.89, 2)
	}
}

func BenchmarkParseFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseFloat("3.14159")
	}
}

func BenchmarkChainableOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(0.1, 0.2).
			Mul(10).
			Div(3, 2).
			Round(2).
			ToString()
	}
}

func BenchmarkResult_Float64(b *testing.B) {
	result := NewResult(3.14159)
	for i := 0; i < b.N; i++ {
		result.Float64()
	}
}

func BenchmarkResult_String(b *testing.B) {
	result := NewResult(3.14159)
	for i := 0; i < b.N; i++ {
		_ = result.String()
	}
}

func BenchmarkResult_ToStringFixed(b *testing.B) {
	result := NewResult(3.14159)
	for i := 0; i < b.N; i++ {
		result.ToStringFixed(2)
	}
}

func BenchmarkResult_Clean(b *testing.B) {
	result := NewResult(3.14000)
	for i := 0; i < b.N; i++ {
		result.Clean()
	}
}

func BenchmarkResult_Round(b *testing.B) {
	result := NewResult(3.14159)
	for i := 0; i < b.N; i++ {
		result.Round(2)
	}
}

func BenchmarkResult_Truncate(b *testing.B) {
	result := NewResult(3.14159)
	for i := 0; i < b.N; i++ {
		result.Truncate(2)
	}
}

func BenchmarkResult_FormatMoney(b *testing.B) {
	result := NewResult(1234567.89)
	for i := 0; i < b.N; i++ {
		result.FormatMoney(2)
	}
}

func BenchmarkResult_Add(b *testing.B) {
	result := NewResult(3.14)
	for i := 0; i < b.N; i++ {
		result.Add(2.0)
	}
}

func BenchmarkResult_Sub(b *testing.B) {
	result := NewResult(3.14)
	for i := 0; i < b.N; i++ {
		result.Sub(2.0)
	}
}

func BenchmarkResult_Mul(b *testing.B) {
	result := NewResult(3.14)
	for i := 0; i < b.N; i++ {
		result.Mul(2.0)
	}
}

func BenchmarkResult_Div(b *testing.B) {
	result := NewResult(3.14)
	for i := 0; i < b.N; i++ {
		result.Div(2.0, 2)
	}
}

func BenchmarkResult_Abs(b *testing.B) {
	result := NewResult(-3.14)
	for i := 0; i < b.N; i++ {
		result.Abs()
	}
}

func BenchmarkResult_Neg(b *testing.B) {
	result := NewResult(3.14)
	for i := 0; i < b.N; i++ {
		result.Neg()
	}
}

// Compare with native Go operations
func BenchmarkNativeAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = 0.1 + 0.2
	}
}

func BenchmarkNativeMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = 0.1 * 0.2
	}
}

func BenchmarkNativeDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = 1.0 / 3.0
	}
}

func BenchmarkNativeAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = math.Abs(-3.14)
	}
}

func BenchmarkNativeCeil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = math.Ceil(3.14)
	}
}

func BenchmarkNativeFloor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = math.Floor(3.14)
	}
}

func BenchmarkNativePow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = math.Pow(2.0, 3.0)
	}
}

func BenchmarkNativeSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = math.Sqrt(16.0)
	}
}
