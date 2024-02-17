package pkg

import "time"

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
