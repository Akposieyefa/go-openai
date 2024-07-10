package helpers

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// generate reference
func Reference(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// convert naira to kobo
func ConvertToKobo(amount int) int {
	return amount * 100
}

// convert kobo to naira
func ConvertKoboToNaira(amount int) int {
	return amount / 100
}
