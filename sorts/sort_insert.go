package sorts

import "cmp"

// SortInsert TimeComplex[Avg:O(n2),Bad:O(n2),Good:O(n)] SpaceComplex:O(1) Stable
func SortInsert[S ~[]E, E cmp.Ordered](datum S) {
	var key E
	for i, size := 1, len(datum); i < size; i++ {
		key = datum[i]
		var j int
		for j = i - 1; j >= 0 && datum[j] > key; j-- {
			datum[j+1] = datum[j]
		}
		datum[j+1] = key
	}
}
