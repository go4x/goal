# MathX - 高精度数学计算库

MathX 是一个基于 `shopspring/decimal` 库的高精度数学计算包，专为 Go 语言设计。它提供了精确的十进制运算、链式 API 调用和丰富的数学函数，解决了 Go 原生浮点数运算精度不足的问题。

## 特性

- 🎯 **高精度计算** - 基于 `shopspring/decimal` 库，避免浮点数精度问题
- 🔗 **链式 API** - 支持流畅的方法链调用
- 📊 **统计函数** - 提供平均值、中位数、标准差等统计计算
- 💰 **货币格式化** - 支持千位分隔符和货币格式化
- 🧮 **数学函数** - 包含幂运算、开方、取整等常用数学函数
- 🛡️ **类型安全** - 使用 Go 泛型确保类型安全
- 📝 **完整文档** - 提供详细的 API 文档和使用示例

## 安装

```bash
go get github.com/go4x/goal/mathx
```

## 快速开始

### 基本运算

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/mathx"
)

func main() {
    // 精确的加法运算
    result := mathx.Add(0.1, 0.2)
    fmt.Printf("0.1 + 0.2 = %s\n", result.ToString())
    // 输出: 0.1 + 0.2 = 0.3
    
    // 链式操作
    chainResult := mathx.Add(0.1, 0.2).
        Mul(10).
        Div(3, 2).
        Round(2).
        ToStringFixed(2)
    fmt.Printf("链式结果: %s\n", chainResult)
    // 输出: 链式结果: 1.00
}
```

### 货币格式化

```go
// 货币格式化
price := mathx.Mul(99.99, 1.15).
    Round(2).
    FormatMoney(2)
fmt.Printf("价格: $%s\n", price)
// 输出: 价格: $114.99

// 大数字格式化
amount := mathx.Mul(1000, 12.5).
    FormatMoney(2)
fmt.Printf("金额: $%s\n", amount)
// 输出: 金额: $12,500.00
```

### 统计计算

```go
// 统计计算
scores := []float64{85, 92, 78, 96, 88}

avg := mathx.Average(scores...)
median := mathx.Median(scores...)
std := mathx.StandardDeviation(scores...)

fmt.Printf("平均分: %.2f\n", avg)
fmt.Printf("中位数: %.2f\n", median)
fmt.Printf("标准差: %.2f\n", std)
// 输出:
// 平均分: 87.80
// 中位数: 88.00
// 标准差: 6.87
```

## API 参考

### 基本运算

#### 链式操作

```go
type Result struct {
    // 内部使用 decimal.Decimal 确保精度
}

// 创建结果对象
func NewResult(value float64) Result

// 获取 float64 值
func (r Result) Float64() float64

// 获取字符串表示
func (r Result) String() string
func (r Result) ToString() string

// 固定小数位数
func (r Result) ToStringFixed(places int32) string
func (r Result) ToStringBank(places int32) string

// 清理尾随零
func (r Result) Clean() Result

// 四舍五入和截断
func (r Result) Round(places int32) Result
func (r Result) Truncate(places int32) Result

// 货币格式化
func (r Result) FormatMoney(decimalPlaces int32) string

// 数学运算
func (r Result) Abs() Result
func (r Result) Neg() Result
func (r Result) Add(other float64) Result
func (r Result) Sub(other float64) Result
func (r Result) Mul(other float64) Result
func (r Result) Div(other float64, precision int32) Result
func (r Result) DivTrunc(other float64, precision int32) Result
```

#### 顶级函数

```go
// 基本运算
func Add(a, b float64) Result
func Sub(a, b float64) Result
func Mul(a, b float64) Result
func Div(a, b float64, precision int32) Result
func DivTrunc(a, b float64, precision int32) Result

// 数学函数
func Round(value float64, precision int32) Result
func Truncate(value float64, precision int32) Result
func Abs(value float64) float64
func Ceil(value float64) float64
func Floor(value float64) float64
func Pow(base, exponent float64) float64
func Sqrt(value float64) float64

// 比较函数
func IsZero(value float64) bool
func IsEqual(a, b float64) bool
func IsPositive(value float64) bool
func IsNegative(value float64) bool
func Sign(value float64) int

// 工具函数
func Clamp(value, min, max float64) float64
func Lerp(a, b, t float64) float64
func ToFixed(value float64, places int32) float64
```

### 统计函数

```go
// 统计计算
func Average[T constraints.Integer | constraints.Float](ns ...T) float64
func Median[T constraints.Integer | constraints.Float](ns ...T) float64
func StandardDeviation[T constraints.Integer | constraints.Float](ns ...T) float64

// 最值和求和
func Max[T constraints.Ordered](ns ...T) T
func Min[T constraints.Ordered](ns ...T) T
func Sum[T constraints.Integer | constraints.Float](ns ...T) T
```

### 金融计算

```go
// 百分比计算
func Percentage(value, percent float64) float64

