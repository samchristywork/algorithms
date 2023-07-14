package main

import (
  "fmt"
)

type Node struct {
  Value int
  Left *Node
  Right *Node
}

func print_element(value int) {
  fmt.Println(value)
}

func in_order_traversal(tree *Node, callback func(int)) []int {
  var values []int

  if tree.Left != nil {
    values = append(values, in_order_traversal(tree.Left, callback)...)
  }

  values = append(values, tree.Value)
  callback(tree.Value, depth)

  if tree.Right != nil {
    values = append(values, in_order_traversal(tree.Right, callback)...)
  }

  return values
}
