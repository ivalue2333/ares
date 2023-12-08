package rands

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Word 后期优化.
func Word() string {
	length := Intn(len(charset))
	return Str(length)
}

func Str(length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[randGen.Intn(len(charset))]
	}
	return string(result)
}
