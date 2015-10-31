package main

import (
	// "fmt"
	"github.com/kogai/algorithms/helper"
)

func main() {
	helper.Benchmark(sort, 10000) // 平均値: 48.565721ms
}

func sort(a []int) []int {
	count := len(a)
	for index := 1; index < count; index++ {
		insert(a, index, a[index])
	}
	return a
}

func insert(arr []int, pos int, val int) {
	i := pos - 1
	for i >= 0 && helper.Compare(arr[i], val) > 0 {
		arr[i+1] = arr[i]
		i--
	}
	arr[i+1] = val
}
