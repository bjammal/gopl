package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

func main() {
	data := make([]int, 10)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	fmt.Println(data)
}

func (t *tree) String() string {
	var buf bytes.Buffer
	if (*t).left != nil {
		buf.WriteString((*t).left.String())
	}
	if t != nil {
		fmt.Fprintf(&buf, "%d ", (*t).value)
	}

	if (*t).right != nil {
		buf.WriteString((*t).right.String())
	}
	return buf.String()
}

//!+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	fmt.Println(root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
