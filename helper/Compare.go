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
