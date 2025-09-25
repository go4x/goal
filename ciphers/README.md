# ciphers

A comprehensive cryptographic package for Go that provides encryption, decryption, hashing, and data compression functionality.

## Features

- **AES Encryption/Decryption**: Support for ECB and CBC modes with security warnings
- **Hash Functions**: SHA256, MD5, and BCrypt for various hashing needs
- **Data Compression**: Base36, Base62, and Base64 encoding for compact data representation
- **Base64 Variants**: Standard, URL-safe, and Base64URL encoding for different use cases
- **Security-First Design**: Clear warnings about insecure modes and best practices

## Installation

```bash
go get github.com/go4x/goal/ciphers
```

## Quick Start

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
    // AES encryption (secure CBC mode)
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
    
    fmt.Println(string(decrypted)) // Output: Hello World
    
    // Hash functions
    hash := hash.SHA256("hello world")
    fmt.Println(hash)
    
    // Data compression
    compressed := base.Base62(12345)
    fmt.Println(compressed) // Output: 3d7
}
```

## AES Encryption/Decryption

### Secure Encryption (Recommended)

```go
import "github.com/go4x/goal/ciphers"

// Use CBC mode with random IV for secure encryption
data := []byte("sensitive data")
key := []byte("your-32-byte-key-here-123456789012") // 32 bytes for AES-256
iv := []byte("random-16-byte-iv") // 16 bytes for AES

// Encrypt
encrypted, err := ciphers.AES.Encrypt(data, key, ciphers.CBC, iv)
if err != nil {
    log.Fatal(err)
}

// Decrypt
decrypted, err := ciphers.AES.Decrypt(encrypted, key, ciphers.CBC, iv)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(decrypted)) // Output: sensitive data
```

### Security Warnings

⚠️ **Important Security Notes:**

- **ECB mode is NOT secure** and should NOT be used in production
- Always use **CBC mode with a random IV** for secure encryption
- Ensure keys are cryptographically random and properly sized
- Never reuse IVs for the same key

```go
// ❌ DO NOT USE - Insecure ECB mode
encrypted, err := ciphers.AES.Encrypt(data, key, ciphers.ECB, nil)

// ✅ USE THIS - Secure CBC mode
encrypted, err := ciphers.AES.Encrypt(data, key, ciphers.CBC, randomIV)
```

## Hash Functions

### SHA256

```go
import "github.com/go4x/goal/ciphers/hash"

// SHA256 hashing
hash := hash.SHA256("hello world")
fmt.Println(hash) // Output: b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
```

### MD5

```go
import "github.com/go4x/goal/ciphers/hash"

// MD5 hashing (not recommended for security-critical applications)
hash := hash.MD5("hello world")
fmt.Println(hash) // Output: 5d41402abc4b2a76b9719d911017c592
```

### BCrypt (Recommended for Passwords)

```go
import "github.com/go4x/goal/ciphers/hash"

// Hash password
password := "mysecretpassword"
hashedPassword, err := hash.BCrypt(password)
if err != nil {
    log.Fatal(err)
}

// Verify password
isValid := hash.BCryptVerify(password, hashedPassword)
fmt.Println(isValid) // Output: true
```

## Data Compression

### Base36 Encoding

```go
import "github.com/go4x/goal/ciphers/base"

// Encode number to Base36
compressed := base.Base36(12345)
fmt.Println(compressed) // Output: 9IX

// Decode Base36 to number
number, ok := base.Base36Decode("9IX")
if ok {
    fmt.Println(number) // Output: 12345
}
```

### Base62 Encoding

```go
import "github.com/go4x/goal/ciphers/base"

// Encode number to Base62
compressed := base.Base62(12345)
fmt.Println(compressed) // Output: 3d7

// Decode Base62 to number
number, ok := base.Base62Decode("3d7")
if ok {
    fmt.Println(number) // Output: 12345
}
```

## Base64 Variants

### Standard Base64

```go
import "github.com/go4x/goal/ciphers/base64x"

// Standard base64 encoding
data := []byte("hello world")
encoded := base64x.StdEncoding.Encode(data)
fmt.Println(encoded) // Output: aGVsbG8gd29ybGQ=

// Decoding
decoded, err := base64x.StdEncoding.Decode(encoded)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(decoded)) // Output: hello world
```

### URL-Safe Base64

```go
import "github.com/go4x/goal/ciphers/base64x"

// URL-safe base64 encoding
data := []byte("hello world")
encoded := base64x.URLEncoding.Encode(data)
fmt.Println(encoded) // Output: aGVsbG8gd29ybGQ

