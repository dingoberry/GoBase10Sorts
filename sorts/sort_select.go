package sorts

import "cmp"

// SortSelect TimeComplex[Avg:O(n2),Bad:O(n2),Good:O(n2)] SpaceComplex:O(1) Unstable
func SortSelect[S ~[]E, E cmp.Ordered](datum S) {
	for i, size := 0, len(datum); i < size-1; i++ {
		m := i + 1
		for j := i; j < size; j++ {
			if datum[j] < datum[m] {
				m = j
			}
		}
		Swap(datum, i, m)
	}
}
