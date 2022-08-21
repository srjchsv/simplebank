package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}
func RandomInt(min, max int64) int64 {
	return min + int64(rand.Int63n(max-min+1))
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
