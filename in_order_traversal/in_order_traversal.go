package main

import (
)

type Node struct {
  Value int
  Left *Node
  Right *Node
}

func in_order_traversal(tree *Node) []int {
  var values []int

  if tree.Left != nil {
    values = append(values, in_order_traversal(tree.Left)...)
  }

  values = append(values, tree.Value)

  if tree.Right != nil {
    values = append(values, in_order_traversal(tree.Right)...)
  }

  return values
}
