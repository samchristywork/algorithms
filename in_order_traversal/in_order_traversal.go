package main

import (
  "fmt"
)

type Node struct {
  Value int
  Left *Node
  Right *Node
}

func print_element(value int, depth int) {
  for i := 0; i < depth; i++ {
    fmt.Print("  ")
  }
  fmt.Println(value)
}

func _in_order_traversal(tree *Node, callback func(int, int), depth int) []int {
  var values []int

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

func in_order_traversal(tree *Node, callback func(int, int)) []int {
  if tree == nil {
    return []int{}
  }

  return _in_order_traversal(tree, callback, 0)
}
