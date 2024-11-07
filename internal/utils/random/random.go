package random

import (
	"math/rand"
	"time"
)

func GenerateSixDigitOTP() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := 100000 + rng.Intn(900000) // 100000 - 999999
	return otp
}
