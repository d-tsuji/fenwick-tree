// Package fenwick is an implementation of the Fenwick tree.
// This is a 0-indexed Fenwick tree.
//
// A Fenwick tree is a data structure
// that can efficiently update elements
// and calculate prefix sums in a table of numbers.
// ref: https://en.wikipedia.org/wiki/Fenwick_tree
package fenwick

type Tree struct {
	n   int
	bit []int
}

// NewTree initializes a slice of the Fenwick tree.
func NewTree(n int) *Tree {
	return &Tree{
		n:   n,
		bit: make([]int, n+1),
	}
}

// Add adds the value x to the element array[i].
func (t *Tree) Add(i, x int) {
	for i <= t.n {
		t.bit[i] += x
		i |= i + 1
	}
}

// Query computes and returns the sum of array[l, r].
func (t *Tree) Query(l, r int) int {
	var ret int
	if l <= r {
		ret = t.query(r) - t.query(l-1)
	}
	return ret
}

// query computes and returns the sum of array[0, r].
func (t *Tree) query(i int) int {
	var ret int
	for i >= 0 {
		ret += t.bit[i]
		i = (i & (i + 1)) - 1
	}
	return ret
}
