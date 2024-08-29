package sorts

import "cmp"

func Swap[S ~[]E, E cmp.Ordered](datum S, first int, second int) {
	var temp = datum[first]
	datum[first] = datum[second]
	datum[second] = temp
}
