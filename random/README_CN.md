# random - Go 随机数生成包

一个全面的 Go 随机数生成、采样和概率操作包，具有高性能和类型安全特性。

## 特性

- **高性能**：全局随机数生成器，性能最优
- **类型安全**：支持任意数据类型的泛型函数
- **全面性**：整数、浮点数、布尔值和分布随机数
- **字符串生成**：多种字符集的随机字符串生成
- **采样**：洗牌、采样和加权选择操作
- **概率**：基于百分比的概率函数
- **统计**：正态分布和指数分布生成
- **安全性**：加密安全的随机数生成

## 安装

```bash
go get github.com/go4x/goal/random
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // 基础随机数
    intValue := random.Int(100)           // 随机整数 0-99
    floatValue := random.Float64()        // 随机浮点数 0.0-1.0
    boolValue := random.Bool()           // 随机布尔值
    
    // 范围生成
    rangeInt := random.Between(10, 50)    // 随机整数 10-49
    rangeFloat := random.Float64Between(1.5, 3.7) // 随机浮点数 1.5-3.7
    
    // 概率
    if random.Percent(30) {              // 30% 概率
        fmt.Println("成功!")
    }
    
    // 采样
    items := []string{"A", "B", "C", "D", "E"}
    choice := random.Choice(items)        // 随机选择
    sampled := random.Sample(items, 3)   // 采样 3 个元素
    
    fmt.Printf("随机值: %d, %.4f, %t\n", intValue, floatValue, boolValue)
    fmt.Printf("范围值: %d, %.4f\n", rangeInt, rangeFloat)
    fmt.Printf("选择: %s, 采样: %v\n", choice, sampled)
}
```

## API 参考

### 基础随机数

#### `Int(max int) int`
生成范围 [0, max) 内的随机整数。

```go
value := random.Int(100)  // 返回 0-99
```

#### `Between(min, max int) int`
生成范围 [min, max) 内的随机整数。

```go
value := random.Between(10, 50)  // 返回 10-49
```

#### `Float64() float64`
生成范围 [0.0, 1.0) 内的随机浮点数。

```go
value := random.Float64()  // 返回 0.0-0.999...
```

#### `Float64Between(min, max float64) float64`
生成范围 [min, max) 内的随机浮点数。

```go
value := random.Float64Between(1.5, 3.7)  // 返回 1.5-3.699...
```

#### `Bool() bool`
生成随机布尔值。

```go
value := random.Bool()  // 返回 true 或 false
```

### 选择和采样

#### `Choice[T any](slice []T) T`
从切片中随机选择一个元素。

```go
options := []string{"A", "B", "C"}
choice := random.Choice(options)  // 返回 "A", "B", 或 "C"
```

#### `Shuffle[T any](slice []T)`
使用 Fisher-Yates 算法原地洗牌。

```go
numbers := []int{1, 2, 3, 4, 5}
random.Shuffle(numbers)  // 原地洗牌
```

#### `Sample[T any](slice []T, k int) []T`
从切片中无放回地采样 k 个元素。

```go
items := []string{"A", "B", "C", "D", "E"}
sampled := random.Sample(items, 3)  // 返回 3 个随机元素
```

#### `SelectWeighted[T any](choices []WeightedChoice[T]) T`
基于权重随机选择一个元素。

```go
choices := []random.WeightedChoice[string]{
    {Weight: 1, Value: "普通"},
    {Weight: 2, Value: "稀有"},
    {Weight: 3, Value: "史诗"},
}
choice := random.SelectWeighted(choices)
```

### 概率函数

#### `Percent(percentage int) bool`
以给定百分比概率返回 true。

```go
if random.Percent(30) {  // 30% 概率
    fmt.Println("成功!")
}
```

#### `PercentFloat(probability float64) bool`
以给定概率返回 true。

```go
if random.PercentFloat(0.3) {  // 30% 概率
    fmt.Println("成功!")
}
```

### 统计分布

#### `Normal(mean, stdDev float64) float64`
生成正态分布的随机数。

```go
value := random.Normal(100, 15)  // 均值=100, 标准差=15
```

#### `Exponential(lambda float64) float64`
生成指数分布的随机数。

```go
value := random.Exponential(2.0)  // 速率=2.0
```

### 字符串生成

#### 基础字符集

#### `Lowercase(length int) string`
生成随机小写字符串。

```go
str := random.Lowercase(8)  // 返回 "abcdefgh"
```

