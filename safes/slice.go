package safes

func Cut(v string, size int) string {
	if len(v) > size {
		return v[:size]
	}
	return v
}
