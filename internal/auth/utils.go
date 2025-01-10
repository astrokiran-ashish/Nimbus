package auth

import (
	"math/rand"
)

func (auth *Auth) generateRandomSixDigit() int64 {
	otp := rand.Intn(999999)
	return int64(otp)
}
