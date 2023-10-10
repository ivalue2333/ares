package ares

import "math/rand"

func RandSlice(size int, randN int, randOffset int) []int {
	res := make([]int, 0, size)
	for i := 0; i < size; i++ {
		res = append(res, rand.Intn(randN)+randOffset)
	}
	return res
}
