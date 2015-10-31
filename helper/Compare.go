package helper

func Compare(target int, base int) int {
	if base > target {
		return -1
	}
	if base < target {
		return 1
	}
	return 0
}

// 大->小を較べる
// small, big -1
// big, small 1
// even 0
