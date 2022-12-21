package util

func PointerFor[T interface{}](value T) *T {
	return &value
}
