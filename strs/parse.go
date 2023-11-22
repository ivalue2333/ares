package strs

import (
	"fmt"
	"strconv"
)

func Parse(v string) (int64, error) {
	return strconv.ParseInt(v, 10, 64)
}

func ParseTwo(v1, v2 string) (r1 int64, r2 int64, err error) {
	r1, err = Parse(v1)
	if err != nil {
		err = fmt.Errorf("v1 parse %w", err)
		return
	}
	r2, err = Parse(v2)
	if err != nil {
		err = fmt.Errorf("v2 parse %w", err)
		return
	}
	return
}

func ParseThree(v1, v2, v3 string) (r1 int64, r2 int64, r3 int64, err error) {
	r1, err = Parse(v1)
	if err != nil {
		err = fmt.Errorf("v1 parse %w", err)
		return
	}
	r2, err = Parse(v2)
	if err != nil {
		err = fmt.Errorf("v2 parse %w", err)
		return
	}
	r3, err = Parse(v3)
	if err != nil {
		err = fmt.Errorf("v3 parse %w", err)
		return
	}
	return
}
