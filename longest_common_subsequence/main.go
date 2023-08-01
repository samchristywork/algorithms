package main

import (
	"fmt"
)

const (
	unchanged = 0
	added     = 1
	removed   = 2
)

type token struct {
	kind int
	val  string
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(r []string, c []string) [][]int {
}

func main() {
	r := []string{"The", "fat", "dog"}
	c := []string{"The", "great", "dog", "is", "here"}
	lcs(r, c)
}
