package utils

import (
	"fmt"
	"github.com/v-mars/library/logs"
)

func SafeGoroutine(fn func()) {
	var err error
	go func() {
		defer func() {
			if r := recover(); r != nil {
				var ok bool
				err, ok = r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				logs.Errorf("goroutine panic: %v", err)
			}
		}()
		fn()
	}()
}
