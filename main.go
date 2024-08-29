package main

import (
	"awesomeProject/sorts"
	"cmp"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

const seedSortA = 10
const baseSortA = 10000
const seedSortB = 1000
const baseSortB = 5000000
const bucketAmount = 100000

type Pair struct {
	data  any
	count int
}

func sort[S ~[]E, E cmp.Ordered](name string, datum *S, impl func(datum *S)) S {
	copies := make(S, len(*datum))
	copy(copies, *datum)
	start := time.Now()
	impl(&copies)
	println("Sort[", name, "] Spent", time.Since(start).Milliseconds(), "MS")
	return copies
}

func main() {
	var goods []string
	for i, length := 0, rand.Intn(seedSortA)+baseSortA; i < length; i++ {
		if str, err := uuid.NewRandom(); err == nil {
			goods = append(goods, str.String())
		}
	}

	var sortList [][]string
	sortList = append(sortList, sort("Bubble", &goods, func(datum *[]string) {
		sorts.SortBubble(*datum)
	}))
	sortList = append(sortList, sort("Merge", &goods, func(datum *[]string) {
		sorts.SortMerge(*datum)
	}))
	sortList = append(sortList, sort("Select", &goods, func(datum *[]string) {
		sorts.SortSelect(*datum)
	}))
	sortList = append(sortList, sort("Insert", &goods, func(datum *[]string) {
		sorts.SortInsert(*datum)
	}))
	sortList = append(sortList, sort("Shell", &goods, func(datum *[]string) {
		sorts.SortShell(*datum)
	}))
	sortList = append(sortList, sort("Heap", &goods, func(datum *[]string) {
		sorts.SortHeap(*datum)
	}))

	var sortVerify []Pair
outer:
	for _, item := range sortList {
		if len(sortVerify) == 0 {
			sortVerify = make([]Pair, len(item))
			for index, sd := range item {
				sortVerify[index].data = sd
			}
		}

		end := len(sortVerify)
		for index, sd := range item {
			if index < end && sortVerify[index].data == sd {
				sortVerify[index].count++
			} else {
				break outer
			}
		}
	}

	var compareCount = -1
	var compareFail = false
	for _, item := range sortVerify {
		if compareCount == -1 {
			compareCount = item.count
		} else if item.count != compareCount {
			compareFail = true
			break
		}
	}

	if compareFail {
		println("Not matched sort data list")
	} else {
		println("Matched sort data list")
	}

	println("__________________________________________________________")

	var numbers []int
	for i, length := 0, rand.Intn(seedSortB)+baseSortB; i < length; i++ {
		numbers = append(numbers, rand.Intn(1000000))
	}

	var sortNumberList [][]int
	sortNumberList = append(sortNumberList, sort("Counting", &numbers, func(datum *[]int) {
		sorts.SortCounting(*datum)
	}))
	sortNumberList = append(sortNumberList, sort("Bucket", &numbers, func(datum *[]int) {
		sorts.SortBucket(*datum, bucketAmount)
	}))
	sortNumberList = append(sortNumberList, sort("Radix", &numbers, func(datum *[]int) {
		sorts.SortRadix(*datum)
	}))

	var sortCountVerify []Pair
outerC:
	for _, item := range sortNumberList {
		if len(sortCountVerify) == 0 {
			sortCountVerify = make([]Pair, len(item))
			for index, sd := range item {
				sortCountVerify[index].data = sd
			}
		}

		end := len(sortCountVerify)
		for index, sd := range item {
			if index < end && sortCountVerify[index].data == sd {
				sortCountVerify[index].count++
			} else {
				break outerC
			}
		}
	}

	compareCount = -1
	compareFail = false
	for _, item := range sortCountVerify {
		if compareCount == -1 {
			compareCount = item.count
		} else if item.count != compareCount {
			compareFail = true
			break
		}
	}

	if compareFail {
		println("Not matched sort number list")
	} else {
		println("Matched sort number list")
	}
}
