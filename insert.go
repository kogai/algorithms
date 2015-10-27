package main

import (
	"fmt"
)

func main() {
	var arr []int = []int{15, 9, 8, 1, 4, 11, 7, 12, 13, 6, 5, 3, 16, 2, 10, 14}
	fmt.Println(arr)
	s := sort(arr)
	fmt.Println(s)
}

func sort(a []int) []int {
	count := len(a)
	for index := 1; index < count; index++ {
		insert(a, index, a[index])
	}
	return a
}

func insert(arr []int, pos int, val int) {
	fmt.Println("-------start")
	fmt.Println(arr)
	// 1, 9
	// 2, 8
	i := pos - 1
	for i >= 0 && compare(arr[i], val) > 0 {
		arr[i+1] = arr[i]
		fmt.Println(arr)
		i--
	}
	arr[i+1] = val
	fmt.Println(arr)
	fmt.Println("-------end")
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
