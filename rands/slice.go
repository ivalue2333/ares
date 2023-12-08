package rands

func SliceInt(size int, n int, offset int) []int {
	res := make([]int, 0, size)
	for i := 0; i < size; i++ {
		res = append(res, Ints(n, offset))
	}
	return res
}

func SliceStr(size int, wordLen int) []string {
	res := make([]string, 0, size)
	for i := 0; i < size; i++ {
		res = append(res, Str(wordLen))
	}
	return res
}
