# IOx - 文件和目录操作包

IOx 是一个功能全面的 Go 包，提供便捷的文件和目录操作工具。它既提供基础功能，也包含高级特性，如支持复杂过滤器组合的灵活文件遍历。

## 特性

### 🔧 核心操作
- **文件操作**: 检查存在性、获取文件信息、处理文件类型
- **目录操作**: 创建、删除、检查空目录、递归遍历
- **路径工具**: 获取可执行文件路径、当前路径、项目根目录
- **文本文件处理**: 带缓冲的读写操作，提供便捷方法

### 🎯 高级特性
- **灵活的文件遍历器**: 支持强大过滤功能的递归目录遍历
- **复杂过滤器组合**: 支持多过滤器组的 AND/OR 逻辑
- **性能优化**: 高效的算法，提供基准测试
- **跨平台**: 支持 Windows、macOS 和 Linux

### 📊 质量保证
- **高测试覆盖率**: 61.4% 语句覆盖率，包含全面的测试套件
- **性能基准测试**: 所有主要功能都有性能基准测试
- **完整文档**: 完整的 API 文档和使用示例
- **错误处理**: 强大的错误处理，提供清晰的错误信息

## 安装

```bash
go get github.com/go4x/goal/iox
```

## 快速开始

### 基础文件操作

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/iox"
)

func main() {
    // 检查文件是否存在
    if iox.Exists("/path/to/file.txt") {
        fmt.Println("文件存在!")
    }

    // 检查目录是否存在
    if iox.IsDir("/path/to/directory") {
        fmt.Println("这是一个目录!")
    }

    // 检查是否为普通文件
    if iox.IsRegularFile("/path/to/file.txt") {
        fmt.Println("这是一个普通文件!")
    }
}
```

### 目录操作

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/iox"
)

func main() {
    // 创建目录
    err := iox.Dir.Create("/path/to/new/directory")
    if err != nil {
        fmt.Printf("创建目录错误: %v\n", err)
        return
    }

    // 检查目录是否为空
    isEmpty, err := iox.Dir.IsEmpty("/path/to/directory")
    if err != nil {
        fmt.Printf("检查目录错误: %v\n", err)
        return
    }
    
    if isEmpty {
        fmt.Println("目录为空")
    }

    // 递归遍历目录
    files, err := iox.Dir.Walk("/path/to/directory")
    if err != nil {
        fmt.Printf("遍历目录错误: %v\n", err)
        return
    }
    
    for _, file := range files {
        fmt.Println("找到文件:", file)
    }
}
```

### 文本文件操作

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/iox"
)

func main() {
    // 创建新的文本文件
    tf, err := iox.NewTxtFile("/path/to/file.txt")
    if err != nil {
        fmt.Printf("创建文本文件错误: %v\n", err)
        return
    }
    defer tf.Close()

    // 写入行到文件
    _, err = tf.WriteLine("Hello, World!")
    if err != nil {
        fmt.Printf("写入行错误: %v\n", err)
        return
    }
    
    _, err = tf.WriteLine("这是第二行")
    if err != nil {
        fmt.Printf("写入行错误: %v\n", err)
        return
    }

    // 刷新缓冲区确保数据写入
    err = tf.Flush()
    if err != nil {
        fmt.Printf("刷新错误: %v\n", err)
        return
    }

    // 从文件读取所有行
    lines, err := tf.ReadAll()
    if err != nil {
        fmt.Printf("读取文件错误: %v\n", err)
        return
    }
    
    for _, line := range lines {
        fmt.Println("行:", line)
    }
}
```

### 高级文件遍历和过滤

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/iox"
)

func main() {
    // 简单过滤 - 查找所有 .go 文件
    goFiles, err := iox.WalkDir("/path/to/project", iox.FilterByExtension(".go"))
    if err != nil {
        fmt.Printf("遍历目录错误: %v\n", err)
        return
    }
    
    fmt.Printf("找到 %d 个 Go 文件\n", len(goFiles))
    for _, file := range goFiles {
        fmt.Println("Go 文件:", file)
    }

    // 复杂过滤，使用多个过滤器组
    goGroup := iox.NewFilterGroup(iox.FilterAnd,
        iox.FilterByExtension(".go"),
        iox.FilterHidden, // 排除隐藏文件
    )
    
    txtGroup := iox.NewFilterGroup(iox.FilterAnd,
        iox.FilterByExtension(".txt"),
        iox.FilterByName("README"),
    )
    
    // 匹配任一组的文件都会被包含（组间使用 OR 逻辑）
    files, err := iox.WalkDirWithFilters("/path/to/project", goGroup, txtGroup)
    if err != nil {
        fmt.Printf("遍历目录错误: %v\n", err)
        return
    }
    
    fmt.Printf("找到 %d 个匹配文件\n", len(files))
    for _, file := range files {
        fmt.Println("匹配文件:", file)
    }
}
```

