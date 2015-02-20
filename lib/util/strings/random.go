package strings

import (
	"crypto/rand"
	"math/big"
)

const (
	Alphabet     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	Numerals     = "1234567890"
	Alphanumeric = Alphabet + Numerals
)

func Random(n int) string {
	charset := Alphanumeric
	randstr := make([]byte, n) // Random string to return
	charlen := big.NewInt(int64(len(charset)))
	for i := 0; i < n; i++ {
		b, _ := rand.Int(rand.Reader, charlen)
		// if err != nil {
		// 	return "", err
		// }
		r := int(b.Int64())
		randstr[i] = charset[r]
	}
	return string(randstr)
}
