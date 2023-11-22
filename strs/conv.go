package strs

import "strconv"

func Str2Int(data string) int {
	res, _ := strconv.ParseInt(data, 10, 64)
	return int(res)
}

func Str2Int64(data string) int64 {
	res, _ := strconv.ParseInt(data, 10, 64)
	return res
}

func Str2Ints(dataList []string) []int {
	res := make([]int, 0, len(dataList))
	for _, data := range dataList {
		d, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return nil
		}
		res = append(res, int(d))
	}
	return res
}

func Str2Int64s(dataList []string) []int64 {
	res := make([]int64, 0, len(dataList))
	for _, data := range dataList {
		d, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return nil
		}
		res = append(res, d)
	}
	return res
}