// 复利计算
func CompoundInterest(principal, rate float64, periods int) float64

// 安全除法（除零保护）
func SafeDiv(dividend, divisor float64, precision int32) float64
```

### 格式化函数

```go
// 货币格式化
func FormatCurrency(amount float64, decimalPlaces int32) string
func FormatMoney(amount float64, decimalPlaces int32) string
func FormatMoneyInt(amount int64, decimalPlaces int32) string

// 字符串转换
func ToString(value float64) string
func ToStringFixed(value float64, places int32) string
func ToStringBank(value float64, places int32) string

// 清理函数
func RemoveTrailingZeros(value float64) string
func RemoveTrailingZerosFixed(value float64, places int32) string
func CleanFloat(value float64) float64
func CleanFloatString(value float64) string

// 解析函数
func ParseFloat(s string) (float64, error)
```

## 使用示例

### 1. 精确计算

```go
// 避免浮点数精度问题
result1 := 0.1 + 0.2  // 原生 Go: 0.30000000000000004
result2 := mathx.Add(0.1, 0.2).ToString()  // MathX: "0.3"
```

### 2. 链式操作

```go
// 复杂的链式计算
result := mathx.Add(100, 50).
    Mul(1.1).           // 增加 10%
    Div(3, 2).          // 除以 3，保留 2 位小数
    Round(2).           // 四舍五入到 2 位小数
    FormatMoney(2)      // 货币格式化

fmt.Printf("最终结果: $%s\n", result)
// 输出: 最终结果: $55.00
```

### 3. 统计分析

```go
// 学生成绩分析
scores := []float64{85, 92, 78, 96, 88, 91, 87, 89, 93, 86}

// 计算统计指标
avg := mathx.Average(scores...)
median := mathx.Median(scores...)
std := mathx.StandardDeviation(scores...)
maxScore := mathx.Max(scores...)
minScore := mathx.Min(scores...)

fmt.Printf("平均分: %.2f\n", avg)
fmt.Printf("中位数: %.2f\n", median)
fmt.Printf("标准差: %.2f\n", std)
fmt.Printf("最高分: %.0f\n", maxScore)
fmt.Printf("最低分: %.0f\n", minScore)
```

### 4. 金融计算

```go
// 计算小费
bill := 50.00
tipPercent := 18.0
tip := mathx.Percentage(bill, tipPercent)
total := mathx.Add(bill, tip)

fmt.Printf("账单: $%.2f\n", bill)
fmt.Printf("小费 (%.0f%%): $%.2f\n", tipPercent, tip)
fmt.Printf("总计: $%.2f\n", total.Float64())

// 复利计算
principal := 1000.0
rate := 0.05  // 5% 年利率
years := 10
finalAmount := mathx.CompoundInterest(principal, rate, years)

fmt.Printf("本金: $%.2f\n", principal)
fmt.Printf("年利率: %.1f%%\n", rate*100)
fmt.Printf("年限: %d 年\n", years)
fmt.Printf("最终金额: $%.2f\n", finalAmount)
```

### 5. 数据清理

```go
// 清理浮点数尾随零
value := 3.140000
cleanValue := mathx.RemoveTrailingZeros(value)
fmt.Printf("清理前: %f\n", value)
fmt.Printf("清理后: %s\n", cleanValue)
// 输出:
// 清理前: 3.140000
// 清理后: 3.14

// 使用链式操作清理
result := mathx.NewResult(3.140000).
    Clean().
    ToString()
fmt.Printf("链式清理: %s\n", result)
// 输出: 链式清理: 3.14
```

## 性能对比

MathX 使用高精度 decimal 库，虽然比原生 Go 数学运算稍慢，但提供了更高的精度：

```go
// 基准测试结果 (Apple M4 Pro)
BenchmarkAdd-12                     	 1658899	       737.8 ns/op
BenchmarkMul-12                     	 1613665	       734.8 ns/op
BenchmarkDiv-12                     	 4808594	       260.8 ns/op
BenchmarkChainableOperations-12     	 1000000	      1029 ns/op

// 原生 Go 运算对比
BenchmarkNativeAdd-12               	1000000000	         0.2595 ns/op
BenchmarkNativeMul-12               	1000000000	         0.2586 ns/op
```

## 测试覆盖率

- **测试覆盖率**: 66.7%
- **测试用例**: 包含单元测试、示例函数和基准测试
- **示例函数**: 提供完整的使用示例

## 依赖

- `github.com/shopspring/decimal` - 高精度十进制运算库
- `golang.org/x/exp/constraints` - Go 泛型约束

## 许可证

本项目采用 MIT 许可证。

## 贡献

欢迎提交 Issue 和 Pull Request！

## 更新日志

### v1.0.0
- 初始版本发布
- 支持基本数学运算
- 支持链式 API
- 支持统计函数
- 支持货币格式化
- 完整的测试覆盖
