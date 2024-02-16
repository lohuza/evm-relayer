package pkg

const EmptyString = ""

func GetPointerOf[T any](value T) *T {
	return &value
}