#### `Uppercase(length int) string`
生成随机大写字符串。

```go
str := random.Uppercase(8)  // 返回 "ABCDEFGH"
```

#### `Digits(length int) string`
生成随机数字字符串。

```go
str := random.Digits(6)  // 返回 "123456"
```

#### `Symbols(length int) string`
生成随机符号字符串。

```go
str := random.Symbols(5)  // 返回 "!@#$%"
```

#### `HexUpper(length int) string`
生成随机大写十六进制字符串。

```go
str := random.HexUpper(8)  // 返回 "12345678"
```

#### 高级字符串生成

#### `AlphanumericSymbols(length int) string`
生成包含字母、数字和符号的随机字符串。

```go
str := random.AlphanumericSymbols(10)  // 返回 "aB3dEfGh!@"
```

#### `StrongPassword(length int) string`
生成强密码（不含易混淆字符）。

```go
str := random.StrongPassword(12)  // 返回 "Kj9mN2pQ7&xY"
```

#### `Readable(length int) string`
生成易读字符串（不含易混淆字符）。

```go
str := random.Readable(10)  // 返回 "abcdefghjk"
```

#### `ShortID(length int) string`
生成适合 URL 的短 ID。

```go
str := random.ShortID(8)  // 返回 "aB3dEfGh"
```

#### `Password(length int, includeSymbols bool) string`
生成密码（可选符号）。

```go
str := random.Password(12, true)   // 带符号: "Kj9#mN2$pQ7&"
str := random.Password(12, false)  // 不带符号: "Kj9mN2pQ7xYz"
```

#### `Username(length int) string`
生成用户名。

```go
str := random.Username(8)  // 返回 "user1234"
```

#### `Email(length int) string`
生成邮箱前缀。

```go
str := random.Email(8)  // 返回 "user1234"
```

#### `Token(length int) string`
生成安全令牌。

```go
str := random.Token(32)  // 返回 "aB3dEfGhJkLmN2pQ7rStUvWxYz123456"
```

#### 颜色生成

#### `ColorHex() string`
生成随机颜色十六进制代码。

```go
color := random.ColorHex()  // 返回 "#FF5733"
```

#### `ColorRGB() string`
生成随机 RGB 颜色字符串。

```go
color := random.ColorRGB()  // 返回 "rgb(255, 87, 51)"
```

#### 网络地址生成

#### `MACAddress() string`
生成随机 MAC 地址。

```go
mac := random.MACAddress()  // 返回 "00:1B:44:11:3A:B7"
```

#### `IPAddress() string`
生成随机 IP 地址。

```go
ip := random.IPAddress()  // 返回 "192.168.1.100"
```

#### 高级字符串生成

#### `WeightedString(chars []WeightedChar, length int) string`
基于字符权重生成字符串。

```go
chars := []random.WeightedChar{
    {Char: 'a', Weight: 1},
    {Char: 'b', Weight: 2},
    {Char: 'c', Weight: 3},
}
str := random.WeightedString(chars, 10)  // 返回 "ccbacccbac"
```

#### `PatternString(pattern string) string`
基于模式生成字符串。

```go
// 模式字符:
// a = 小写字母
// A = 大写字母
// n = 数字
// s = 符号
// x = 任意字母数字
// ? = 任意字符

str := random.PatternString("aAn")  // 返回 "aB3"
```

## 使用示例

### 基础随机生成

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // 生成随机数
    fmt.Printf("随机整数: %d\n", random.Int(100))
    fmt.Printf("随机浮点数: %.4f\n", random.Float64())
    fmt.Printf("随机布尔值: %t\n", random.Bool())
    
    // 生成范围内的值
    fmt.Printf("10-50 随机数: %d\n", random.Between(10, 50))
    fmt.Printf("1.5-3.7 随机浮点数: %.4f\n", random.Float64Between(1.5, 3.7))
}
```

### 采样和选择

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // 随机选择
    colors := []string{"红色", "绿色", "蓝色", "黄色", "紫色"}
    color := random.Choice(colors)
    fmt.Printf("随机颜色: %s\n", color)
    
    // 洗牌
    cards := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    random.Shuffle(cards)
    fmt.Printf("洗牌后的牌: %v\n", cards)
    
    // 无放回采样
    items := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
    sampled := random.Sample(items, 3)
    fmt.Printf("采样元素: %v\n", sampled)
}
```