### 路径工具

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/iox"
)

func main() {
    // 获取可执行文件目录
    execPath := iox.Path.ExecPath()
    fmt.Printf("可执行文件目录: %s\n", execPath)

    // 获取当前源文件目录
    currentPath := iox.Path.CurrentPath()
    fmt.Printf("当前目录: %s\n", currentPath)

    // 获取 Go 项目根目录
    projectPath := iox.Path.ProjectPath()
    fmt.Printf("项目根目录: %s\n", projectPath)

    // 检查路径是否为文件
    if iox.Path.IsFile("/path/to/file.txt") {
        fmt.Println("这是一个文件!")
    }

    // 检查路径是否为目录
    if iox.Path.IsDir("/path/to/directory") {
        fmt.Println("这是一个目录!")
    }
}
```

## API 参考

### 核心函数

#### 文件和目录存在性检查
```go
// 检查文件或目录是否存在
func Exists(file string) bool

// 检查路径是否存在并返回是否为目录
func PathExists(path string) (bool, error)

// 检查路径是否为目录
func IsDir(path string) bool

// 检查路径是否为普通文件
func IsRegularFile(path string) bool
```

#### 全局实例
```go
// 文件操作
var File *files

// 目录操作  
var Dir *dirs

// 路径工具
var Path *paths
```

### 文件操作

#### 文件实例方法
```go
// 检查文件是否存在（非目录）
func (f *files) Exists(file string) bool

// 获取文件信息
func (f *files) Info(file string) (os.FileInfo, error)
```

### 目录操作

#### 目录实例方法
```go
// 检查目录是否存在
func (d *dirs) Exists(dir string) (bool, error)

// 如果需要则添加路径分隔符
func (d *dirs) AppendSeparator(dir string) string

// 创建目录和父目录
func (d *dirs) Create(dir string) error

// 仅在目录不存在时创建
func (d *dirs) CreateIfNotExists(dir string) error

// 检查目录是否为空
func (d *dirs) IsEmpty(dir string) (bool, error)

// 递归删除目录
func (d *dirs) Delete(dir string) error

// 仅在目录存在时删除
func (d *dirs) DeleteIfExists(dir string) error

// 仅在目录为空时删除
func (d *dirs) DeleteIfEmpty(dir string) error

// 递归遍历目录（返回所有文件）
func (d *dirs) Walk(dir string) ([]string, error)
```

### 路径工具

#### 路径实例方法
```go
// 获取可执行文件目录路径
func (ps *paths) ExecPath() string

// 获取当前源文件目录
func (ps *paths) CurrentPath() string

// 检查路径是否存在
func (ps *paths) PathExists(path string) bool

// 检查路径是否为文件
func (ps *paths) IsFile(path string) bool

// 检查路径是否为目录
func (ps *paths) IsDir(path string) bool

// 获取 Go 项目根目录路径
func (ps *paths) ProjectPath() string
```

### 文本文件操作

#### TxtFile 方法
```go
// 创建新的文本文件
func NewTxtFile(f string) (*TxtFile, error)

// 写入行到文件（缓冲）
func (tf *TxtFile) WriteLine(s string) (*TxtFile, error)

// 刷新缓冲数据到磁盘
func (tf *TxtFile) Flush() error

// 关闭文件并清理
func (tf *TxtFile) Close() error

// 从文件读取所有行
func (tf *TxtFile) ReadAll() ([]string, error)
```

### 文件遍历和过滤

#### 遍历函数
```go
// 使用简单过滤器遍历目录
func WalkDir(dir string, filters ...WalkFilter) ([]string, error)

// 使用复杂过滤器组遍历目录
func WalkDirWithFilters(dir string, filterGroups ...FilterGroup) ([]string, error)
```

#### 过滤器函数
```go
// 按文件扩展名过滤
func FilterByExtension(extensions ...string) WalkFilter

