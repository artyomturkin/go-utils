package utils

import (
	"context"
	"sync"
	"testing"
)

func TestGracefullSHutdwonContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())

	ctxGS, cancelGS := context.WithCancel(context.TODO())
	wg := &sync.WaitGroup{}

	go cancelGS()

	GracefulShutdown(ctxGS, cancel, wg)

	if ctx.Err() != nil {
		t.Error("expected context to be canceled")
	}
}
