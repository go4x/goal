# ciphers

一个功能全面的 Go 加密包，提供加密、解密、哈希和数据压缩功能。

## 功能特性

- **AES 加密/解密**：支持 ECB 和 CBC 模式，包含安全警告
- **哈希函数**：SHA256、MD5 和 BCrypt，满足各种哈希需求
- **数据压缩**：Base36、Base62 和 Base64 编码，用于紧凑的数据表示
- **Base64 变体**：标准、URL 安全和 Base64URL 编码，适用于不同用例
- **安全优先设计**：明确警告不安全模式并提供最佳实践

## 安装

```bash
go get github.com/go4x/goal/ciphers
```

## 快速开始

```go
package main

import (
    "fmt"
    "log"
    "github.com/go4x/goal/ciphers"
    "github.com/go4x/goal/ciphers/hash"
    "github.com/go4x/goal/ciphers/base"
)

func main() {
    // AES 加密（安全的 CBC 模式）
    data := []byte("Hello World")
    key := []byte("my-32-byte-long-key-123456789012")
    iv := []byte("my-16-byte-iv-12")
    
    encrypted, err := ciphers.AES.Encrypt(data, key, ciphers.CBC, iv)
    if err != nil {
        log.Fatal(err)
    }
    
    decrypted, err := ciphers.AES.Decrypt(encrypted, key, ciphers.CBC, iv)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(string(decrypted)) // 输出: Hello World
    
    // 哈希函数
    hash := hash.SHA256("hello world")
    fmt.Println(hash)
    
    // 数据压缩
    compressed := base.Base62(12345)
    fmt.Println(compressed) // 输出: 3d7
}
```

## AES 加密/解密

### 安全加密（推荐）

```go
import "github.com/go4x/goal/ciphers"

// 使用 CBC 模式和随机 IV 进行安全加密
data := []byte("敏感数据")
key := []byte("your-32-byte-key-here-123456789012") // AES-256 需要 32 字节
iv := []byte("random-16-byte-iv") // AES 需要 16 字节

// 加密
encrypted, err := ciphers.AES.Encrypt(data, key, ciphers.CBC, iv)
if err != nil {
    log.Fatal(err)
}

// 解密
decrypted, err := ciphers.AES.Decrypt(encrypted, key, ciphers.CBC, iv)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(decrypted)) // 输出: 敏感数据
```

### 安全警告

⚠️ **重要安全注意事项：**

- **ECB 模式不安全**，不应在生产环境中使用
- 始终使用 **CBC 模式和随机 IV** 进行安全加密
- 确保密钥是密码学随机的且大小合适
- 永远不要为同一密钥重复使用 IV

```go
// ❌ 不要使用 - 不安全的 ECB 模式
encrypted, err := ciphers.AES.Encrypt(data, key, ciphers.ECB, nil)

// ✅ 使用这个 - 安全的 CBC 模式
encrypted, err := ciphers.AES.Encrypt(data, key, ciphers.CBC, randomIV)
```

## 哈希函数

### SHA256

```go
import "github.com/go4x/goal/ciphers/hash"

// SHA256 哈希
hash := hash.SHA256("hello world")
fmt.Println(hash) // 输出: b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
```

### MD5

```go
import "github.com/go4x/goal/ciphers/hash"

// MD5 哈希（不推荐用于安全关键应用）
hash := hash.MD5("hello world")
fmt.Println(hash) // 输出: 5d41402abc4b2a76b9719d911017c592
```

### BCrypt（推荐用于密码）

```go
import "github.com/go4x/goal/ciphers/hash"

// 哈希密码
password := "mysecretpassword"
hashedPassword, err := hash.BCrypt(password)
if err != nil {
    log.Fatal(err)
}

// 验证密码
isValid := hash.BCryptVerify(password, hashedPassword)
fmt.Println(isValid) // 输出: true
```

## 数据压缩

### Base36 编码

```go
import "github.com/go4x/goal/ciphers/base"

// 将数字编码为 Base36
compressed := base.Base36(12345)
fmt.Println(compressed) // 输出: 9IX

// 将 Base36 解码为数字
number, ok := base.Base36Decode("9IX")
if ok {
    fmt.Println(number) // 输出: 12345
}
```

### Base62 编码

```go
import "github.com/go4x/goal/ciphers/base"

// 将数字编码为 Base62
compressed := base.Base62(12345)
fmt.Println(compressed) // 输出: 3d7

// 将 Base62 解码为数字
number, ok := base.Base62Decode("3d7")
if ok {
    fmt.Println(number) // 输出: 12345
}
```

## Base64 变体

### 标准 Base64

```go
import "github.com/go4x/goal/ciphers/base64x"

// 标准 base64 编码
data := []byte("hello world")
encoded := base64x.StdEncoding.Encode(data)
fmt.Println(encoded) // 输出: aGVsbG8gd29ybGQ=

// 解码
decoded, err := base64x.StdEncoding.Decode(encoded)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(decoded)) // 输出: hello world
```

### URL 安全 Base64

