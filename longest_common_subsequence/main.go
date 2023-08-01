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
	table := make([][]int, len(r)+1)
	for i := range table {
		table[i] = make([]int, len(c)+1)
	}

	for i := 1; i < len(r)+1; i++ {
		for j := 1; j < len(c)+1; j++ {
			if r[i-1] == c[j-1] {
				table[i][j] = table[i-1][j-1] + 1
			} else {
				table[i][j] = max(table[i-1][j], table[i][j-1])
			}
		}
	}

	return table
}

func diff(r []string, c []string, table [][]int) []token {
	var output []token

	i := len(r)
	j := len(c)
	for i > 0 && j > 0 {
		if r[i-1] == c[j-1] {
			output = append(output, token{unchanged, r[i-1]})
			i--
			j--
		} else if table[i-1][j] > table[i][j-1] {
			output = append(output, token{removed, r[i-1]})
			i--
		} else {
			output = append(output, token{added, c[j-1]})
			j--
		}
	}

	for i > 0 {
		output = append(output, token{removed, r[i-1]})
		i--
	}

	for j > 0 {
		output = append(output, token{added, c[j-1]})
		j--
	}

	return output
}

func tokenizeString(s string) []string {
	var output []string
	var current string

	for _, c := range s {
		if c == ' ' {
			output = append(output, current)
			current = ""
		} else {
			current += string(c)
		}
	}

	if current != "" {
		output = append(output, current)
	}

	return output
}

func main() {
	r := tokenizeString("The fat dog")
	c := tokenizeString("The great dog is here")
	lcs := lcs(r, c)

	diff(r, c, lcs)
}
