package sorts

import "cmp"

// SortQuick TimeComplex[Avg:O(nLogn),Bad:O(n2),Good:O(nLogn)] SpaceComplex:O(Logn) Unstable
func SortQuick[S ~[]E, E cmp.Ordered](datum S) {
	SortQuickPartition(datum, 0, len(datum)-1)
}

func SortQuickPartitionTransform[S ~[]E, E cmp.Ordered](datum S, low int, high int) {
	SortQuickPartition(datum, low, high)
}

func SortQuickPartition[S ~[]E, E cmp.Ordered](datum S, low int, high int) {
	var i, j = low, high
	if i >= j {
		return
	}

	var temp = datum[low]
	for i != j {
		for datum[j] >= temp && i < j {
			j--
		}
		for datum[i] <= temp && i < j {
			i++
		}
		if i < j && (i != low || j != high) {
			Swap(datum, i, j)
		}
	}

	if i != low {
		Swap(datum, i, low)
	}

	SortQuickPartitionTransform(datum, low, i-1)
	SortQuickPartitionTransform(datum, i+1, high)
}
