# Stringx 包

一个功能全面且强大的 Go 字符串操作库，提供丰富的字符串处理、验证、转换和构建工具。

## 特性

- **字符串工具**: 全面的字符串操作函数集合
- **字符串构建器**: 具有错误处理和方法链的高级字符串构建
- **字符串验证**: 各种字符串模式验证函数
- **字符串转换**: 大小写转换、格式化和编码工具
- **字符串常量**: 大量常用字符串常量集合
- **模式匹配**: 使用 Aho-Corasick 算法的高效多模式字符串替换
- **Unicode 支持**: 完整的 Unicode 和多字节字符支持
- **性能优化**: 基于 Go 标准库构建，性能优异

## 安装

```bash
go get github.com/go4x/goal/stringx
```

## 快速开始

### 基础字符串操作

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/stringx"
)

func main() {
    // 字符串验证
    fmt.Println(stringx.IsEmail("user@example.com"))        // true
    fmt.Println(stringx.IsNumeric("12345"))                 // true
    fmt.Println(stringx.IsAlpha("Hello"))                  // true
    
    // 字符串转换
    fmt.Println(stringx.ToSnakeCase("HelloWorld"))          // "hello_world"
    fmt.Println(stringx.ToCamelCase("hello_world"))        // "helloWorld"
    fmt.Println(stringx.Reverse("Hello"))                   // "olleH"
    
    // 字符串工具
    fmt.Println(stringx.Cut(10, "This is a long string"))   // "This is a ..."
    fmt.Println(stringx.BlurEmail("john@example.com"))      // "jo****@example.com"
}
```

### 高级字符串构建

```go
// 使用增强的构建器
builder := stringx.NewBuilder()
result := builder.
    WriteString("Hello").
    WriteSpace().
    WriteQuoted("World").
    WriteNewline().
    WriteIndent("  ", 1).
    WriteString("Indented content").
    String()

fmt.Println(result)
// 输出:
// Hello "World"
//   Indented content
```

### 模式替换

```go
// 多模式字符串替换
replacer := stringx.NewReplacer(map[string]string{
    "hello": "hi",
    "world": "universe",
    "goodbye": "bye",
})

