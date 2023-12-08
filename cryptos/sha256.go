package cryptos

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Md5SumStr(v string) string {
	return Md5Sum([]byte(v))
}

func Md5Sum(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256Str(v string) string {
	return Sha256([]byte(v))
}

func Sha256(b []byte) string {
	h := sha256.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
