package main

import (
	"testing"
)

type point struct {
	x int
	y int
}

func reverse(s []token) []token {
	r := make([]token, len(s))
	for i := range s {
		r[len(s)-1-i] = s[i]
	}
	return r
}

func assertEqual(t *testing.T, expected []token, actual []token) {
	expected = reverse(expected)

	if len(expected) != len(actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	for i := range expected {
		if expected[i].kind != actual[i].kind {
			t.Errorf("\nExpected:\t%v\nGot:     \t%v", expected, actual)
			return
		}
		if expected[i].val != actual[i].val {
			t.Errorf("\nExpected:\t%v\nGot:     \t%v", expected, actual)
			return
		}
	}
}

func TestString(t *testing.T) {
	r := tokenizeString("The fat dog")
	c := tokenizeString("The great dog is here")
	lcs := lcs(r, c)

	expected := []token{
		{kind: unchanged, val: "The"},
		{kind: removed, val: "fat"},
		{kind: added, val: "great"},
		{kind: unchanged, val: "dog"},
		{kind: added, val: "is"},
		{kind: added, val: "here"},
	}

	assertEqual(t, expected, diff(r, c, lcs))

	d := diff(r, c, lcs)
	printDiff(d)
}

func TestInt(t *testing.T) {
	r := tokenize([]int{1, 2, 3, 4, 5})
	c := tokenize([]int{1, 3, 5, 7, 9})
	lcs := lcs(r, c)

	expected := []token{
		{kind: unchanged, val: 1},
		{kind: removed, val: 2},
		{kind: unchanged, val: 3},
		{kind: removed, val: 4},
		{kind: unchanged, val: 5},
		{kind: added, val: 7},
		{kind: added, val: 9},
	}

	assertEqual(t, expected, diff(r, c, lcs))

	d := diff(r, c, lcs)
	printDiff(d)
}

func Test(t *testing.T) {
	r := tokenize([]float64{1.1, 2.2, 3.3, 4.4, 5.5})
	c := tokenize([]float64{1.1, 3.3, 5.5, 7.7, 9.9})
	lcs := lcs(r, c)

	expected := []token{
		{kind: unchanged, val: 1.1},
		{kind: removed, val: 2.2},
		{kind: unchanged, val: 3.3},
		{kind: removed, val: 4.4},
		{kind: unchanged, val: 5.5},
		{kind: added, val: 7.7},
		{kind: added, val: 9.9},
	}

	assertEqual(t, expected, diff(r, c, lcs))

	d := diff(r, c, lcs)
	printDiff(d)
}

func TestStruct(t *testing.T) {
	r := tokenize([]point{{1, 2}, {3, 4}, {5, 6}})
	c := tokenize([]point{{1, 2}, {3, 5}, {5, 6}})
	lcs := lcs(r, c)

	expected := []token{
		{kind: unchanged, val: point{1, 2}},
		{kind: removed, val: point{3, 4}},
		{kind: added, val: point{3, 5}},
		{kind: unchanged, val: point{5, 6}},
	}

	assertEqual(t, expected, diff(r, c, lcs))

	d := diff(r, c, lcs)
	printDiff(d)
}
