package main

import (
	// "fmt"
	"github.com/kogai/algorithms/helper"
)

func main() {
	helper.Benchmark(sort, 10000) // 平均値: 205.251865ms
}

func sort(arr []int) []int {
	var max int = len(arr)
	for i := 0; i < max; i++ {
		for k := i; k < max; k++ {
			if helper.Compare(arr[i], arr[k]) > 0 {
				var tmp int = arr[i]
				arr[i] = arr[k]
				arr[k] = tmp
			}
		}
	}
	return arr
}
