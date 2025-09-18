package constraint

import (
	"testing"
)

// 测试 Number 约束接口
func TestNumberConstraint(t *testing.T) {
	// 测试整数类型
	testIntegerTypes(t)

	// 测试浮点数类型
	testFloatTypes(t)

	// 测试不兼容的类型
	testIncompatibleTypes(t)
}

func testIntegerTypes(t *testing.T) {
	// 测试各种整数类型
	testNumberConstraint[int](t, "int")
	testNumberConstraint[int8](t, "int8")
	testNumberConstraint[int16](t, "int16")
	testNumberConstraint[int32](t, "int32")
	testNumberConstraint[int64](t, "int64")

	testNumberConstraint[uint](t, "uint")
	testNumberConstraint[uint8](t, "uint8")
	testNumberConstraint[uint16](t, "uint16")
	testNumberConstraint[uint32](t, "uint32")
	testNumberConstraint[uint64](t, "uint64")

	testNumberConstraint[uintptr](t, "uintptr")
}

func testFloatTypes(t *testing.T) {
	// 测试浮点数类型
	testNumberConstraint[float32](t, "float32")
	testNumberConstraint[float64](t, "float64")
}

func testIncompatibleTypes(t *testing.T) {
	// 这些类型不应该满足 Number 约束
	// 注意：这些测试在编译时就会失败，所以用注释形式展示

	// 以下类型不满足 Number 约束：
	// - string
	// - bool
	// - complex64
	// - complex128
	// - []int
	// - map[string]int
	// - struct{}

	t.Log("不兼容的类型（如 string, bool 等）在编译时就会被拒绝")
}

// 泛型函数，用于测试 Number 约束
func testNumberConstraint[T Number](t *testing.T, typeName string) {
	t.Run("Test_"+typeName, func(t *testing.T) {
		// 测试基本数值操作
		var zero T
		var one T = 1
		var two T = 2

		// 测试零值
		if zero != 0 {
			t.Errorf("%s 零值应该是 0，实际是 %v", typeName, zero)
		}

		// 测试加法
		result := one + one
		if result != two {
			t.Errorf("%s 加法测试失败：1 + 1 = %v，期望 %v", typeName, result, two)
		}

		// 测试乘法
		result = one * two
		if result != two {
			t.Errorf("%s 乘法测试失败：1 * 2 = %v，期望 %v", typeName, result, two)
		}

		// 测试比较
		if one >= two {
			t.Errorf("%s 比较测试失败：1 不应该大于等于 2", typeName)
		}

		if two <= one {
			t.Errorf("%s 比较测试失败：2 不应该小于等于 1", typeName)
		}

		t.Logf("%s 类型约束测试通过", typeName)
	})
}

// 测试 Number 约束在函数参数中的使用
func TestNumberConstraintInFunction(t *testing.T) {
	// 测试整数
	testNumberFunction(t, 42, "int")
	testNumberFunction(t, int8(42), "int8")
	testNumberFunction(t, int16(42), "int16")
	testNumberFunction(t, int32(42), "int32")
	testNumberFunction(t, int64(42), "int64")

	// 测试无符号整数
	testNumberFunction(t, uint(42), "uint")
	testNumberFunction(t, uint8(42), "uint8")
	testNumberFunction(t, uint16(42), "uint16")
	testNumberFunction(t, uint32(42), "uint32")
	testNumberFunction(t, uint64(42), "uint64")

	// 测试浮点数
	testNumberFunction(t, float32(42.5), "float32")
	testNumberFunction(t, float64(42.5), "float64")
}

func testNumberFunction[T Number](t *testing.T, value T, typeName string) {
	result := processNumber(value)
	expected := value * 2

	if result != expected {
		t.Errorf("%s 函数处理失败：输入 %v，输出 %v，期望 %v",
			typeName, value, result, expected)
	}

	t.Logf("%s 函数处理测试通过：%v -> %v", typeName, value, result)
}

// 使用 Number 约束的泛型函数
func processNumber[T Number](value T) T {
	return value * 2
}

