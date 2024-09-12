package sorts

import "cmp"

func SortHeapifyImpl[S ~[]E, E cmp.Ordered](datum S, size, index int) {
	SortHeapify(datum, size, index)
}

func SortHeapify[S ~[]E, E cmp.Ordered](datum S, size, index int) {
	if index >= size {
		return
	}

	c1 := 2*index + 1
	c2 := 2*index + 2
	var top = index
	if c1 < size && datum[c1] > datum[top] {
		top = c1
	}
	if c2 < size && datum[c2] > datum[top] {
		top = c2
	}
	if top != index {
		Swap(datum, top, index)
		SortHeapifyImpl(datum, size, top)
	}
}

// SortHeap TimeComplex[Avg:O(nLogn),Bad:O(nLogn),Good:O(nLogn)] SpaceComplex:O(1) Unstable
func SortHeap[S ~[]E, E cmp.Ordered](datum S) {
	size := len(datum)
	lastNode := size - 1
	for i := (lastNode - 1) / 2; i >= 0; i-- {
		SortHeapify(datum, size, i)
	}

	for i := size - 1; i >= 0; i-- {
		Swap(datum, i, 0)
		SortHeapify(datum, i, 0)
	}
}
