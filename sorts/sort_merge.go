package sorts

import (
	"cmp"
)

// SortMerge TimeComplex[Avg:O(nLogn),Bad:O(nLogn),Good:O(nLogn)] SpaceComplex:O(n) Stable
func SortMerge[S ~[]E, E cmp.Ordered](datum S) {
	for i, size := 1, len(datum); i < size; i *= 2 {
		for j := 0; j < size; j += i * 2 {
			SortQuickPartition(datum, j, min(j+i*2, size)-1)
		}
	}
}
