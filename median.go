package main

import (
	"fmt"
	"github.com/kogai/algorithms/helper"
)

func main() {
	// arr := helper.Create(16)
	// 要素数[16] 中央値[8]
	// 最初の軸値[7] 最初の軸要素[12]
	arr := []int{15, 9, 8, 1, 4, 11, 7, 12, 13, 6, 5, 3, 16, 2, 10, 14}
	fmt.Println("課題配列", arr)

	var _ []int = medianSort(arr, 0, len(arr)-1)
	// fmt.Println(s)
	// helper.Benchmark(medianSort, 10000)
}

func medianSort(arr []int, left, right int) []int {
	if left < right {
		var mid int = (left + right) / 2
		var me int = selectKth(arr, mid, left, right)
		// var _ int = partition(arr, 0, 15, 9)
		// var _ int = partition(arr, left, right, mid)
		fmt.Println(me)
	}
	return arr
}

func selectKth(arr []int, k, left, right int) int {
	fmt.Println("中央idx:", k)
	var pivotIndex int = partition(arr, left, right, k)
	if left+k-1 == pivotIndex {
		return pivotIndex
	}
	if left+k-1 < pivotIndex {
		return selectKth(arr, k, left, pivotIndex-1)
	} else {
		return selectKth(arr, k-(pivotIndex-left+1), pivotIndex+1, right)
	}
	return pivotIndex
}

func partition(arr []int, left, right, pivotIdx int) int {
	var idx, store int
	var pivot int = arr[pivotIdx]

	// fmt.Printf("pivotIdx: %d\n", pivotIdx)
	// fmt.Printf("pivot: %d\n", pivot)

	var tmp int = arr[right]
	arr[right] = arr[pivotIdx]
	arr[pivotIdx] = tmp
	// fmt.Println(arr)

	store = left
	for idx = left; idx < right; idx++ {
		if helper.Compare(arr[idx], pivot) <= 0 {
			// fmt.Println(arr[idx], pivot, helper.Compare(arr[idx], pivot))
			tmp = arr[idx]
			arr[store] = tmp
			// fmt.Println(arr)
			store++
		}
	}

	tmp = arr[right]
	arr[right] = arr[store]
	arr[store] = tmp

	// fmt.Println(arr)
	// fmt.Println(store)
	return store
}

/*
// -1 = Compare(small, big)
// +1 = Compare(big, small)

func findMedian(arr []int) int {
	var arrCount = len(arr)
}
*/