text := "hello world, goodbye world"
result := replacer.Replace(text)
fmt.Println(result) // "hi universe, bye universe"
```

## API 参考

### 字符串工具

#### 验证函数
- `IsNumeric(s string) bool` - 检查字符串是否只包含数字字符
- `IsAlpha(s string) bool` - 检查字符串是否只包含字母字符
- `IsAlphaNumeric(s string) bool` - 检查字符串是否只包含字母数字字符
- `IsEmail(s string) bool` - 基本邮箱格式验证
- `IsEmpty(str string) bool` - 检查字符串是否为空或只包含空白字符
- `IsSpace(s string) bool` - 检查字符串是否只包含空格
- `HasLen(s string) bool` - 检查字符串是否有有意义的内容

#### 转换函数
- `ToSnakeCase(s string) string` - 转换为 snake_case
- `ToKebabCase(s string) string` - 转换为 kebab-case
- `ToPascalCase(s string) string` - 转换为 PascalCase
- `ToCamelCase(s string) string` - 转换为 camelCase
- `ToTitle(s string) string` - 转换为标题格式
- `Reverse(s string) string` - 反转字符串
- `RemoveDuplicates(s string) string` - 移除连续重复字符

#### 字符串操作
- `Cut(max int, s string) string` - 截断字符串并添加省略号
- `Trim(s, cut string) string` - 从两端修剪字符
- `TrimLeft(s, cut string) string` - 从左端修剪字符
- `TrimRight(s, cut string) string` - 从右端修剪字符
- `TrimSpace(s string) string` - 修剪空白字符
- `RemSpace(s string) string` - 移除所有空格
- `Replace(s, old, new string, n int) string` - 替换出现次数
- `ReplaceAll(s, old, new string) string` - 替换所有出现次数

#### 填充函数
- `PadLeft(s string, length int, padChar rune) string` - 左填充
- `PadRight(s string, length int, padChar rune) string` - 右填充
- `PadCenter(s string, length int, padChar rune) string` - 居中填充

#### 字符串分析
- `CountWords(s string) int` - 统计字符串中的单词数
- `CountLines(s string) int` - 统计字符串中的行数
- `CountOccurrences(s, substr string) int` - 统计子字符串出现次数
- `FindAll(s, substr string) []int` - 查找所有子字符串位置
- `ContainsAny(s string, substrings ...string) bool` - 检查是否包含任意子字符串
- `ContainsAll(s string, substrings ...string) bool` - 检查是否包含所有子字符串

#### 字符串分割和连接
- `SplitAndTrim(s, sep string) []string` - 分割并修剪空白字符
- `JoinNonEmpty(sep string, strs ...string) string` - 连接非空字符串
- `Chunk(s string, size int) []string` - 分割成块
- `Wrap(s string, width int) []string` - 按宽度换行

#### 隐私函数
- `BlurEmail(email string) string` - 模糊化邮箱地址
- `Blur(str string, start, end int, sep string, num int) string` - 模糊化字符串内容

### 字符串构建器

`Builder` 类型提供具有错误处理和方法链的高级字符串构建功能。

#### 基础操作
- `NewBuilder() *Builder` - 创建新构建器
- `WriteString(s string) *Builder` - 追加字符串
- `WriteRune(r rune) *Builder` - 追加 rune
- `WriteByte(c byte) *Builder` - 追加字节
- `Write(p []byte) *Builder` - 追加字节切片

#### 格式化操作
- `Writef(format string, args ...interface{}) *Builder` - 格式化写入
- `WriteLine(s string) *Builder` - 写入字符串并换行
- `WriteLinef(format string, args ...interface{}) *Builder` - 格式化写入并换行

#### 条件操作
- `WriteIf(condition bool, s string) *Builder` - 条件写入
- `WriteIfElse(condition bool, ifTrue, ifFalse string) *Builder` - 条件选择

#### 重复和连接
- `WriteRepeat(s string, n int) *Builder` - 重复字符串
- `WriteJoin(sep string, strs ...string) *Builder` - 连接字符串

#### 空白字符操作
- `WriteSpace() *Builder` - 写入空格
- `WriteTab() *Builder` - 写入制表符
- `WriteNewline() *Builder` - 写入换行符
- `WriteIndent(indent string, n int) *Builder` - 写入缩进

#### 包装操作
- `WriteWrap(prefix, content, suffix string) *Builder` - 包装内容
- `WriteQuoted(content string) *Builder` - 用双引号包装
- `WriteSingleQuoted(content string) *Builder` - 用单引号包装
- `WriteBacktickQuoted(content string) *Builder` - 用反引号包装
- `WriteBrackets(content string) *Builder` - 用方括号包装
- `WriteParentheses(content string) *Builder` - 用圆括号包装
- `WriteBraces(content string) *Builder` - 用花括号包装

#### 状态操作
- `Len() int` - 获取当前长度
- `Cap() int` - 获取当前容量
- `Reset() *Builder` - 重置构建器
- `Grow(n int) *Builder` - 预分配容量
- `Error() error` - 获取任何错误
- `String() string` - 获取最终字符串

### 字符串常量

包提供了按类别组织的广泛字符串常量：

#### 标点符号和符号
```go
stringx.Exclamation     // "!"
stringx.AtSign          // "@"
stringx.HashTag         // "#"
stringx.DollarSign      // "$"
stringx.PercentSign     // "%"
stringx.Caret           // "^"
stringx.AmpersandSign   // "&"
stringx.StarSign        // "*"
stringx.PlusSign        // "+"
stringx.MinusSign       // "-"
stringx.EqualsSign      // "="
stringx.UnderscoreSign  // "_"
stringx.PipeSign        // "|"
stringx.BackslashSign   // "\\"
stringx.ForwardSlash    // "/"
stringx.ColonSign       // ":"
stringx.SemicolonSign   // ";"
stringx.CommaSign       // ","
stringx.DotSign         // "."
stringx.QuestionSign    // "?"
```

#### 括号和引号
```go
stringx.LeftParen       // "("
stringx.RightParen      // ")"
stringx.LeftBracket     // "["
stringx.RightBracket    // "]"
stringx.LeftBrace       // "{"
stringx.RightBrace      // "}"
stringx.LeftAngle       // "<"
stringx.RightAngle      // ">"
stringx.DoubleQuote     // "\""
stringx.SingleQuote     // "'"
stringx.BacktickQuote   // "`"
```

#### 空白字符
```go
stringx.SpaceChar       // " "
stringx.TabChar         // "\t"
stringx.NewlineChar     // "\n"
stringx.CarriageReturn  // "\r"
stringx.FormFeed        // "\f"
stringx.VerticalTab     // "\v"
```

#### 常用分隔符
```go
stringx.CommaSpace      // ", "
stringx.SemicolonSpace  // "; "
stringx.ColonSpace      // ": "
stringx.PipeSpace       // " | "
stringx.SlashSpace      // " / "
stringx.BackslashSpace  // " \\ "
```

#### 布尔值
```go
stringx.BooleanTrue     // "true"
stringx.BooleanFalse    // "false"
stringx.BooleanYes      // "yes"
stringx.BooleanNo       // "no"
stringx.BooleanOn       // "on"
stringx.BooleanOff      // "off"
stringx.BooleanEnabled  // "enabled"
stringx.BooleanDisabled // "disabled"
```

#### 文件扩展名
```go
stringx.ExtJSON         // ".json"
stringx.ExtXML          // ".xml"
stringx.ExtHTML         // ".html"
stringx.ExtCSS          // ".css"
stringx.ExtJS           // ".js"
stringx.ExtGo           // ".go"
stringx.ExtTxt          // ".txt"
stringx.ExtLog          // ".log"
stringx.ExtYAML         // ".yaml"
stringx.ExtYML          // ".yml"
stringx.ExtTOML         // ".toml"
stringx.ExtINI          // ".ini"
stringx.ExtConf         // ".conf"
stringx.ExtConfig       // ".config"
```

#### 协议
```go
stringx.ProtocolHTTP            // "http"
stringx.ProtocolHTTPS           // "https"
stringx.ProtocolFTP             // "ftp"
stringx.ProtocolFTPS            // "ftps"
stringx.ProtocolSFTP            // "sftp"
stringx.ProtocolSSH             // "ssh"
stringx.ProtocolTCP             // "tcp"
stringx.ProtocolUDP             // "udp"
stringx.ProtocolWebSocket       // "ws"
stringx.ProtocolWebSocketSecure // "wss"
```

#### 编码类型
```go
stringx.EncodingUTF8   // "utf-8"
stringx.EncodingUTF16  // "utf-16"
stringx.EncodingASCII  // "ascii"
stringx.EncodingBase64 // "base64"
stringx.EncodingHex    // "hex"
stringx.EncodingURL    // "url"
```

#### 时区
```go
stringx.TimezoneUTC // "UTC"
stringx.TimezoneGMT // "GMT"
stringx.TimezoneEST // "EST"
stringx.TimezonePST // "PST"
stringx.TimezoneCST // "CST"
stringx.TimezoneMST // "MST"
```

#### 单位
```go
stringx.UnitBytes // "bytes"
stringx.UnitKB    // "KB"
stringx.UnitMB    // "MB"
stringx.UnitGB    // "GB"
stringx.UnitTB    // "TB"
stringx.UnitPB    // "PB"
```

## 示例

### JSON 生成
```go
builder := stringx.NewBuilder()
json := builder.
    WriteBraces("").
    WriteNewline().
    WriteIndent("  ", 1).
    WriteQuoted("name").
    WriteString(stringx.Colon + stringx.Space).
    WriteQuoted("John").
    WriteString(stringx.Comma).
    WriteNewline().
    WriteIndent("  ", 1).
    WriteQuoted("age").
    WriteString(stringx.Colon + stringx.Space).
    WriteString("25").
    WriteNewline().
    WriteString("}").
    String()

