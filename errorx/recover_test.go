package errorx_test

import (
	"context"
	"sync/atomic"
	"testing"

	"github.com/gophero/goal/errorx"
	"github.com/gophero/goal/tests"
	"github.com/stretchr/testify/assert"
)

func TestRescue(t *testing.T) {
	var count int32
	assert.NotPanics(t, func() {
		defer errorx.Recover(tests.NewLog(), func() {
			atomic.AddInt32(&count, 2)
		}, func() {
			atomic.AddInt32(&count, 3)
		})

		panic("hello")
	})
	assert.Equal(t, int32(5), atomic.LoadInt32(&count))
}

func TestRescueCtx(t *testing.T) {
	var count int32
	assert.NotPanics(t, func() {
		defer errorx.RecoverCtx(context.Background(), tests.NewLog(), func() {
			atomic.AddInt32(&count, 2)
		}, func() {
			atomic.AddInt32(&count, 3)
		})

		panic("hello")
	})
	assert.Equal(t, int32(5), atomic.LoadInt32(&count))
}
