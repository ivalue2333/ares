package rands

func Intn(n int) int {
	return randGen.Intn(n)
}

func Ints(n int, offset int) int {
	return randGen.Intn(n) + offset
}

func Intn64(n int) int64 {
	return int64(randGen.Intn(n))
}
