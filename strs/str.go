package strs

import "strconv"

// Str like python str.
func Str[T int | int8 | int16 | int32 | int64](v T) string {
	return strconv.FormatInt(int64(v), 10)
}

func Strs[T int | int8 | int16 | int32 | int64](vv []T) []string {
	res := make([]string, 0, len(vv))
	for _, v := range vv {
		s := Str(v)
		res = append(res, s)
	}
	return res
}
