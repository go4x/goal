# TimeX - Go 语言综合时间工具包

TimeX 是一个综合性的时间工具包，扩展了 Go 标准 `time` 包的功能，提供时间格式化、工作日计算、时区处理以及各种常用时间操作的实用函数。

## 特性

- **时间格式化**: 多种预定义格式和自定义格式化
- **工作日计算**: 计算工作日和业务操作
- **相对时间格式化**: 人类可读的时间差（如"2小时前"）
- **时区处理**: 夏令时支持和时区转换
- **批量操作**: 时间切片排序、过滤和分组
- **性能工具**: 执行时间测量和优化

## 安装

```bash
go get github.com/go4x/goal/timex
```

## 快速开始

```go
package main

import (
    "fmt"
    "time"
    "github.com/go4x/goal/timex"
)

func main() {
    // 时间格式化
    t := time.Now()
    formatted := timex.Format(t, timex.YmdDash)
    fmt.Println(formatted)
    
    // 工作日计算
    start := time.Date(2023, 1, 9, 0, 0, 0, 0, time.UTC) // 周一
    end := time.Date(2023, 1, 13, 0, 0, 0, 0, time.UTC)   // 周五
    businessDays := timex.BusinessDaysBetween(start, end)
    fmt.Printf("工作日数: %d\n", businessDays)
    
    // 相对时间格式化
    relative := timex.FormatRelative(time.Now().Add(-2 * time.Hour))
    fmt.Printf("时间差: %s\n", relative) // "2小时前"
}
```

## API 参考

### 时间解析

#### `From(s string, layout Layout) (time.Time, error)`
使用特定格式解析时间字符串。

```go
// 使用特定格式解析
t1, _ := timex.From("2023-01-01 12:30:45", timex.YmdhmsDash)
t2, _ := timex.From("2023/01/01 12:30:45", timex.YmdhmsSlash)
t3, _ := timex.From("2023-01-01", timex.YmdDash)
```

#### `MustFrom(s string, layout Layout) time.Time`
使用特定格式解析时间字符串，解析失败时抛出异常。

### 时间格式化

#### `Format(date time.Time, layout Layout) string`
使用指定格式格式化时间。

#### `FormatRelative(t time.Time) string`
将时间格式化为相对文本（如"2小时前"、"3天后"）。

#### `FormatDuration(d time.Duration) string`
将持续时间格式化为人类可读文本（如"2小时30分钟"）。

#### `FormatChinese(t time.Time) string`
将时间格式化为中文格式（如"2023年01月01日 12时30分45秒"）。

#### `FormatISO(t time.Time) string`
将时间格式化为 ISO 标准格式（RFC3339）。

### 时间计算

#### `DiffDay(src, dst time.Time) int`
计算两个时间之间的绝对天数差。

#### `BusinessDaysBetween(start, end time.Time) int`
计算两个时间之间的工作日数。

#### `AddBusinessDays(t time.Time, days int) time.Time`
向时间添加指定数量的工作日。

#### `NextBusinessDay(t time.Time) time.Time`
获取给定时间后的下一个工作日。

#### `PrevBusinessDay(t time.Time) time.Time`
获取给定时间前的上一个工作日。

### 时间查询

#### `IsToday(t time.Time) bool`
检查时间是否为今天。

#### `IsYesterday(t time.Time) bool`
检查时间是否为昨天。

#### `IsTomorrow(t time.Time) bool`
检查时间是否为明天。

#### `IsWeekend(t time.Time) bool`
检查时间是否为周末（周六或周日）。

#### `IsBusinessDay(t time.Time) bool`
检查时间是否为工作日（周一到周五）。

#### `IsInRange(t, start, end time.Time) bool`
检查时间是否在指定范围内。

### 周期边界

#### `StartTime(t time.Time) time.Time`
获取一天的开始时间（00:00:00.000000000）。

#### `EndTime(t time.Time) time.Time`
获取一天的结束时间（23:59:59.999999999）。

#### `StartOfWeek(t time.Time) time.Time`
获取一周的开始时间（周一）。

#### `EndOfWeek(t time.Time) time.Time`
获取一周的结束时间（周日）。

#### `StartOfMonth(t time.Time) time.Time`
获取一个月的开始时间（第一天）。

#### `EndOfMonth(t time.Time) time.Time`
获取一个月的结束时间（最后一天）。

#### `StartOfYear(t time.Time) time.Time`
获取一年的开始时间（1月1日）。

#### `EndOfYear(t time.Time) time.Time`
获取一年的结束时间（12月31日）。

