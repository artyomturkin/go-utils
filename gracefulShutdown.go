package utils

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// GracefullShutdown waits for either ctx to be cancelled or a signal from OS. Calls supplied cancel func and waits for wait group to be done
func GracefulShutdown(ctx context.Context, cancel func(), wg *sync.WaitGroup) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
	case <-sig:
	}

	wg.Wait()
}
