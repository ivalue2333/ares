package strs

import (
	"strconv"

	"github.com/ivalue2333/ares/internal/constraints"
)

// Str like python str.
func Str[T constraints.Integer](v T) string {
	if v < 0 {
		return strconv.FormatInt(int64(v), 10)
	}
	return strconv.FormatUint(uint64(v), 10)
}

func StrF[T constraints.Float](v T) string {
	return strconv.FormatFloat(float64(v), 'g', -1, 64)
}

func Strs[T constraints.Integer](vv []T) []string {
	res := make([]string, 0, len(vv))
	for _, v := range vv {
		s := Str(v)
		res = append(res, s)
	}
	return res
}

func StrFs[T constraints.Float](vv []T) []string {
	res := make([]string, 0, len(vv))
	for _, v := range vv {
		s := StrF(v)
		res = append(res, s)
	}
	return res
}
