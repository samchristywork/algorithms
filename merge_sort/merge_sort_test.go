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

func TestEmpty(t *testing.T) {
	data := []int{}
	expected := []int{}

	actual := merge_sort(data)

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
	}
}

func TestSingle(t *testing.T) {
	data := []int{1}
	expected := []int{1}

	actual := merge_sort(data)

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
	}
}

func TestFloats(t *testing.T) {
	data := []float64{1.0, 4.0, 2.0, 5.0, 3.0}
	expected := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	actual := merge_sort(data)

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("Expected %f, got %f", expected[i], v)
		}
	}
}

func TestStrings(t *testing.T) {
	data := []string{"foo", "bar", "baz"}
	expected := []string{"bar", "baz", "foo"}

	actual := merge_sort(data)

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], v)
		}
	}
}
