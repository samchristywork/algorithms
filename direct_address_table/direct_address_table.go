package main

const (
	MAX_KEY = 10 // Very small for testing purposes
)

type TableElement[T interface{}] struct {
	key   string
	value T
}

type DirectAddressTable[T interface{}] struct {
	table       [MAX_KEY][]TableElement[T]
	numElements int
}

func hashString(key string) int {
	sum := 0
	for _, c := range key {
		sum += int(c)
	}
	return sum
}

func hashKey(key string) int {
	return hashString(key) % MAX_KEY
}

func (t *DirectAddressTable[T]) Len() int {
	return t.numElements
}
