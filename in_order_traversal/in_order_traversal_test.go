package main

import (
	"testing"
)

func TestGeneral(t *testing.T) {
	tree := &Node[int]{Value: 1}
	tree.Left = &Node[int]{Value: 2}
	tree.Right = &Node[int]{Value: 3}
	tree.Left.Left = &Node[int]{Value: 4}
	tree.Left.Right = &Node[int]{Value: 5}
	tree.Right.Left = &Node[int]{Value: 6}
	tree.Right.Right = &Node[int]{Value: 7}
	tree.Left.Left.Left = &Node[int]{Value: 8}
	tree.Left.Left.Right = &Node[int]{Value: 9}
	tree.Left.Right.Left = &Node[int]{Value: 10}
	tree.Left.Right.Right = &Node[int]{Value: 11}
	tree.Right.Left.Left = &Node[int]{Value: 12}

	expected := []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 3, 7}
	actual := in_order_traversal(tree, print_element[int])

	if len(expected) != len(actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	for i, v := range expected {
		if v != actual[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestEmpty(t *testing.T) {
	tree := &Node[int]{Value: 1}

	expected := []int{1}
	actual := in_order_traversal(tree, print_element[int])

	if len(expected) != len(actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	for i, v := range expected {
		if v != actual[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestLeftOnly(t *testing.T) {
	tree := &Node[int]{Value: 1}
	tree.Left = &Node[int]{Value: 2}
	tree.Left.Left = &Node[int]{Value: 3}
	tree.Left.Left.Left = &Node[int]{Value: 4}
	tree.Left.Left.Left.Left = &Node[int]{Value: 5}

	expected := []int{5, 4, 3, 2, 1}
	actual := in_order_traversal(tree, print_element[int])

	if len(expected) != len(actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	for i, v := range expected {
		if v != actual[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestNil(t *testing.T) {
	expected := []int{}
	actual := in_order_traversal(nil, print_element[int])

	if len(expected) != len(actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	for i, v := range expected {
		if v != actual[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}

	tree := &Node[int]{Value: 1}
	tree.Left = &Node[int]{Value: 2}
	tree.Left.Left = &Node[int]{Value: 3}
	tree.Left.Left.Left = &Node[int]{Value: 4}
	tree.Left.Left.Left.Left = &Node[int]{Value: 5}

	expected = []int{5, 4, 3, 2, 1}
	actual = in_order_traversal(tree, nil)

	if len(expected) != len(actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	for i, v := range expected {
		if v != actual[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}