### 加权选择

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // 加权战利品系统
    lootChoices := []random.WeightedChoice[string]{
        {Weight: 50, Value: "普通物品"},
        {Weight: 30, Value: "稀有物品"},
        {Weight: 15, Value: "史诗物品"},
        {Weight: 4, Value: "传说物品"},
        {Weight: 1, Value: "神话物品"},
    }
    
    // 模拟 10 次战利品掉落
    for i := 0; i < 10; i++ {
        loot := random.SelectWeighted(lootChoices)
        fmt.Printf("掉落 %d: %s\n", i+1, loot)
    }
}
```

### 概率系统

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // 暴击系统
    if random.Percent(15) {  // 15% 暴击概率
        fmt.Println("暴击!")
    } else {
        fmt.Println("普通攻击")
    }
    
    // 技能检定系统
    if random.PercentFloat(0.75) {  // 75% 成功率
        fmt.Println("技能检定通过!")
    } else {
        fmt.Println("技能检定失败!")
    }
}
```

### 统计分析

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // 生成样本数据
    samples := make([]float64, 1000)
    for i := range samples {
        samples[i] = random.Normal(100, 15)  // 均值=100, 标准差=15
    }
    
    // 计算统计量
    sum := 0.0
    for _, sample := range samples {
        sum += sample
    }
    mean := sum / float64(len(samples))
    
    fmt.Printf("样本均值: %.2f\n", mean)
    fmt.Printf("期望均值: 100.00\n")
}
```

### 字符串生成示例

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // 基础字符串生成
    fmt.Printf("小写: %s\n", random.Lowercase(8))
    fmt.Printf("大写: %s\n", random.Uppercase(8))
    fmt.Printf("数字: %s\n", random.Digits(6))
    fmt.Printf("符号: %s\n", random.Symbols(5))
    
    // 高级字符串生成
    fmt.Printf("强密码: %s\n", random.StrongPassword(12))
    fmt.Printf("用户名: %s\n", random.Username(8))
    fmt.Printf("邮箱: %s@example.com\n", random.Email(8))
    fmt.Printf("令牌: %s\n", random.Token(32))
    
    // 颜色生成
    fmt.Printf("颜色十六进制: %s\n", random.ColorHex())
    fmt.Printf("颜色 RGB: %s\n", random.ColorRGB())
    
    // 网络地址
    fmt.Printf("MAC 地址: %s\n", random.MACAddress())
    fmt.Printf("IP 地址: %s\n", random.IPAddress())
}
```

### 模式字符串生成

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // 基于模式的生成
    patterns := []string{
        "aaa",     // 小写字母
        "AAA",     // 大写字母
        "nnn",     // 数字
        "sss",     // 符号
        "xxx",     // 字母数字
        "aAn",     // 混合模式
    }
    
    for _, pattern := range patterns {
        fmt.Printf("模式 %s: %s\n", pattern, random.PatternString(pattern))
    }
}
```

### 加权字符串生成

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // 加权字符选择
    chars := []random.WeightedChar{
        {Char: 'a', Weight: 1},  // 16.7% 概率
        {Char: 'b', Weight: 2},  // 33.3% 概率
        {Char: 'c', Weight: 3},  // 50.0% 概率
    }
    
    // 生成加权字符串
    for i := 0; i < 5; i++ {
        str := random.WeightedString(chars, 10)
        fmt.Printf("加权字符串 %d: %s\n", i+1, str)
    }
}
```

## 性能

该包针对性能进行了优化：

- **全局随机生成器**：单一共享随机数生成器
- **高效算法**：Fisher-Yates 洗牌、水库采样
- **最小分配**：尽可能减少内存分配
- **类型安全**：泛型函数进行编译时类型检查

### 基准测试结果

- **基础操作**：~2-5ns/op，0 分配
- **选择操作**：~10-50ns/op，0-1 分配
- **采样操作**：~100ns-1μs/op，1-3 分配
- **分布生成**：~20-100ns/op，0 分配
- **字符串生成**：~50-200ns/op，1-2 分配
- **大字符串生成**：~1-10μs/op，1-3 分配

## 测试

运行测试：

```bash
go test ./random
```

运行覆盖率测试：

```bash
go test ./random -cover
```

运行示例：

```bash
go test ./random -run Example
```

运行基准测试：

```bash
go test ./random -bench=.
```

运行模糊测试：

```bash
go test ./random -fuzz=.
```

## 许可证

此包是 `goal` 项目的一部分，遵循相同的许可证条款。

## 贡献

欢迎贡献！请随时提交 Pull Request。
