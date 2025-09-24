package mathx

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
)

// Result represents a calculation result with chainable methods
type Result struct {
	v decimal.Decimal
}

// NewResult creates a new Result from a float64
func NewResult(value float64) Result {
	return Result{v: decimal.NewFromFloat(value)}
}

// Float64 returns the float64 value
func (r Result) Float64() float64 {
	f, _ := r.v.Float64()
	return f
}

// String returns the string representation
func (r Result) String() string {
	return r.v.String()
}

// ToString returns the string representation
func (r Result) ToString() string {
	return r.v.String()
}

// ToStringFixed returns the string with fixed decimal places
func (r Result) ToStringFixed(places int32) string {
	return r.v.StringFixed(places)
}

// ToStringBank returns the string with banker's rounding
func (r Result) ToStringBank(places int32) string {
	return r.v.StringFixedBank(places)
}

// Clean removes trailing zeros and returns a new Result
func (r Result) Clean() Result {
	// 转换为字符串去除尾随零，再转回decimal
	str := r.v.String()
	if strings.Contains(str, ".") {
		str = strings.TrimRight(str, "0")
		str = strings.TrimRight(str, ".")
	}
	cleanValue, _ := decimal.NewFromString(str)
	return Result{v: cleanValue}
}

// Round rounds to specified precision and returns a new Result
func (r Result) Round(places int32) Result {
	return Result{v: r.v.Round(places)}
}

// Truncate truncates to specified precision and returns a new Result
func (r Result) Truncate(places int32) Result {
	if places < 0 {
		// For negative precision, truncate to integer places
		// e.g., precision -1 means truncate to tens place
		multiplier := decimal.NewFromFloat(math.Pow(10, float64(-places)))
		result := r.v.Div(multiplier).Truncate(0).Mul(multiplier)
		return Result{v: result}
	}
	return Result{v: r.v.Truncate(places)}
}

// FormatMoney formats as currency with thousands separator
func (r Result) FormatMoney(decimalPlaces int32) string {
	rounded := r.v.Round(decimalPlaces)
	str := rounded.StringFixed(decimalPlaces)

	// 分离整数和小数部分
	parts := strings.Split(str, ".")
	integerPart := parts[0]
	decimalPart := ""
	if len(parts) > 1 {
		decimalPart = "." + parts[1]
	}

	// 添加千位分隔符
	if len(integerPart) > 3 {
		var result strings.Builder
		for i, char := range integerPart {
			if i > 0 && (len(integerPart)-i)%3 == 0 {
				result.WriteString(",")
			}
			result.WriteRune(char)
		}
		integerPart = result.String()
	}

	return integerPart + decimalPart
}

// Abs returns the absolute value
func (r Result) Abs() Result {
	return Result{v: r.v.Abs()}
}

// Neg returns the negative value
func (r Result) Neg() Result {
	return Result{v: r.v.Neg()}
}

// Add adds another value to this result
func (r Result) Add(other float64) Result {
	return Result{v: r.v.Add(decimal.NewFromFloat(other))}
}

// Sub subtracts another value from this result
func (r Result) Sub(other float64) Result {
	return Result{v: r.v.Sub(decimal.NewFromFloat(other))}
}

// Mul multiplies this result by another value
func (r Result) Mul(other float64) Result {
	return Result{v: r.v.Mul(decimal.NewFromFloat(other))}
}

// Div divides this result by another value
func (r Result) Div(other float64, precision int32) Result {
	return Result{v: r.v.DivRound(decimal.NewFromFloat(other), precision)}
}

// DivTrunc truncates the division
func (r Result) DivTrunc(other float64, precision int32) Result {
	return Result{v: r.v.Div(decimal.NewFromFloat(other)).Truncate(precision)}
}

// Add adds two float64 values using decimal precision and returns a Result
func Add(a, b float64) Result {
	result := decimal.NewFromFloat(a).Add(decimal.NewFromFloat(b))
	return Result{v: result}
}

// Sub subtracts two float64 values using decimal precision and returns a Result
func Sub(a, b float64) Result {
	result := decimal.NewFromFloat(a).Sub(decimal.NewFromFloat(b))
	return Result{v: result}
}

// Mul multiplies two float64 values using decimal precision and returns a Result
func Mul(a, b float64) Result {
	result := decimal.NewFromFloat(a).Mul(decimal.NewFromFloat(b))
	return Result{v: result}
}

