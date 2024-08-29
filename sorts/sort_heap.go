package sorts

import "cmp"

func SortHeapifyImpl[S ~[]E, E cmp.Ordered](datum S, n, i int) {
	SortHeapify(datum, n, i)
}

func SortHeapify[S ~[]E, E cmp.Ordered](datum S, n, i int) {
	if i >= n {
		return
	}

	c1 := 2*i + 1
	c2 := 2*i + 2
	var m = i
	if c1 < n && datum[c1] > datum[m] {
		m = c1
	}
	if c2 < n && datum[c2] > datum[m] {
		m = c2
	}
	if m != i {
		Swap(datum, m, i)
		SortHeapifyImpl(datum, n, m)
	}
}

// SortHeap TimeComplex[Avg:O(nLogn),Bad:O(nLogn),Good:O(nLogn)] SpaceComplex:O(1) Unstable
func SortHeap[S ~[]E, E cmp.Ordered](datum S) {
	size := len(datum)
	lastNode := size - 1
	parent := (lastNode - 1) / 2
	for i := parent; i >= 0; i-- {
		SortHeapify(datum, size, i)
	}

	for i := size - 1; i >= 0; i-- {
		Swap(datum, i, 0)
		SortHeapify(datum, i, 0)
	}
}
