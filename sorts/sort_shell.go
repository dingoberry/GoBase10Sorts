package sorts

import "cmp"

// SortShell TimeComplex[Avg:O(nLogn),Bad:O(nLog2n),Good:O(nLog2n)] SpaceComplex:O(1) Unstable
func SortShell[S ~[]E, E cmp.Ordered](datum S) {
	size := len(datum)
	for gap := size / 2; gap > 0; gap /= 2 {
		for i := gap; i < size; i++ {
			var j = i
			var tmp = datum[j]
			for ; j-gap >= 0 && tmp < datum[j-gap]; j -= gap {
				datum[j] = datum[j-gap]
			}
			datum[j] = tmp
		}
	}
}
