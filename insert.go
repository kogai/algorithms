package main

import (
	"github.com/kogai/algorithms/helper"
)

func main() {
	helper.Benchmark(sort, 10000) // 平均値: 165.659119ms
}

func sort(arr []int) []int {
	max := len(arr)

	for i := 1; i < max; i++ {
		for index := 0; index < i; index++ {
			if helper.Compare(arr[index], arr[i]) > 0 {
				tmp := arr[index]
				arr[index] = arr[i]
				arr[i] = tmp
			}
		}
	}

	return arr
}
