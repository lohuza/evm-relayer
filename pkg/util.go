package pkg

import (
	"context"
	"sync"
	"time"
)

const EmptyString = ""

func GetPointerOf[T any](value T) *T {
	return &value
}

func GetTimeUtc() time.Time {
	return time.Now().UTC()
}

func GetUtcTimeSecond() int64 {
	return time.Now().UTC().Unix()
}

func HandleShutdown(ctx context.Context, wg *sync.WaitGroup, handleShutdown func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		handleShutdown()
	}()
}

func ExecuteWithRetry[T any](funcToExecute func() (T, error)) (T, error) {
	var result T
	var err error
	for i := 1; i <= 3; i++ {
		result, err = funcToExecute()
		if err == nil {
			return result, nil
		}
		time.Sleep(time.Millisecond * time.Duration(i) * 100)
	}
	return result, err
}