// Div divides two float64 values using decimal precision and returns a Result
func Div(a, b float64, precision int32) Result {
	result := decimal.NewFromFloat(a).DivRound(decimal.NewFromFloat(b), precision)
	return Result{v: result}
}

// DivTrunc truncates the division of two float64 values and returns a Result
func DivTrunc(a, b float64, precision int32) Result {
	result := decimal.NewFromFloat(a).Div(decimal.NewFromFloat(b)).Truncate(precision)
	return Result{v: result}
}

// Round rounds a float64 to specified precision and returns a Result
func Round(value float64, precision int32) Result {
	result := decimal.NewFromFloat(value).Round(precision)
	return Result{v: result}
}

// Truncate truncates a float64 to specified precision and returns a Result
func Truncate(value float64, precision int32) Result {
	if precision < 0 {
		// For negative precision, truncate to integer places
		// e.g., precision -1 means truncate to tens place
		multiplier := decimal.NewFromFloat(math.Pow(10, float64(-precision)))
		result := decimal.NewFromFloat(value).Div(multiplier).Truncate(0).Mul(multiplier)
		return Result{v: result}
	}
	result := decimal.NewFromFloat(value).Truncate(precision)
	return Result{v: result}
}

// Int64Div divides two int64 values with specified precision
func Int64Div(dividend, divisor int64, precision int32) float64 {
	result := decimal.NewFromInt(dividend).DivRound(decimal.NewFromInt(divisor), precision)
	f, _ := result.Float64()
	return f
}

// Int64DivTrunc truncates the division of two int64 values
func Int64DivTrunc(dividend, divisor int64, precision int32) float64 {
	result := decimal.NewFromInt(dividend).Div(decimal.NewFromInt(divisor)).Truncate(precision)
	f, _ := result.Float64()
	return f
}

// Int64MulFloat64 multiplies int64 and float64 using decimal precision
func Int64MulFloat64(multiplicand int64, multiplier float64) float64 {
	result := decimal.NewFromInt(multiplicand).Mul(decimal.NewFromFloat(multiplier))
	f, _ := result.Float64()
	return f
}

// BigFloatMul multiplies two big.Float values using decimal precision
func BigFloatMul(multiplicand, multiplier *big.Float) *big.Float {
	m1, _ := multiplicand.Float64()
	m2, _ := multiplier.Float64()
	result := Mul(m1, m2)
	return big.NewFloat(result.Float64())
}

// Percentage calculates percentage of a value
func Percentage(value, percent float64) float64 {
	return Mul(value, Div(percent, 100, 10).Float64()).Float64()
}

// CompoundInterest calculates compound interest
func CompoundInterest(principal, rate float64, periods int) float64 {
	if periods <= 0 {
		return principal
	}
	multiplier := Add(1, rate)
	power := multiplier
	for i := 1; i < periods; i++ {
		power = power.Mul(multiplier.Float64())
	}
	return Mul(principal, power.Float64()).Float64()
}

// SafeDiv safely divides two numbers, returns 0 if divisor is 0
func SafeDiv(dividend, divisor float64, precision int32) float64 {
	if divisor == 0 {
		return 0
	}
	return Div(dividend, divisor, precision).Float64()
}

