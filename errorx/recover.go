package errorx

import (
	"context"
	"runtime/debug"

	"github.com/go4x/logx"
)

// Recover is used with defer to do cleanup on panics.
// Use it like:
//
//	defer Recover(func() {})
func Recover(logger logx.Logger, cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		logger.Error("Recover", "error", p, "stack", debug.Stack())
	}
}

// RecoverCtx is used with defer to do cleanup on panics.
func RecoverCtx(ctx context.Context, logger logx.Logger, cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		logger.Error("Recover", "error", p, "stack", debug.Stack())
	}
}
