package common_utils

import "math/rand"

func GenerateRandomSixDigit() int64 {
	otp := rand.Intn(999999)
	return int64(otp)
}
