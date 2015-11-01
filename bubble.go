package main

import (
	"github.com/kogai/algorithms/helper"
)

func main() {
	helper.Benchmark(sort, 10000) // 平均値: 305.297838ms
}

func sort(arr []int) []int {
	left := 0
	max := len(arr) - 1

	for index := 0; index < max; index++ {
		for k := max; k > left; k-- {
			if helper.Compare(arr[k-1], arr[k]) > 0 {
				tmp := arr[k-1]
				arr[k-1] = arr[k]
				arr[k] = tmp
			}
		}
		left++
	}
	return arr
}
