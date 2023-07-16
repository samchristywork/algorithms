package main

import (
	"math/rand"
	"sort"
	"testing"
)

func TestGeneral(t *testing.T) {
	data := []int{1, 4, 2, 5, 3}
	expected := []int{1, 2, 3, 4, 5}

	actual := merge_sort(data)

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
	}
}

func TestLotsOfData(t *testing.T) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = rand.Int()
	}

	expected := make([]int, 10000)
	copy(expected, data)
	sort.Ints(expected)

	actual := merge_sort(data)

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
	}
}