```go
import "github.com/go4x/goal/ciphers/base64x"

// URL 安全 base64 编码
data := []byte("hello world")
encoded := base64x.URLEncoding.Encode(data)
fmt.Println(encoded) // 输出: aGVsbG8gd29ybGQ

// 解码
decoded, err := base64x.URLEncoding.Decode(encoded)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(decoded)) // 输出: hello world
```

### 大整数 Base64URL

```go
import (
    "math/big"
    "github.com/go4x/goal/ciphers/base64x"
)

// 将大整数编码为 Base64URL
bi := big.NewInt(12345)
encoded := base64x.Base64UrlUint.Encode(bi)
fmt.Println(encoded)

// 将 Base64URL 解码为大整数
decoded, err := base64x.Base64UrlUint.Decode(encoded)
if err != nil {
    log.Fatal(err)
}
fmt.Println(decoded) // 输出: 12345
```

## 使用场景

### URL 缩短

```go
import "github.com/go4x/goal/ciphers/base"

// 使用 Base62 缩短 URL
urlID := 123456789
shortCode := base.Base62(urlID)
shortURL := "https://short.ly/" + shortCode
fmt.Println(shortURL) // 输出: https://short.ly/8M0k5
```

### 数据库 ID 压缩

```go
import "github.com/go4x/goal/ciphers/base"

// 压缩数据库 ID 用于 API 响应
userID := 987654321
compressedID := base.Base62(userID)
fmt.Println(compressedID) // 输出: 15lOj
```

### 令牌生成

```go
import (
    "crypto/rand"
    "github.com/go4x/goal/ciphers/hash"
)

// 生成安全令牌
func generateToken() string {
    randomBytes := make([]byte, 32)
    rand.Read(randomBytes)
    return hash.SHA256(string(randomBytes))
}
```

### 数据完整性验证

```go
import "github.com/go4x/goal/ciphers/hash"

// 验证数据完整性
originalData := "重要数据"
originalHash := hash.SHA256(originalData)

// 稍后，验证数据未被篡改
receivedData := "重要数据"
receivedHash := hash.SHA256(receivedData)

if originalHash == receivedHash {
    fmt.Println("数据完整性已验证")
} else {
    fmt.Println("数据已被篡改")
}
```

## 安全最佳实践

### 加密

1. **使用 CBC 模式和随机 IV** 进行安全加密
2. **永远不要在生产环境中使用 ECB 模式**
3. **生成密码学随机的密钥和 IV**
4. **永远不要为同一密钥重复使用 IV**
5. **使用适当的密钥大小**（AES-256 需要 32 字节）

### 哈希

1. **使用 SHA256** 进行通用哈希
2. **使用 BCrypt** 进行密码哈希
3. **避免在安全关键应用中使用 MD5**
4. **在哈希前为密码添加盐值**

### 数据压缩

1. **使用 Base62** 进行区分大小写的应用
2. **使用 Base36** 进行不区分大小写的应用
3. **使用 Base64URL** 进行 URL 安全编码
4. **在编码/解码前验证输入**

## API 参考

### AES 加密

| 函数 | 描述 |
|------|------|
| `AES.Encrypt(data, key, mode, iv)` | 使用 AES 加密数据 |
| `AES.Decrypt(data, key, mode, iv)` | 使用 AES 解密数据 |

### 哈希函数

| 函数 | 描述 |
|------|------|
| `hash.SHA256(str)` | 生成 SHA256 哈希 |
| `hash.MD5(str)` | 生成 MD5 哈希 |
| `hash.BCrypt(password)` | 使用 BCrypt 哈希密码 |
| `hash.BCryptVerify(password, hash)` | 验证 BCrypt 密码 |

### 数据压缩

| 函数 | 描述 |
|------|------|
| `base.Base36(num)` | 将数字编码为 Base36 |
| `base.Base36Decode(str)` | 将 Base36 解码为数字 |
| `base.Base62(num)` | 将数字编码为 Base62 |
| `base.Base62Decode(str)` | 将 Base62 解码为数字 |

### Base64 变体

| 函数 | 描述 |
|------|------|
| `base64x.StdEncoding.Encode(data)` | 标准 Base64 编码 |
| `base64x.StdEncoding.Decode(str)` | 标准 Base64 解码 |
| `base64x.URLEncoding.Encode(data)` | URL 安全 Base64 编码 |
| `base64x.URLEncoding.Decode(str)` | URL 安全 Base64 解码 |
| `base64x.Base64UrlUint.Encode(bi)` | 大整数 Base64URL 编码 |
| `base64x.Base64UrlUint.Decode(str)` | Base64URL 解码为大整数 |

## 性能考虑

- **AES 加密/解密** 速度快但需要适当的密钥管理
- **BCrypt** 故意设计得较慢以防止暴力攻击
- **Base36/Base62** 压缩对于小数字非常快
- **SHA256** 速度快，适合大多数哈希需求
- **MD5** 速度快但密码学上已被破解

## 线程安全

此包中的所有函数都是线程安全的，可以从多个 goroutine 并发调用。

## 许可证

此包是 goal 项目的一部分，遵循相同的许可证条款。