// Max returns the maximum value from a slice of numbers
func Max[T constraints.Ordered](ns ...T) T {
	if len(ns) == 0 {
		var zero T
		return zero
	}
	max := ns[0]
	for _, n := range ns[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

// Min returns the minimum value from a slice of numbers
func Min[T constraints.Ordered](ns ...T) T {
	if len(ns) == 0 {
		var zero T
		return zero
	}
	min := ns[0]
	for _, n := range ns[1:] {
		if n < min {
			min = n
		}
	}
	return min
}

// Sum returns the sum of a slice of numbers
func Sum[T constraints.Integer | constraints.Float](ns ...T) T {
	var sum T
	for _, n := range ns {
		sum += n
	}
	return sum
}

// DecimalSum returns the sum of decimal values
func DecimalSum(ds ...decimal.Decimal) decimal.Decimal {
	if len(ds) == 0 {
		return decimal.Zero
	}
	sum := ds[0]
	for _, d := range ds[1:] {
		sum = sum.Add(d)
	}
	return sum
}

// DecimalMax returns the maximum decimal value
func DecimalMax(ds ...decimal.Decimal) decimal.Decimal {
	if len(ds) == 0 {
		return decimal.Zero
	}
	max := ds[0]
	for _, d := range ds[1:] {
		if d.GreaterThan(max) {
			max = d
		}
	}
	return max
}

// DecimalMin returns the minimum decimal value
func DecimalMin(ds ...decimal.Decimal) decimal.Decimal {
	if len(ds) == 0 {
		return decimal.Zero
	}
	min := ds[0]
	for _, d := range ds[1:] {
		if d.LessThan(min) {
			min = d
		}
	}
	return min
}

// Abs returns the absolute value of a number
func Abs(value float64) float64 {
	if value < 0 {
		return -value
	}
	return value
}

// Ceil returns the smallest integer greater than or equal to the value
func Ceil(value float64) float64 {
	result := decimal.NewFromFloat(value).Ceil()
	f, _ := result.Float64()
	return f
}

// Floor returns the largest integer less than or equal to the value
func Floor(value float64) float64 {
	result := decimal.NewFromFloat(value).Floor()
	f, _ := result.Float64()
	return f
}

// Pow raises a number to the power of another
func Pow(base, exponent float64) float64 {
	result := decimal.NewFromFloat(base).Pow(decimal.NewFromFloat(exponent))
	f, _ := result.Float64()
	return f
}

// Sqrt returns the square root of a number
func Sqrt(value float64) float64 {
	if value < 0 {
		return 0 // 负数返回0，或者可以返回错误
	}
	// 使用牛顿法计算平方根
	if value == 0 {
		return 0
	}

	// 初始猜测值
	guess := value / 2
	for i := 0; i < 10; i++ { // 迭代10次
		guess = Div(Add(guess, Div(value, guess, 10).Float64()).Float64(), 2, 10).Float64()
	}
	return guess
}

// IsZero checks if a number is zero (within a small epsilon)
func IsZero(value float64) bool {
	return Abs(value) < 1e-10
}

// IsEqual checks if two numbers are equal (within a small epsilon)
func IsEqual(a, b float64) bool {
	return Abs(a-b) < 1e-8
}

// Clamp clamps a value between min and max
func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Lerp performs linear interpolation between two values
func Lerp(a, b, t float64) float64 {
	return Add(a, Mul(Sub(b, a).Float64(), t).Float64()).Float64()
}

// Average calculates the average of a slice of numbers
func Average[T constraints.Integer | constraints.Float](ns ...T) float64 {
	if len(ns) == 0 {
		return 0
	}
	sum := Sum(ns...)
	return Div(float64(sum), float64(len(ns)), 10).Float64()
}

// Median calculates the median of a slice of numbers
func Median[T constraints.Integer | constraints.Float](ns ...T) float64 {
	if len(ns) == 0 {
		return 0
	}
	// 简单的冒泡排序
	for i := 0; i < len(ns)-1; i++ {
		for j := 0; j < len(ns)-i-1; j++ {
			if ns[j] > ns[j+1] {
				ns[j], ns[j+1] = ns[j+1], ns[j]
			}
		}
	}

	n := len(ns)
	if n%2 == 0 {
		mid1 := float64(ns[n/2-1])
		mid2 := float64(ns[n/2])
		return Div(Add(mid1, mid2).Float64(), 2, 10).Float64()
	}
	return float64(ns[n/2])
}

// StandardDeviation calculates the standard deviation of a slice of numbers
func StandardDeviation[T constraints.Integer | constraints.Float](ns ...T) float64 {
	if len(ns) == 0 {
		return 0
	}
	if len(ns) == 1 {
		return 0
	}

	avg := Average(ns...)
	var sum float64
	for _, n := range ns {
		diff := Sub(float64(n), avg).Float64()
		sum = Add(sum, Mul(diff, diff).Float64()).Float64()
	}

	// Use sample standard deviation (n-1)
	variance := Div(sum, float64(len(ns)-1), 10).Float64()
	return Sqrt(variance)
}

// FormatCurrency formats a number as currency with specified decimal places
func FormatCurrency(amount float64, decimalPlaces int32) string {
	rounded := Round(amount, decimalPlaces)
	if decimalPlaces == 0 {
		return fmt.Sprintf("%.0f", rounded.Float64())
	}
	format := fmt.Sprintf("%%.%df", decimalPlaces)
	return fmt.Sprintf(format, rounded.Float64())
}

// ParseFloat safely parses a string to float64
func ParseFloat(s string) (float64, error) {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return 0, err
	}
	f, _ := d.Float64()
	return f, nil
}

// ToFixed formats a number to a fixed number of decimal places
func ToFixed(value float64, places int32) float64 {
	return Round(value, places).Float64()
}

// IsPositive checks if a number is positive
func IsPositive(value float64) bool {
	return value > 0
}

// IsNegative checks if a number is negative
func IsNegative(value float64) bool {
	return value < 0
}

// Sign returns the sign of a number (-1, 0, or 1)
func Sign(value float64) int {
	if value > 0 {
		return 1
	}
	if value < 0 {
		return -1
	}
	return 0
}

// ToString converts a float64 to string
func ToString(value float64) string {
	return decimal.NewFromFloat(value).String()
}

// ToStringFixed converts a float64 to string with fixed decimal places
func ToStringFixed(value float64, places int32) string {
	return decimal.NewFromFloat(value).StringFixed(places)
}

// ToStringBank converts a float64 to string with banker's rounding
func ToStringBank(value float64, places int32) string {
	return decimal.NewFromFloat(value).StringFixedBank(places)
}

// FormatMoney formats a number as currency with thousands separator
func FormatMoney(amount float64, decimalPlaces int32) string {
	d := decimal.NewFromFloat(amount)
	rounded := d.Round(decimalPlaces)

	// 转换为字符串并添加千位分隔符
	str := rounded.String()

	// 分离整数和小数部分
	parts := strings.Split(str, ".")
	integerPart := parts[0]
	decimalPart := ""
	if len(parts) > 1 {
		decimalPart = "." + parts[1]
	}

	// 添加千位分隔符
	if len(integerPart) > 3 {
		var result strings.Builder
		for i, char := range integerPart {
			if i > 0 && (len(integerPart)-i)%3 == 0 {
				result.WriteString(",")
			}
			result.WriteRune(char)
		}
		integerPart = result.String()
	}

	return integerPart + decimalPart
}

// FormatMoneyInt formats an int64 as currency with thousands separator
func FormatMoneyInt(amount int64, decimalPlaces int32) string {
	d := decimal.NewFromInt(amount)
	rounded := d.Round(decimalPlaces)

	// 转换为字符串并添加千位分隔符
	str := rounded.String()

	// 分离整数和小数部分
	parts := strings.Split(str, ".")
	integerPart := parts[0]
	decimalPart := ""
	if len(parts) > 1 {
		decimalPart = "." + parts[1]
	}

	// 添加千位分隔符
	if len(integerPart) > 3 {
		var result strings.Builder
		for i, char := range integerPart {
			if i > 0 && (len(integerPart)-i)%3 == 0 {
				result.WriteString(",")
			}
			result.WriteRune(char)
		}
		integerPart = result.String()
	}

	return integerPart + decimalPart
}

// RemoveTrailingZeros removes trailing zeros from a float64 string representation
func RemoveTrailingZeros(value float64) string {
	str := decimal.NewFromFloat(value).String()

	// 如果包含小数点
	if strings.Contains(str, ".") {
		// 移除尾随的零
		str = strings.TrimRight(str, "0")
		// 如果小数点后没有数字了，也移除小数点
		str = strings.TrimRight(str, ".")
	}

	return str
}

// RemoveTrailingZerosFixed removes trailing zeros from a float64 with fixed decimal places
func RemoveTrailingZerosFixed(value float64, places int32) string {
	str := decimal.NewFromFloat(value).StringFixed(places)

	// 如果包含小数点
	if strings.Contains(str, ".") {
		// 移除尾随的零
		str = strings.TrimRight(str, "0")
		// 如果小数点后没有数字了，也移除小数点
		str = strings.TrimRight(str, ".")
	}

	return str
}

// CleanFloat removes trailing zeros and returns a clean float64
func CleanFloat(value float64) float64 {
	// 先转换为decimal，再转回float64，这样可以去除浮点数精度问题
	d := decimal.NewFromFloat(value)
	f, _ := d.Float64()
	return f
}

// CleanFloatString removes trailing zeros and returns a clean string representation
func CleanFloatString(value float64) string {
	d := decimal.NewFromFloat(value)
	return d.String()
}
