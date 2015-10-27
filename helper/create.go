package helper

import (
	// "fmt"
	"math/rand"
	"time"
)

func Create(max int) []int {
	var sortable []int = make([]int, max)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < max; i++ {
		val := rand.Intn(max)
		sortable[i] = val
	}
	// fmt.Println(sortable)
	return sortable
}