// 测试 Number 约束在结构体中的使用
func TestNumberConstraintInStruct(t *testing.T) {
	// 测试整数容器
	intContainer := NewNumberContainer(42)
	if intContainer.Value != 42 {
		t.Errorf("整数容器值错误：期望 42，实际 %v", intContainer.Value)
	}

	// 测试浮点数容器
	floatContainer := NewNumberContainer(3.14)
	if floatContainer.Value != 3.14 {
		t.Errorf("浮点数容器值错误：期望 3.14，实际 %v", floatContainer.Value)
	}

	// 测试容器方法
	intContainer.Double()
	if intContainer.Value != 84 {
		t.Errorf("整数容器翻倍后值错误：期望 84，实际 %v", intContainer.Value)
	}

	floatContainer.Double()
	if floatContainer.Value != 6.28 {
		t.Errorf("浮点数容器翻倍后值错误：期望 6.28，实际 %v", floatContainer.Value)
	}
}

// 使用 Number 约束的泛型结构体
type NumberContainer[T Number] struct {
	Value T
}

func NewNumberContainer[T Number](value T) *NumberContainer[T] {
	return &NumberContainer[T]{Value: value}
}

func (nc *NumberContainer[T]) Double() {
	nc.Value *= 2
}

func (nc *NumberContainer[T]) Add(other T) {
	nc.Value += other
}

func (nc *NumberContainer[T]) GetValue() T {
	return nc.Value
}

// 测试 Number 约束在切片中的使用
func TestNumberConstraintInSlice(t *testing.T) {
	// 测试整数切片
	intSlice := []int{1, 2, 3, 4, 5}
	result := sumNumbers(intSlice)
	expected := 15
	if result != expected {
		t.Errorf("整数切片求和错误：期望 %d，实际 %d", expected, result)
	}

	// 测试浮点数切片
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	resultFloat := sumNumbers(floatSlice)
	expectedFloat := 16.5
	if resultFloat != expectedFloat {
		t.Errorf("浮点数切片求和错误：期望 %f，实际 %f", expectedFloat, resultFloat)
	}

	// 测试空切片
	emptySlice := []int{}
	resultEmpty := sumNumbers(emptySlice)
	if resultEmpty != 0 {
		t.Errorf("空切片求和错误：期望 0，实际 %v", resultEmpty)
	}
}

// 使用 Number 约束的泛型求和函数
func sumNumbers[T Number](slice []T) T {
	var sum T
	for _, value := range slice {
		sum += value
	}
	return sum
}

// 测试 Number 约束在比较函数中的使用
func TestNumberConstraintInComparison(t *testing.T) {
	// 测试整数比较
	testNumberComparison(t, 10, 20, "int")
	testNumberComparison(t, int32(10), int32(20), "int32")

	// 测试浮点数比较
	testNumberComparison(t, 10.5, 20.5, "float64")
	testNumberComparison(t, float32(10.5), float32(20.5), "float32")
}

func testNumberComparison[T Number](t *testing.T, a, b T, typeName string) {
	// 测试最大值
	max := maxNumber(a, b)
	if max != b {
		t.Errorf("%s 最大值错误：max(%v, %v) = %v，期望 %v", typeName, a, b, max, b)
	}

	// 测试最小值
	min := minNumber(a, b)
	if min != a {
		t.Errorf("%s 最小值错误：min(%v, %v) = %v，期望 %v", typeName, a, b, min, a)
	}

	t.Logf("%s 比较测试通过：max(%v, %v) = %v, min(%v, %v) = %v",
		typeName, a, b, max, a, b, min)
}

// 使用 Number 约束的泛型比较函数
func maxNumber[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func minNumber[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// 测试 Number 约束的边界情况
func TestNumberConstraintEdgeCases(t *testing.T) {
	// 测试零值
	testNumberEdgeCase(t, 0, "int zero")
	testNumberEdgeCase(t, 0.0, "float64 zero")

	// 测试负数
	testNumberEdgeCase(t, -1, "int negative")
	testNumberEdgeCase(t, -1.5, "float64 negative")

	// 测试大数
	testNumberEdgeCase(t, 999999, "int large")
	testNumberEdgeCase(t, 999999.999, "float64 large")
}

func testNumberEdgeCase[T Number](t *testing.T, value T, description string) {
	// 测试绝对值（对于有符号类型）
	if value < 0 {
		abs := -value
		if abs <= 0 {
			t.Errorf("%s 绝对值计算错误：abs(%v) = %v", description, value, abs)
		}
	}

	// 测试平方
	square := value * value
	if square < 0 {
		t.Errorf("%s 平方值不应该为负数：%v^2 = %v", description, value, square)
	}

	t.Logf("%s 边界情况测试通过：%v", description, value)
}