fmt.Println(json)
// 输出:
// {
//   "name": "John",
//   "age": 25
// }
```

### URL 构建
```go
url := stringx.ProtocolHTTPS + 
    stringx.Colon + 
    stringx.ForwardSlash + 
    stringx.ForwardSlash + 
    "api.example.com" + 
    stringx.Slash + 
    "v1" + 
    stringx.Slash + 
    "users"
```

### 文件路径构建
```go
configPath := "config" + stringx.Slash + "app" + stringx.ExtJSON
logPath := "logs" + stringx.Slash + "app" + stringx.ExtLog
```

### 字符串验证管道
```go
func validateUserInput(input string) error {
    if stringx.IsEmpty(input) {
        return errors.New("输入不能为空")
    }
    
    if !stringx.IsAlphaNumeric(input) {
        return errors.New("输入必须为字母数字")
    }
    
    if stringx.CountWords(input) > 10 {
        return errors.New("输入过长")
    }
    
    return nil
}
```

### 文本处理
```go
func processText(text string) string {
    // 移除重复并标准化
    processed := stringx.RemoveDuplicates(text)
    
    // 转换为标题格式
    processed = stringx.ToTitle(processed)
    
    // 长行换行
    lines := stringx.Wrap(processed, 80)
    
    // 使用适当格式连接
    return stringx.JoinNonEmpty(stringx.NewlineChar, lines...)
}
```

## 性能考虑

- **内存效率**: 包设计为最小内存分配
- **Unicode 支持**: 完整支持多字节字符和 Unicode
- **构建器模式**: 通过预分配实现高效的字符串构建
- **算法优化**: 使用 Aho-Corasick 等高效算法进行模式匹配
- **零拷贝操作**: 许多操作避免不必要的字符串复制

## 最佳实践

1. **使用常量**: 优先使用字符串常量而不是硬编码字符串
2. **复杂字符串使用构建器**: 使用构建器进行复杂的字符串构建
3. **早期验证**: 在管道早期使用验证函数
4. **处理错误**: 使用构建器时始终检查错误
5. **Unicode 意识**: 在字符串操作中注意多字节字符

## 许可证

此包是 goal 项目的一部分。有关详细信息，请参阅主项目许可证。

## 贡献

欢迎贡献！请随时提交 Pull Request。

## 更新日志

### v1.0.0
- 具有全面字符串工具的初始版本
- 具有方法链的高级构建器
- 广泛的字符串常量集合
- 使用 Aho-Corasick 算法的模式匹配
- 完整的 Unicode 支持
- 全面的测试覆盖
