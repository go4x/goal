package errorx_test

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/go4x/goal/errorx"
	"github.com/stretchr/testify/assert"
)

func TestRescue(t *testing.T) {
	var count int32
	assert.NotPanics(t, func() {
		defer errorx.Recover(func(r any) {
			fmt.Println(r)
		}, func() {
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
		defer errorx.RecoverCtx(context.Background(), func(r any) {
			fmt.Println(r)
		}, func() {
			atomic.AddInt32(&count, 2)
		}, func() {
			atomic.AddInt32(&count, 3)
		})

		panic("hello")
	})
	assert.Equal(t, int32(5), atomic.LoadInt32(&count))
}
