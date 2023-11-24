package rands

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var (
	randGen = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func SliceInt(size int, randN int, offset int) []int {
	res := make([]int, 0, size)
	for i := 0; i < size; i++ {
		res = append(res, Intn(randN, offset))
	}
	return res
}

func Intn(randN int, offset int) int {
	return rand.Intn(randN) + offset
}

func SliceStr(size int, wordLen int) []string {
	res := make([]string, 0, size)
	for i := 0; i < size; i++ {
		res = append(res, Str(wordLen))
	}
	return res
}

func Str(length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[randGen.Intn(len(charset))]
	}
	return string(result)
}
