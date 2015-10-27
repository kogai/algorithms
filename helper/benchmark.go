package helper

import (
	"fmt"
	"time"
)

const maxTry int = 100

func Benchmark(sort func(a []int) []int, maxElement int) {
	var durations [maxTry]time.Duration

	for i := 0; i < maxTry; i++ {
		arr := Create(maxElement)
		start := time.Now()
		_ = sort(arr)
		end := time.Now()
		durations[i] = end.Sub(start)
		fmt.Printf("%d回目: %v\n", i+1, end.Sub(start))
	}
	r := durationAverage(durations, maxTry)
	fmt.Printf("--------------------\n平均値: %v\n", r)
}

func durationAverage(durations [maxTry]time.Duration, maxTry int) time.Duration {
	var total time.Duration
	for _, d := range durations {
		total = total + d
	}
	return time.Duration(int(total)/maxTry) * time.Nanosecond
}
