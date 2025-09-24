# prob - Go 概率与统计包

一个全面的 Go 概率计算、统计分布和随机采样操作包。

## 特性

- **基础概率**：基于百分比和浮点数的概率计算
- **权重选择**：基于权重的类型安全选择
- **统计分布**：二项式、泊松、正态、几何、超几何分布
- **随机生成**：均匀、正态、指数分布随机数
- **采样操作**：洗牌、采样、权重采样
- **性能优化**：高效算法，最小内存分配
- **类型安全**：支持任意数据类型的泛型函数

## 安装

```bash
go get github.com/go4x/goal/prob
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/prob"
)

func main() {
    // 基础概率
    if prob.Percent(30) {
        fmt.Println("30% 概率命中！")
    }
    
    // 权重选择
    weights := []int{1, 2, 3, 4}
    index, _ := prob.Select(weights)
    fmt.Printf("选择的索引: %d\n", index)
    
    // 统计分布
    probability := prob.Binomial(10, 3, 0.5)
    fmt.Printf("二项式概率: %.4f\n", probability)
}
```

## API 参考

### 基础概率

#### `Percent(percentage int) bool`
计算基于百分比的概率（0-100）。

```go
if prob.Percent(30) {
    // 30% 成功概率
}
```

#### `PercentFloat(probability float64) bool`
使用 float64 计算概率（0.0-1.0）。

```go
if prob.PercentFloat(0.3) {
    // 30% 成功概率
}
```

#### `Half() bool`
以 50% 概率返回 true。

```go
if prob.Half() {
    fmt.Println("正面")
} else {
    fmt.Println("反面")
}
```

### 权重选择

#### `Select(weights []int) (int, error)`
基于整数权重选择索引。

```go
weights := []int{1, 2, 3, 4}
index, err := prob.Select(weights)
if err != nil {
    // 处理错误
}
```

#### `SelectFloat(weights []float64) (int, error)`
基于 float64 权重选择索引。

```go
weights := []float64{0.1, 0.2, 0.3, 0.4}
index, err := prob.SelectFloat(weights)
```

#### `SelectSafe(weights []int) int`
安全版本，出错时返回 -1 而不是 panic。

```go
index := prob.SelectSafe(weights)
if index == -1 {
    // 处理错误
}
```

#### `SelectWeighted[T any](choices []WeightedChoice[T]) (T, error)`
从权重选择中类型安全地选择值。

```go
choices := []prob.WeightedChoice[string]{
    {Weight: 1, Value: "低"},
    {Weight: 2, Value: "中"},
    {Weight: 3, Value: "高"},
}

value, err := prob.SelectWeighted(choices)
```

#### `SelectWeightedFloat[T any](choices []WeightedChoiceFloat[T]) (T, error)`
从 float64 权重选择中类型安全地选择值。

```go
choices := []prob.WeightedChoiceFloat[string]{
    {Weight: 0.1, Value: "稀有"},
    {Weight: 0.3, Value: "常见"},
    {Weight: 0.6, Value: "非常常见"},
}

value, err := prob.SelectWeightedFloat(choices)
```

### 统计分布

#### `Binomial(n, k int, p float64) float64`
计算二项式概率。

```go
// 在 10 次试验中恰好 3 次成功的概率，成功概率为 0.5
prob := prob.Binomial(10, 3, 0.5)
```

#### `Poisson(k int, lambda float64) float64`
计算泊松概率。

```go
// k=3, λ=2.0 的泊松概率
prob := prob.Poisson(3, 2.0)
```

#### `Normal(x, mean, stdDev float64) float64`
计算正态分布概率密度。

```go
// x=0, 均值=0, 标准差=1 的正态密度
density := prob.Normal(0, 0, 1)
```

#### `Geometric(k int, p float64) float64`
计算几何概率。

```go
// 第 3 次试验首次成功的概率
prob := prob.Geometric(3, 0.3)
```

#### `Hypergeometric(n, k, K, N int) float64`
计算超几何概率。

```go
// 超几何概率
prob := prob.Hypergeometric(10, 3, 20, 100)
```

### 随机生成

#### `Uniform(min, max float64) float64`
从均匀分布生成随机数。

```go
// 0 到 10 之间的随机数
value := prob.Uniform(0, 10)
```

#### `NormalRandom(mean, stdDev float64) float64`
从正态分布生成随机数。

```go
// 均值=0, 标准差=1 的正态随机数
value := prob.NormalRandom(0, 1)
```

#### `Exponential(lambda float64) float64`
从指数分布生成随机数。

