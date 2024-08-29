package sorts

// SortCounting TimeComplex[Avg:O(n+k),Bad:O(n+k),Good:O(n+k)] SpaceComplex:O(k) Stable
func SortCounting(datum []int) {
	var m = 0
	for _, item := range datum {
		m = max(m, item)
	}

	countingSpace := make([]int, m+1)
	for _, item := range datum {
		countingSpace[item]++
	}

	newIndex := 0
	for index, item := range countingSpace {
		for ; item > 0; item-- {
			datum[newIndex] = index
			newIndex++
		}
	}
}
