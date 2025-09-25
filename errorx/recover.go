package errorx

import (
	"context"
)

// Recover is used with defer to do cleanup on panics.
// Use it like:
//
//	defer Recover(func(r any) { fmt.Println(r) }, func() { /* cleanup code */ })
func Recover(f func(r any), cleanups ...func()) {
	if p := recover(); p != nil {
		// Execute cleanup functions when panic occurs
		for _, cleanup := range cleanups {
			func() {
				defer func() {
					if r := recover(); r != nil {
						// Handle panics in cleanup functions
						if f != nil {
							f(r)
						}
					}
				}()
				cleanup()
			}()
		}

		// callback to handle the panic information
		if f != nil {
			f(p)
		}
	}
}

// RecoverCtx is used with defer to do cleanup on panics with context support.
// It respects context cancellation and timeout during cleanup.
func RecoverCtx(ctx context.Context, f func(r any), cleanups ...func()) {
	if p := recover(); p != nil {
		// Check if context is already cancelled
		select {
		case <-ctx.Done():
			if f != nil {
				f(p)
			}
			return
		default:
		}

		// Execute cleanup functions with context awareness
		for _, cleanup := range cleanups {
			select {
			case <-ctx.Done():
				return // Stop cleanup if context is cancelled
			default:
				func() {
					defer func() {
						if r := recover(); r != nil {
							// Handle panics in cleanup functions
							if f != nil {
								f(r)
							}
						}
					}()
					cleanup()
				}()
			}
		}

		// callback to handle the panic information
		if f != nil {
			f(p)
		}
	}
}
