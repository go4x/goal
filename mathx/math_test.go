package mathx

import (
	"math"
	"testing"
)

func TestAddPrev(t *testing.T) {
	a := 3.123456789
	b := 2.123456789
	expected := 5.246913578
	result := Add(a, b)
	if result.Float64() != expected {
		t.Errorf("Add(%f, %f) = %f, want %f", a, b, result.Float64(), expected)
	}
}

func TestResult_Float64(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"positive number", 3.14, 3.14},
		{"negative number", -3.14, -3.14},
		{"zero", 0, 0},
		{"large number", 123456.789, 123456.789},
		{"small number", 0.000001, 0.000001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewResult(tt.value)
			if got := result.Float64(); got != tt.expected {
				t.Errorf("Result.Float64() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResult_String(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected string
	}{
		{"positive number", 3.14, "3.14"},
		{"negative number", -3.14, "-3.14"},
		{"zero", 0, "0"},
		{"integer", 42, "42"},
		{"decimal", 0.5, "0.5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewResult(tt.value)
			if got := result.String(); got != tt.expected {
				t.Errorf("Result.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResult_ToStringFixed(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		places   int32
		expected string
	}{
		{"2 decimal places", 3.14159, 2, "3.14"},
		{"4 decimal places", 3.14159, 4, "3.1416"},
		{"0 decimal places", 3.14159, 0, "3"},
		{"negative places", 123.456, -1, "120"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewResult(tt.value)
			if got := result.ToStringFixed(tt.places); got != tt.expected {
				t.Errorf("Result.ToStringFixed() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResult_Clean(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected string
	}{
		{"trailing zeros", 3.14000, "3.14"},
		{"no trailing zeros", 3.14, "3.14"},
		{"integer", 42.0, "42"},
		{"zero", 0.0, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewResult(tt.value)
			cleaned := result.Clean()
			if got := cleaned.String(); got != tt.expected {
				t.Errorf("Result.Clean() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResult_Round(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		places   int32
		expected string
	}{
		{"round up", 3.145, 2, "3.15"},
		{"round down", 3.144, 2, "3.14"},
		{"round to integer", 3.6, 0, "4"},
		{"negative places", 123.456, -1, "120"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewResult(tt.value)
			rounded := result.Round(tt.places)
			if got := rounded.String(); got != tt.expected {
				t.Errorf("Result.Round() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResult_Truncate(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		places   int32
		expected string
	}{
		{"truncate positive", 3.145, 2, "3.14"},
		{"truncate negative", -3.145, 2, "-3.14"},
		{"truncate to integer", 3.9, 0, "3"},
		{"negative places", 123.456, -1, "120"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewResult(tt.value)
			truncated := result.Truncate(tt.places)
			if got := truncated.String(); got != tt.expected {
				t.Errorf("Result.Truncate() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResult_FormatMoney(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		places   int32
		expected string
	}{
		{"thousands separator", 1234567.89, 2, "1,234,567.89"},
		{"no thousands separator", 123.45, 2, "123.45"},
		{"integer", 1000, 0, "1,000"},
		{"negative", -1234567.89, 2, "-1,234,567.89"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewResult(tt.value)
			if got := result.FormatMoney(tt.places); got != tt.expected {
				t.Errorf("Result.FormatMoney() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResult_Abs(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"positive number", 3.14, 3.14},
		{"negative number", -3.14, 3.14},
		{"zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewResult(tt.value)
			abs := result.Abs()
			if got := abs.Float64(); got != tt.expected {
				t.Errorf("Result.Abs() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResult_Neg(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"positive to negative", 3.14, -3.14},
		{"negative to positive", -3.14, 3.14},
		{"zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewResult(tt.value)
			neg := result.Neg()
			if got := neg.Float64(); got != tt.expected {
				t.Errorf("Result.Neg() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResult_Add(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		other    float64
		expected float64
	}{
		{"positive addition", 3.14, 2.86, 6.0},
		{"negative addition", -3.14, 2.86, -0.28},
		{"zero addition", 0, 5.5, 5.5},
		{"decimal precision", 0.1, 0.2, 0.3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewResult(tt.value)
			sum := result.Add(tt.other)
			if got := sum.Float64(); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Result.Add() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResult_Sub(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		other    float64
		expected float64
	}{
		{"positive subtraction", 5.0, 2.0, 3.0},
		{"negative result", 2.0, 5.0, -3.0},
		{"zero result", 5.0, 5.0, 0.0},
		{"decimal precision", 0.3, 0.1, 0.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewResult(tt.value)
			diff := result.Sub(tt.other)
			if got := diff.Float64(); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Result.Sub() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResult_Mul(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		other    float64
		expected float64
	}{
		{"positive multiplication", 3.0, 4.0, 12.0},
		{"negative multiplication", -3.0, 4.0, -12.0},
		{"zero multiplication", 5.0, 0, 0.0},
		{"decimal multiplication", 0.1, 0.2, 0.02},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewResult(tt.value)
			product := result.Mul(tt.other)
			if got := product.Float64(); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Result.Mul() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResult_Div(t *testing.T) {
	tests := []struct {
		name      string
		value     float64
		other     float64
		precision int32
		expected  float64
	}{
		{"simple division", 10.0, 2.0, 2, 5.0},
		{"decimal division", 1.0, 3.0, 2, 0.33},
		{"negative division", -10.0, 2.0, 2, -5.0},
		{"precision test", 1.0, 3.0, 4, 0.3333},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewResult(tt.value)
			quotient := result.Div(tt.other, tt.precision)
			if got := quotient.Float64(); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Result.Div() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"positive addition", 3.14, 2.86, 6.0},
		{"decimal precision", 0.1, 0.2, 0.3},
		{"negative addition", -3.14, 2.86, -0.28},
		{"zero addition", 0, 5.5, 5.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if got := result.Float64(); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Add() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"positive subtraction", 5.0, 2.0, 3.0},
		{"negative result", 2.0, 5.0, -3.0},
		{"zero result", 5.0, 5.0, 0.0},
		{"decimal precision", 0.3, 0.1, 0.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sub(tt.a, tt.b)
			if got := result.Float64(); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Sub() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"positive multiplication", 3.0, 4.0, 12.0},
		{"negative multiplication", -3.0, 4.0, -12.0},
		{"zero multiplication", 5.0, 0, 0.0},
		{"decimal multiplication", 0.1, 0.2, 0.02},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Mul(tt.a, tt.b)
			if got := result.Float64(); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Mul() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		precision int32
		expected  float64
	}{
		{"simple division", 10.0, 2.0, 2, 5.0},
		{"decimal division", 1.0, 3.0, 2, 0.33},
		{"negative division", -10.0, 2.0, 2, -5.0},
		{"precision test", 1.0, 3.0, 4, 0.3333},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Div(tt.a, tt.b, tt.precision)
			if got := result.Float64(); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Div() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDivTrunc(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		precision int32
		expected  float64
	}{
		{"simple truncation", 10.0, 3.0, 2, 3.33},
		{"negative truncation", -10.0, 3.0, 2, -3.33},
		{"zero precision", 10.0, 3.0, 0, 3.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DivTrunc(tt.a, tt.b, tt.precision)
			if got := result.Float64(); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("DivTrunc() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		places   int32
		expected float64
	}{
		{"round up", 3.145, 2, 3.15},
		{"round down", 3.144, 2, 3.14},
		{"round to integer", 3.6, 0, 4.0},
		{"negative places", 123.456, -1, 120.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Round(tt.value, tt.places)
			if got := result.Float64(); got != tt.expected {
				t.Errorf("Round() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		places   int32
		expected float64
	}{
		{"truncate positive", 3.145, 2, 3.14},
		{"truncate negative", -3.145, 2, -3.14},
		{"truncate to integer", 3.9, 0, 3.0},
		{"negative places", 123.456, -1, 120.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Truncate(tt.value, tt.places)
			if got := result.Float64(); got != tt.expected {
				t.Errorf("Truncate() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"positive number", 3.14, 3.14},
		{"negative number", -3.14, 3.14},
		{"zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.value); got != tt.expected {
				t.Errorf("Abs() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCeil(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"positive decimal", 3.2, 4.0},
		{"negative decimal", -3.2, -3.0},
		{"integer", 3.0, 3.0},
		{"negative integer", -3.0, -3.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ceil(tt.value); got != tt.expected {
				t.Errorf("Ceil() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"positive decimal", 3.8, 3.0},
		{"negative decimal", -3.8, -4.0},
		{"integer", 3.0, 3.0},
		{"negative integer", -3.0, -3.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Floor(tt.value); got != tt.expected {
				t.Errorf("Floor() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		name     string
		base     float64
		exponent float64
		expected float64
	}{
		{"positive power", 2.0, 3.0, 8.0},
		{"negative power", 2.0, -1.0, 0.5},
		{"zero power", 5.0, 0.0, 1.0},
		{"fractional power", 4.0, 0.5, 2.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pow(tt.base, tt.exponent); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Pow() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"perfect square", 16.0, 4.0},
		{"decimal square", 2.25, 1.5},
		{"zero", 0.0, 0.0},
		{"negative", -4.0, 0.0}, // Should return 0 for negative
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sqrt(tt.value); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Sqrt() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsZero(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected bool
	}{
		{"exact zero", 0.0, true},
		{"very small positive", 1e-11, true},
		{"very small negative", -1e-11, true},
		{"normal positive", 0.1, false},
		{"normal negative", -0.1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZero(tt.value); got != tt.expected {
				t.Errorf("IsZero() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsEqual(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected bool
	}{
		{"exact equal", 3.14, 3.14, true},
		{"very close", 3.14, 3.1400000001, true},
		{"different", 3.14, 3.15, false},
		{"zero comparison", 0.0, 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEqual(tt.a, tt.b); got != tt.expected {
				t.Errorf("IsEqual() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		min      float64
		max      float64
		expected float64
	}{
		{"within range", 5.0, 0.0, 10.0, 5.0},
		{"below min", -5.0, 0.0, 10.0, 0.0},
		{"above max", 15.0, 0.0, 10.0, 10.0},
		{"at min", 0.0, 0.0, 10.0, 0.0},
		{"at max", 10.0, 0.0, 10.0, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clamp(tt.value, tt.min, tt.max); got != tt.expected {
				t.Errorf("Clamp() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestLerp(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		t        float64
		expected float64
	}{
		{"t=0", 0.0, 10.0, 0.0, 0.0},
		{"t=1", 0.0, 10.0, 1.0, 10.0},
		{"t=0.5", 0.0, 10.0, 0.5, 5.0},
		{"t=0.3", 0.0, 10.0, 0.3, 3.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lerp(tt.a, tt.b, tt.t); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Lerp() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		name     string
		values   []float64
		expected float64
	}{
		{"simple average", []float64{1, 2, 3, 4, 5}, 3.0},
		{"decimal average", []float64{1.5, 2.5, 3.5}, 2.5},
		{"single value", []float64{42.0}, 42.0},
		{"empty slice", []float64{}, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.values...); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Average() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		name     string
		values   []float64
		expected float64
	}{
		{"odd count", []float64{1, 2, 3, 4, 5}, 3.0},
		{"even count", []float64{1, 2, 3, 4}, 2.5},
		{"single value", []float64{42.0}, 42.0},
		{"empty slice", []float64{}, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.values...); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Median() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		name     string
		values   []float64
		expected float64
	}{
		{"simple case", []float64{1, 2, 3, 4, 5}, 1.5811388301},
		{"single value", []float64{42.0}, 0.0},
		{"empty slice", []float64{}, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardDeviation(tt.values...); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("StandardDeviation() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		values   []float64
		expected float64
	}{
		{"positive numbers", []float64{1, 5, 3, 9, 2}, 9.0},
		{"negative numbers", []float64{-1, -5, -3, -9, -2}, -1.0},
		{"mixed numbers", []float64{-1, 5, -3, 9, -2}, 9.0},
		{"single value", []float64{42.0}, 42.0},
		{"empty slice", []float64{}, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.values...); got != tt.expected {
				t.Errorf("Max() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		values   []float64
		expected float64
	}{
		{"positive numbers", []float64{1, 5, 3, 9, 2}, 1.0},
		{"negative numbers", []float64{-1, -5, -3, -9, -2}, -9.0},
		{"mixed numbers", []float64{-1, 5, -3, 9, -2}, -3.0},
		{"single value", []float64{42.0}, 42.0},
		{"empty slice", []float64{}, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.values...); got != tt.expected {
				t.Errorf("Min() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		values   []float64
		expected float64
	}{
		{"positive numbers", []float64{1, 2, 3, 4, 5}, 15.0},
		{"negative numbers", []float64{-1, -2, -3, -4, -5}, -15.0},
		{"mixed numbers", []float64{1, -2, 3, -4, 5}, 3.0},
		{"single value", []float64{42.0}, 42.0},
		{"empty slice", []float64{}, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.values...); got != tt.expected {
				t.Errorf("Sum() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPercentage(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		percent  float64
		expected float64
	}{
		{"50 percent", 100.0, 50.0, 50.0},
		{"25 percent", 200.0, 25.0, 50.0},
		{"100 percent", 100.0, 100.0, 100.0},
		{"0 percent", 100.0, 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Percentage(tt.value, tt.percent); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("Percentage() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCompoundInterest(t *testing.T) {
	tests := []struct {
		name      string
		principal float64
		rate      float64
		periods   int
		expected  float64
	}{
		{"simple case", 1000.0, 0.1, 1, 1100.0},
		{"zero periods", 1000.0, 0.1, 0, 1000.0},
		{"zero rate", 1000.0, 0.0, 2, 1000.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompoundInterest(tt.principal, tt.rate, tt.periods); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("CompoundInterest() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSafeDiv(t *testing.T) {
	tests := []struct {
		name      string
		dividend  float64
		divisor   float64
		precision int32
		expected  float64
	}{
		{"normal division", 10.0, 2.0, 2, 5.0},
		{"division by zero", 10.0, 0.0, 2, 0.0},
		{"decimal division", 1.0, 3.0, 2, 0.33},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SafeDiv(tt.dividend, tt.divisor, tt.precision); math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("SafeDiv() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFormatCurrency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		places   int32
		expected string
	}{
		{"simple currency", 123.45, 2, "123.45"},
		{"zero places", 123.45, 0, "123"},
		{"negative amount", -123.45, 2, "-123.45"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCurrency(tt.amount, tt.places); got != tt.expected {
				t.Errorf("FormatCurrency() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestParseFloat(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  float64
		shouldErr bool
	}{
		{"valid decimal", "3.14", 3.14, false},
		{"valid integer", "42", 42.0, false},
		{"valid negative", "-3.14", -3.14, false},
		{"invalid string", "abc", 0.0, true},
		{"empty string", "", 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFloat(tt.input)
			if (err != nil) != tt.shouldErr {
				t.Errorf("ParseFloat() error = %v, wantErr %v", err, tt.shouldErr)
				return
			}
			if !tt.shouldErr && math.Abs(got-tt.expected) > 1e-10 {
				t.Errorf("ParseFloat() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestToFixed(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		places   int32
		expected float64
	}{
		{"2 decimal places", 3.14159, 2, 3.14},
		{"4 decimal places", 3.14159, 4, 3.1416},
		{"0 decimal places", 3.14159, 0, 3.0},
		{"negative places", 123.456, -1, 120.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFixed(tt.value, tt.places); got != tt.expected {
				t.Errorf("ToFixed() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsPositive(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected bool
	}{
		{"positive number", 3.14, true},
		{"negative number", -3.14, false},
		{"zero", 0.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPositive(tt.value); got != tt.expected {
				t.Errorf("IsPositive() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsNegative(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected bool
	}{
		{"positive number", 3.14, false},
		{"negative number", -3.14, true},
		{"zero", 0.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNegative(tt.value); got != tt.expected {
				t.Errorf("IsNegative() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSign(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected int
	}{
		{"positive number", 3.14, 1},
		{"negative number", -3.14, -1},
		{"zero", 0.0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sign(tt.value); got != tt.expected {
				t.Errorf("Sign() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// Test chainable operations
func TestChainableOperations(t *testing.T) {
	// Test complex chainable operations
	result := Add(0.1, 0.2).
		Mul(10).
		Div(3, 2).
		Round(2).
		ToStringFixed(2)

	expected := "1.00"
	if result != expected {
		t.Errorf("Chainable operations = %v, want %v", result, expected)
	}
}

// Test money formatting
func TestMoneyFormatting(t *testing.T) {
	result := Mul(1000, 12.5).
		FormatMoney(2)

	expected := "12,500.00"
	if result != expected {
		t.Errorf("Money formatting = %v, want %v", result, expected)
	}
}

// Test clean operation
func TestCleanOperation(t *testing.T) {
	result := Add(0.1, 0.2).
		Clean().
		ToString()

	expected := "0.3"
	if result != expected {
		t.Errorf("Clean operation = %v, want %v", result, expected)
	}
}
