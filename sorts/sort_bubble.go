package sorts

import "cmp"

// SortBubble TimeComplex[Avg:O(n2),Bad:O(n2),Good:O(n)] SpaceComplex:O(1) Stable
func SortBubble[S ~[]E, E cmp.Ordered](datum S) {
	for i, size := 0, len(datum); i < size; i++ {
		for j := 1; j < size-i; j++ {
			v1 := datum[j-1]
			v2 := datum[j]
			if v1 > v2 {
				datum[j-1] = v2
				datum[j] = v1
			}
		}
	}
}
