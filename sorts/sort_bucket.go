package sorts

import (
	"math"
)

type Bucket struct {
	next *Bucket
	data int
}

// SortBucket TimeComplex[Avg:O(n+k),Bad:O(n2),Good:O(n+k)] SpaceComplex:O(n+k) Stable
func SortBucket(datum []int, amount int) {
	var m = 0
	for _, item := range datum {
		m = max(m, item)
	}

	rangeScope := int(math.Ceil(float64(m)/float64(amount))) + 1
	buckets := make([]*Bucket, amount)
	for _, item := range datum {
		bucketIndex := item / rangeScope
		header := buckets[bucketIndex]
		bucket := new(Bucket)
		bucket.data = item
		if header == nil {
			buckets[bucketIndex] = bucket
		} else {
			for header.next != nil && header.next.data < item {
				header = header.next
			}

			if header == buckets[bucketIndex] && header.data > item {
				bucket.next = header
				buckets[bucketIndex] = bucket
			} else {
				tmp := header.next
				header.next = bucket
				bucket.next = tmp
			}
		}
	}

	index := 0
	for _, item := range buckets {
		if item == nil {
			continue
		}

		for ; item != nil; item = item.next {
			datum[index] = item.data
			index++
		}
	}
}
