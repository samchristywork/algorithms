package main

import (
	"fmt"
	"testing"
)

func TestGeneral(t *testing.T) {
	table := DirectAddressTable[int]{}
	table.Insert("foo", 1)
	table.Insert("bar", 20)
	table.Insert("baz", 30)
	table.Insert("qux", 40)
	table.Insert("quux", 50)
	table.Insert("corge", 60)
	table.Insert("grault", 70)
	table.Insert("garply", 80)
	table.Insert("waldo", 90)
	table.Insert("fred", 100)
	table.Insert("plugh", 110)

	table.Delete("grault")
	table.Delete("qux")
	table.Delete("bar")

	foo := table.Find("foo")
	if *foo != 1 {
		t.Errorf("Expected 1, got %d", *foo)
	}

	faz := table.Find("faz")
	if faz != nil {
		t.Errorf("Expected nil, got %d", *faz)
	}

	grault := table.Find("grault")
	if grault != nil {
		t.Errorf("Expected nil, got %d", *grault)
	}

	for i := 0; i < MAX_KEY; i++ {
		fmt.Printf("%d: %v\n", i, table.table[i])
	}

	length := table.Len()
	if length != 8 {
		t.Errorf("Expected 8, got %d", length)
	}
}

func TestInsert(t *testing.T) {
	table := DirectAddressTable[int]{}
	table.Insert("foo", 1)
	table.Insert("bar", 20)
	table.Insert("baz", 30)
	table.Insert("qux", 40)
	table.Insert("quux", 50)
	table.Insert("corge", 60)
	table.Insert("grault", 70)
	table.Insert("garply", 80)
	table.Insert("waldo", 90)
	table.Insert("fred", 100)
	table.Insert("plugh", 110)

	foo := table.Find("foo")
	if *foo != 1 {
		t.Errorf("Expected 1, got %d", *foo)
	}

	faz := table.Find("faz")
	if faz != nil {
		t.Errorf("Expected nil, got %d", *faz)
	}

	grault := table.Find("grault")
	if *grault != 70 {
		t.Errorf("Expected 70, got %d", *grault)
	}

	for i := 0; i < MAX_KEY; i++ {
		fmt.Printf("%d: %v\n", i, table.table[i])
	}

	length := table.Len()
	if length != 11 {
		t.Errorf("Expected 11, got %d", length)
	}
}

func TestFloat(t *testing.T) {
	table := DirectAddressTable[float64]{}
	table.Insert("foo", 1.0)
	table.Insert("bar", 20.0)
	table.Insert("baz", 30.0)

	foo := table.Find("foo")
	if *foo != 1.0 {
		t.Errorf("Expected 1.0, got %f", *foo)
	}

	faz := table.Find("faz")
	if faz != nil {
		t.Errorf("Expected nil, got %f", *faz)
	}
}

func TestEmpty(t *testing.T) {
	table := DirectAddressTable[float64]{}

	foo := table.Find("foo")
	if foo != nil {
		t.Errorf("Expected nil, got %f", *foo)
	}

	length := table.Len()
	if length != 0 {
		t.Errorf("Expected 0, got %d", length)
	}
}