// Decoding
decoded, err := base64x.URLEncoding.Decode(encoded)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(decoded)) // Output: hello world
```

### Base64URL for Unsigned Integers

```go
import (
    "math/big"
    "github.com/go4x/goal/ciphers/base64x"
)

// Encode big integer to Base64URL
bi := big.NewInt(12345)
encoded := base64x.Base64UrlUint.Encode(bi)
fmt.Println(encoded)

// Decode Base64URL to big integer
decoded, err := base64x.Base64UrlUint.Decode(encoded)
if err != nil {
    log.Fatal(err)
}
fmt.Println(decoded) // Output: 12345
```

## Use Cases

### URL Shortening

```go
import "github.com/go4x/goal/ciphers/base"

// Shorten URLs using Base62
urlID := 123456789
shortCode := base.Base62(urlID)
shortURL := "https://short.ly/" + shortCode
fmt.Println(shortURL) // Output: https://short.ly/8M0k5
```

### Database ID Compression

```go
import "github.com/go4x/goal/ciphers/base"

// Compress database IDs for API responses
userID := 987654321
compressedID := base.Base62(userID)
fmt.Println(compressedID) // Output: 15lOj
```

### Token Generation

```go
import (
    "crypto/rand"
    "github.com/go4x/goal/ciphers/hash"
)

// Generate secure tokens
func generateToken() string {
    randomBytes := make([]byte, 32)
    rand.Read(randomBytes)
    return hash.SHA256(string(randomBytes))
}
```

### Data Integrity Verification

```go
import "github.com/go4x/goal/ciphers/hash"

// Verify data integrity
originalData := "important data"
originalHash := hash.SHA256(originalData)

// Later, verify the data hasn't been tampered with
receivedData := "important data"
receivedHash := hash.SHA256(receivedData)

if originalHash == receivedHash {
    fmt.Println("Data integrity verified")
} else {
    fmt.Println("Data has been tampered with")
}
```

## Security Best Practices

### Encryption

1. **Use CBC mode with random IVs** for secure encryption
2. **Never use ECB mode** in production
3. **Generate cryptographically random keys and IVs**
4. **Never reuse IVs** for the same key
5. **Use appropriate key sizes** (32 bytes for AES-256)

### Hashing

1. **Use SHA256** for general-purpose hashing
2. **Use BCrypt** for password hashing
3. **Avoid MD5** for security-critical applications
4. **Add salt** to passwords before hashing

### Data Compression

1. **Use Base62** for case-sensitive applications
2. **Use Base36** for case-insensitive applications
3. **Use Base64URL** for URL-safe encoding
4. **Validate input** before encoding/decoding

## API Reference

### AES Encryption

| Function | Description |
|----------|-------------|
| `AES.Encrypt(data, key, mode, iv)` | Encrypt data with AES |
| `AES.Decrypt(data, key, mode, iv)` | Decrypt data with AES |

### Hash Functions

| Function | Description |
|----------|-------------|
| `hash.SHA256(str)` | Generate SHA256 hash |
| `hash.MD5(str)` | Generate MD5 hash |
| `hash.BCrypt(password)` | Hash password with BCrypt |
| `hash.BCryptVerify(password, hash)` | Verify BCrypt password |

### Data Compression

| Function | Description |
|----------|-------------|
| `base.Base36(num)` | Encode number to Base36 |
| `base.Base36Decode(str)` | Decode Base36 to number |
| `base.Base62(num)` | Encode number to Base62 |
| `base.Base62Decode(str)` | Decode Base62 to number |

### Base64 Variants

| Function | Description |
|----------|-------------|
| `base64x.StdEncoding.Encode(data)` | Standard Base64 encoding |
| `base64x.StdEncoding.Decode(str)` | Standard Base64 decoding |
| `base64x.URLEncoding.Encode(data)` | URL-safe Base64 encoding |
| `base64x.URLEncoding.Decode(str)` | URL-safe Base64 decoding |
| `base64x.Base64UrlUint.Encode(bi)` | Base64URL encoding for big integers |
| `base64x.Base64UrlUint.Decode(str)` | Base64URL decoding to big integers |

## Performance Considerations

- **AES encryption/decryption** is fast but requires proper key management
- **BCrypt** is intentionally slow to prevent brute force attacks
- **Base36/Base62** compression is very fast for small numbers
- **SHA256** is fast and suitable for most hashing needs
- **MD5** is fast but cryptographically broken

## Thread Safety

All functions in this package are thread-safe and can be called concurrently from multiple goroutines.

## License

This package is part of the goal project and follows the same license terms.
