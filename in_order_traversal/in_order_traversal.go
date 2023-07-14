package main

import (
  "fmt"
)

type Node[T interface{}] struct {
  Value T
  Left *Node[T]
  Right *Node[T]
}

func print_element[T any](value T, depth int) {
  for i := 0; i < depth; i++ {
    fmt.Print("  ")
  }
  fmt.Println(value)
}

func _in_order_traversal[T any](tree *Node[T], callback func(T, int), depth int) []T {
  var values []T

  if tree.Left != nil {
    values = append(values, _in_order_traversal(tree.Left, callback, depth+1)...)
  }

  values = append(values, tree.Value)
  if callback != nil {
    callback(tree.Value, depth)
  }

  if tree.Right != nil {
    values = append(values, _in_order_traversal(tree.Right, callback, depth+1)...)
  }

  return values
}

func in_order_traversal(tree *Node[int], callback func(int, int)) []int {
  if tree == nil {
    return []int{}
  }

  return _in_order_traversal(tree, callback, 0)
}