// 按名称模式过滤（包含）
func FilterByName(pattern string) WalkFilter

// 按大小范围过滤
func FilterBySize(minSize, maxSize int64) WalkFilter

// 按路径模式过滤（正则表达式）
func FilterByPathPattern(pattern string) WalkFilter

// 仅包含目录
func FilterDirectoriesOnly(entry os.DirEntry, path string) bool

// 仅包含文件
func FilterFilesOnly(entry os.DirEntry, path string) bool

// 排除隐藏文件
func FilterHidden(entry os.DirEntry, path string) bool
```

#### 过滤器组操作
```go
// 创建新的过滤器组
func NewFilterGroup(combiner FilterCombiner, filters ...WalkFilter) FilterGroup

// 将过滤器组应用于目录条目
func (fg FilterGroup) Apply(entry os.DirEntry, path string) bool
```

#### 过滤器组合器
```go
// 逻辑 AND 组合
const FilterAnd FilterCombiner

// 逻辑 OR 组合  
const FilterOr FilterCombiner
```

## 高级用法

### 复杂过滤器组合

文件遍历器支持多种过滤器组的复杂过滤：

```go
// 创建具有不同逻辑的过滤器组
goFiles := iox.NewFilterGroup(iox.FilterAnd,
    iox.FilterByExtension(".go"),
    iox.FilterHidden, // 排除隐藏文件
)

largeFiles := iox.NewFilterGroup(iox.FilterAnd,
    iox.FilterBySize(1024, 1024*1024), // 1KB 到 1MB
    iox.FilterFilesOnly,
)

// 匹配任一组的文件都会被包含
files, err := iox.WalkDirWithFilters("/project", goFiles, largeFiles)
```

### 性能考虑

该包针对性能进行了优化：

```go
// 基准测试结果（典型）：
// Exists: ~1,175 ns/op
// IsDir: ~1,177 ns/op  
// WalkDir: ~25,000 ns/op (100 个文件)
// WriteLine: ~71 ns/op
// ReadAll: ~25,382 ns/op (1000 行)
```

### 错误处理

所有函数都返回适当的错误：

```go
// 始终检查错误
files, err := iox.WalkDir("/path")
if err != nil {
    log.Fatalf("WalkDir 失败: %v", err)
}

// 处理特定错误类型
if iox.Exists("/path") {
    // 路径存在，可以安全继续
} else {
    // 路径不存在，相应处理
}
```

## 测试

该包包含全面的测试：

```bash
# 运行所有测试
go test ./iox

# 运行带覆盖率的测试
go test -cover ./iox

# 运行基准测试
go test -bench=. ./iox

# 运行特定测试
go test -run TestWalkDir ./iox
```

### 测试覆盖率

- **语句覆盖率**: 61.4%
- **测试类型**: 单元测试、集成测试、边界情况测试、性能基准测试
- **测试文件**: 7 个测试文件，包含 25+ 个测试函数

## 性能基准测试

关键性能指标：

| 函数 | 性能 | 说明 |
|------|------|------|
| `Exists` | ~1,175 ns/op | 基础文件存在性检查 |
| `IsDir` | ~1,177 ns/op | 目录类型检查 |
| `WalkDir` | ~25,000 ns/op | 100 个文件遍历 |
| `WriteLine` | ~71 ns/op | 缓冲文本写入 |
| `ReadAll` | ~25,382 ns/op | 1000 行读取 |
| `ExecPath` | ~66 ns/op | 可执行文件路径 |
| `ProjectPath` | ~6.8 ms/op | 包含 `go env GOMOD` |

## 贡献

1. Fork 仓库
2. 创建功能分支
3. 为新功能添加测试
4. 确保所有测试通过
5. 提交拉取请求

### 开发指南

- 遵循 Go 约定
- 添加全面的测试
- 更新文档
- 为性能关键代码运行基准测试
- 保持向后兼容性

## 许可证

本项目采用 Apache License 2.0 - 详见 [LICENSE](LICENSE) 文件。

## 更新日志

### v1.0.0
- 初始发布
- 核心文件和目录操作
- 文本文件处理
- 支持过滤器的高级文件遍历
- 全面的测试套件
- 性能基准测试
- 完整文档