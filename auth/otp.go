package auth

import (
	"math"
	"math/rand"
	"time"
)

func GenerateOTPCode(length int) int {
	min := int(math.Pow10(length - 1))
	max := int(math.Pow10(length))
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
