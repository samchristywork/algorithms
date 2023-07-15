package main

import (
	"fmt"
)

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

func (t *DirectAddressTable[T]) Insert(key string, value T) {
	hash := hashKey(key)
	t.table[hash] = append(t.table[hash], TableElement[T]{key, value})
	t.numElements++
}

func (t *DirectAddressTable[T]) Delete(key string) {
	hash := hashKey(key)
	for i, element := range t.table[hash] {
		if element.key == key {
			t.table[hash] = append(t.table[hash][:i], t.table[hash][i+1:]...)
			t.numElements--
		}
	}
}

func (t *DirectAddressTable[T]) Find(key string) *T {
	hash := hashKey(key)
	for _, element := range t.table[hash] {
		if element.key == key {
			return &element.value
		}
	}

	return nil
}

func (t *DirectAddressTable[T]) Update(key string, value T) {
	hash := hashKey(key)
	for i, element := range t.table[hash] {
		if element.key == key {
			t.table[hash][i].value = value
		}
	}
}

func (t *DirectAddressTable[T]) Print() {
	for i := 0; i < MAX_KEY; i++ {
		fmt.Printf("%d: %v\n", i, t.table[i])
	}
}
