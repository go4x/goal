package limiter

import (
	"testing"
	"time"
)

// TestTokenBucketImplementsLimiter 验证 TokenBucket 是否实现了 Limiter 接口
func TestTokenBucketImplementsLimiter(t *testing.T) {
	// 这个测试会在编译时验证接口实现关系
	var _ Limiter = &TokenBucket{}

	// 创建实例并测试所有接口方法
	bucket := NewTokenBucket(10, 5, time.Second)

	// 测试 TryTake
	_ = bucket.TryTake()

	// 测试 Take (在 goroutine 中避免阻塞)
	go func() {
		bucket.Take()
	}()

	// 测试 TakeWithTimeout
	_ = bucket.TakeWithTimeout(100 * time.Millisecond)

	// 测试 Stat
	total, blocked, successRate := bucket.Stat()
	_ = total
	_ = blocked
	_ = successRate

	t.Log("TokenBucket successfully implements Limiter interface")
}

// TestLimiterInterface 测试接口的多态性
func TestLimiterInterface(t *testing.T) {
	// 使用接口类型
	var limiter Limiter = NewTokenBucket(10, 5, time.Second)

	// 测试所有接口方法
	_ = limiter.TryTake()
	_ = limiter.TakeWithTimeout(100 * time.Millisecond)
	total, blocked, successRate := limiter.Stat()

	_ = total
	_ = blocked
	_ = successRate

	t.Log("Interface polymorphism works correctly")
}
