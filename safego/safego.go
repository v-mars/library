package safego

import (
	"context"
	"github.com/v-mars/library/goutil"
)

func Go(ctx context.Context, fn func()) {
	go func() {
		defer goutil.Recovery(ctx)

		fn()
	}()
}
