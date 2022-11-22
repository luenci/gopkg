package gopkg

import (
	"context"
	"runtime/debug"

	"github.com/luenci/gopkg/log"
)

// Go 友好的 goroutine 异常捕获, 防止野生 goroutine 的 panic.

func Go(ctx context.Context, handler func()) {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				buf := debug.Stack()
				log.Errors("safe go handler panic: \n", map[string]interface{}{
					"stack": string(buf),
					"err":   e,
				})
			}
		}()
		handler()
	}()
}
