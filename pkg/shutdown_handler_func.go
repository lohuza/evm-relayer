package pkg

import (
	"context"
	"sync"
)

func HandleShutdown(ctx context.Context, wg *sync.WaitGroup, handleShutdown func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		handleShutdown()
	}()
}