#### `QuarterOf(t time.Time) int`
获取给定时间的季度（1-4）。

#### `StartOfQuarter(t time.Time) time.Time`
获取季度的开始时间。

#### `EndOfQuarter(t time.Time) time.Time`
获取季度的结束时间。

### 日期工具

#### `Age(birthDate, currentDate time.Time) int`
计算年龄（年数）。

#### `IsLeapYear(year int) bool`
检查年份是否为闰年。

#### `DaysInMonth(year int, month time.Month) int`
获取指定年份和月份的天数。

#### `SameDate(src, dst time.Time) bool`
检查两个时间是否表示同一个日期。

### 时区操作

#### `ConvertTimezone(t time.Time, location *time.Location) time.Time`
将时间转换为不同时区。

#### `GetTimezoneOffset(t time.Time) time.Duration`
获取时区偏移量。

#### `IsDST(t time.Time) bool`
检查时间是否处于夏令时。

#### `GetTimezoneInfo(t time.Time) map[string]interface{}`
获取综合时区信息。

### 批量操作

#### `SortTimes(times []time.Time) []time.Time`
按升序对时间切片进行排序。

#### `FindClosest(times []time.Time, target time.Time) time.Time`
在切片中找到最接近目标时间的时间。

#### `GroupByDay(times []time.Time) map[string][]time.Time`
按天对时间进行分组。

#### `FilterByRange(times []time.Time, start, end time.Time) []time.Time`
过滤出指定范围内的时间。

### 工具函数

#### `RoundToNearest(t time.Time, unit time.Duration) time.Time`
将时间舍入到最近的单位。

#### `TruncateTo(t time.Time, unit time.Duration) time.Time`
将时间截断到指定单位。

#### `MeasureExecution(fn func()) time.Duration`
测量函数的执行时间。

## 预定义格式

TimeX 提供了多个预定义格式常量：

```go
timex.Ymdhms      // "20060102150405"
timex.Ymd         // "20060102"
timex.Hms         // "15:04:05"
timex.YmdhmsDash  // "2006-01-02 15:04:05"
timex.YmdDash     // "2006-01-02"
timex.YmdhmsSlash // "2006/01/02 15:04:05"
timex.YmdSlash    // "2006/01/02"
timex.YmdhmsZh    // "2006年01月02日 15时04分05秒"
timex.YmdZh       // "2006年01月02日"
timex.HmsZh       // "15时04分05秒"
```

## 性能

TimeX 针对性能进行了优化，许多函数实现了零分配：

```
BenchmarkParseSmart-12             	19308361	        62.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkDiffDay-12                	28509064	        42.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsBusinessDay-12          	398281634	         3.024 ns/op	       0 B/op	       0 allocs/op
BenchmarkConvertTimezone-12        	1000000000	         0.3156 ns/op	       0 B/op	       0 allocs/op
```

## 测试

该包包含全面的测试，代码覆盖率达到 86.6%：

```bash
go test -cover ./timex
```

运行基准测试：

```bash
go test -bench=. -benchmem ./timex
```

## 示例

查看 [example_test.go](example_test.go) 文件获取全面的使用示例。

## 许可证

该项目是 Goal 库的一部分，遵循相同的许可证条款。

## 贡献

欢迎贡献！请确保所有测试通过并保持高代码覆盖率。

## 使用场景

### 业务应用
- **工作日计算**: 计算项目工期、休假天数
- **时间格式化**: 统一的时间显示格式
- **时区转换**: 多地区业务的时间处理

### 数据分析
- **时间分组**: 按天、周、月对数据进行分组
- **时间过滤**: 筛选特定时间范围的数据
- **时间排序**: 对时间序列数据进行排序

### 性能监控
- **执行时间测量**: 监控函数执行时间
- **性能分析**: 识别性能瓶颈
- **优化验证**: 验证性能优化效果

## 最佳实践

1. **使用预定义格式**: 优先使用预定义的时间格式常量
2. **时区处理**: 在跨时区应用中始终明确指定时区
3. **工作日计算**: 使用专门的工作日函数而不是简单的天数计算
4. **性能考虑**: 在循环中避免重复的时间格式化操作
5. **错误处理**: 始终检查时间解析函数的错误返回值

## 更新日志

### v1.0.0
- 初始版本发布
- 包含全面的时间工具函数
- 支持工作日计算和时区处理
- 提供中文时间格式化
- 包含性能优化的批量操作
- 全面的测试覆盖
