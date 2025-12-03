# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-09-25

### Added
- **Initial Release**: Comprehensive Go utility library with 20+ specialized packages
- **Core Utilities**: assert, value, ptr, conv packages for basic operations
- **Collections**: col/set, col/mapx, col/slicex for advanced data structures
- **String & Text**: stringx, color, jsonx for text manipulation
- **System & I/O**: cmd, iox, httpx for system operations
- **Cryptography**: ciphers, uuid for security operations
- **Error Handling**: errorx for comprehensive error management
- **Mathematics**: mathx, prob, random for mathematical operations
- **Utilities**: timex, limiter, retry, reflectx, printer for various utilities
- **Generic Type Support**: Modern Go generics throughout the codebase
- **Performance Optimized**: O(1) operations for most collection operations
- **Thread Safety**: All packages designed for concurrent use
- **Minimal Dependencies**: Few external dependencies for easy maintenance
- **Bilingual Documentation**: Complete documentation in English and Chinese
- **Comprehensive Testing**: Extensive test coverage with benchmarks

### Features
- **HashSet, ArraySet, LinkedSet**: Multiple set implementations with different performance characteristics
- **Regular Map, ArrayMap, LinkedMap**: Map implementations with order preservation options
- **Enhanced Slice Operations**: Immutable operations with functional programming support
- **String Utilities**: Case conversion, email blurring, string building
- **HTTP Client**: Async HTTP operations with connection pooling
- **Command Execution**: Async command execution with timeout handling
- **AES Encryption**: Secure encryption with CBC mode support
- **UUID Generation**: Standard UUIDs and distributed IDs (Sonyflake)
- **Error Chaining**: Comprehensive error handling with recovery mechanisms
- **Rate Limiting**: Token bucket algorithm implementation
- **Retry Mechanisms**: Exponential backoff retry strategies
- **Time Utilities**: Advanced time formatting and parsing
- **Random Generation**: Number and string generation utilities

### Documentation
- Complete API documentation for all packages
- Comprehensive usage examples
- Performance characteristics and tips
- Bilingual support (English/Chinese)
- Installation and quick start guides

### Testing
- Unit tests for all packages
- Benchmark tests for performance-critical operations
- Example tests for documentation
- Comprehensive test coverage

## [1.1.0] - 2025-12-03

### Changed
- **value**: Moved generic check and comparison helpers (`Zero`, `NotZero`, `Nil`, `NotNil`, `Empty`, `NotEmpty`, `Equal`, `NotEqual`, `DeepEqual`) out of the `value` package into the new `is` package to better reflect their semantics.

### Added
- **is**:
  - Introduced a dedicated value checking and comparison package with:
    - Boolean helpers: `Not`, `True`, `False`
    - Zero / empty / nil helpers: `Zero`, `NotZero`, `Nil`, `NotNil`, `Empty`, `NotEmpty`
    - Equality helpers: `Eq`, `Neq`, `EqDeep`, `Same`
    - Ordered comparisons: `Gt`, `Gte`, `Lt`, `Lte` based on `cmp.Ordered`
  - `Eq` now supports:
    - Basic types via `==`
    - Pointers by comparing the pointed-to values (including nested pointers)
    - Interfaces by comparing underlying concrete values
    - Slices and maps via deep comparison
    - Channels by reference (standard Go behavior)
    - Functions, returning true only when both are nil
  - Added comprehensive tests covering:
    - Structs, pointers, interfaces, slices, maps, channels, functions
    - Ordered comparisons for integers and strings
  - Added bilingual (`README.md` / `README_CN.md`) documentation describing the new `is` API, behavior, and best practices.
