package main

import (
	// "fmt"
	"github.com/kogai/algorithms/helper"
)

func main() {
	// helper.Benchmark(sort, 60)
	helper.Benchmark(sort, 10000)
}

func sort(a []int) []int {
	count := len(a)
	for index := 1; index < count; index++ {
		insert(a, index, a[index])
	}
	// fmt.Println(a)
	return a
}

func insert(arr []int, pos int, val int) {
	// fmt.Println("-------start")
	// fmt.Println(arr)
	// 1, 9
	// 2, 8
	i := pos - 1
	for i >= 0 && compare(arr[i], val) > 0 {
		arr[i+1] = arr[i]
		// fmt.Println(arr)
		i--
	}
	arr[i+1] = val
	// fmt.Println(arr)
	// fmt.Println("-------end")
}

func compare(target int, base int) int {
	if base > target {
		return -1
	}
	if base < target {
		return 1
	}
	return 0
}
