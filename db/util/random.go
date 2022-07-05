package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // min -> max
}

func RandomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func RandomOwner() string {
	return RandomString(10)
}

func RandomBalance() int64 {
	return RandomInt(100, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "JPY", "VND", "CAD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