```go
// λ=1.0 的指数随机数
value := prob.Exponential(1.0)
```

### 采样操作

#### `Shuffle[T any](slice []T)`
使用 Fisher-Yates 算法原地洗牌。

```go
slice := []int{1, 2, 3, 4, 5}
prob.Shuffle(slice)
```

#### `Sample[T any](slice []T, k int) []T`
从切片中无重复地采样 k 个元素。

```go
slice := []string{"A", "B", "C", "D", "E"}
sampled := prob.Sample(slice, 3)
```

#### `WeightedSample[T any](slice []T, weights []int, k int) ([]T, error)`
基于权重采样 k 个元素。

```go
slice := []string{"A", "B", "C", "D"}
weights := []int{1, 2, 3, 4}
sampled, err := prob.WeightedSample(slice, weights, 2)
```

## 使用示例

### 基础概率

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/prob"
)

func main() {
    // 30% 成功概率
    if prob.Percent(30) {
        fmt.Println("成功！")
    }
    
    // 50% 概率
    if prob.Half() {
        fmt.Println("正面")
    } else {
        fmt.Println("反面")
    }
}
```

### 权重选择

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/prob"
)

func main() {
    // 从权重选择中选择
    choices := []prob.WeightedChoice[string]{
        {Weight: 1, Value: "普通"},
        {Weight: 2, Value: "不常见"},
        {Weight: 3, Value: "稀有"},
        {Weight: 4, Value: "史诗"},
        {Weight: 5, Value: "传说"},
    }
    
    value, err := prob.SelectWeighted(choices)
    if err != nil {
        fmt.Printf("错误: %v\n", err)
        return
    }
    
    fmt.Printf("选择的物品: %s\n", value)
}
```

### 统计分析

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/prob"
)

func main() {
    // 二项式分布
    n, k, p := 20, 5, 0.3
    binomialProb := prob.Binomial(n, k, p)
    fmt.Printf("二项式: P(X=%d) 在 %d 次试验中 = %.4f\n", k, n, binomialProb)
    
    // 泊松分布
    lambda := 3.0
    poissonProb := prob.Poisson(2, lambda)
    fmt.Printf("泊松: P(X=2) 当 λ=%.1f = %.4f\n", lambda, poissonProb)
    
    // 正态分布
    normalDensity := prob.Normal(0, 0, 1)
    fmt.Printf("正态密度在 0: %.4f\n", normalDensity)
}
```

### 随机采样

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/prob"
)

func main() {
    // 洗牌切片
    slice := []int{1, 2, 3, 4, 5}
    fmt.Printf("原始: %v\n", slice)
    
    prob.Shuffle(slice)
    fmt.Printf("洗牌后: %v\n", slice)
    
    // 采样元素
    sampled := prob.Sample(slice, 3)
    fmt.Printf("采样: %v\n", sampled)
}
```

### 游戏应用

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/prob"
)

func main() {
    // 战利品掉落系统
    lootChoices := []prob.WeightedChoice[string]{
        {Weight: 1, Value: "传说"},
        {Weight: 5, Value: "史诗"},
        {Weight: 15, Value: "稀有"},
        {Weight: 30, Value: "普通"},
        {Weight: 49, Value: "垃圾"},
    }
    
    // 模拟战利品掉落
    for i := 0; i < 5; i++ {
        loot, _ := prob.SelectWeighted(lootChoices)
        fmt.Printf("掉落 %d: %s\n", i+1, loot)
    }
    
    // 暴击系统
    critChance := 0.15 // 15% 暴击概率
    if prob.PercentFloat(critChance) {
        fmt.Println("暴击！")
    } else {
        fmt.Println("普通攻击")
    }
}
```

## 性能

该包针对性能进行了优化：

- **全局随机数生成器**：单一共享随机数生成器
- **高效算法**：Fisher-Yates 洗牌，优化采样
- **最小分配**：尽可能减少内存分配
- **类型安全**：泛型函数提供编译时类型检查

### 基准测试结果

- **基础操作**：~50-100ns/op，0 分配
- **权重选择**：~200-300ns/op，0-1 分配
- **统计计算**：~500ns-2μs/op，0 分配
- **采样操作**：~1-5μs/op，1-3 分配

## 测试

运行测试：

```bash
go test ./prob
```

运行覆盖率测试：

```bash
go test ./prob -cover
```

运行示例：

```bash
go test ./prob -run Example
```

运行基准测试：

```bash
go test ./prob -bench=.
```

## 许可证

此包是 `goal` 项目的一部分，遵循相同的许可证条款。

## 贡献

欢迎贡献！请随时提交 Pull Request。
