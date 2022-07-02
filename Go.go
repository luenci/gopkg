package gopkg

import (
	"context"
	"log"
	"runtime/debug"
)

// Go 友好的 goroutine 异常捕获, 防止野生 goroutine 的 panic.

func Go(ctx context.Context, handler func()) {

	go func() {
		defer func() {
			if e := recover(); e != nil {
				buf := debug.Stack()
				log.Panicf("safe go handler panic: %+v", map[string]interface{}{
					"stack": string(buf),
					"err":   e,
				})
			}
		}()
		handler()
	}()
}
