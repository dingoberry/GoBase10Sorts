package sorts

// SortRadix TimeComplex[Avg:O(n*k),Bad:O(n*k),Good:O(n*k)] SpaceComplex:O(n*k) Stable
func SortRadix(datum []int) {
	var m = 0
	for _, item := range datum {
		m = max(m, item)
	}

	size := len(datum)
	copies := make([]int, size)
	for exp := 1; m/exp > 0; exp *= 10 {
		base := make([]int, 10)

		for i := 0; i < size; i++ {
			base[(datum[i]/exp)%10]++
		}

		for i := 1; i < 10; i++ {
			base[i] += base[i-1]
		}

		for i := size - 1; i >= 0; i-- {
			index := (datum[i] / exp) % 10
			base[index]--
			copies[base[index]] = datum[i]
		}

		for i := 0; i < size; i++ {
			datum[i] = copies[i]
		}
	}
}
