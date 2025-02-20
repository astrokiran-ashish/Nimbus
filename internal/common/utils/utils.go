package common_utils

import (
	"encoding/json"
	"math/rand"
)

func GenerateRandomSixDigit() int64 {
	otp := rand.Intn(999999)
	return int64(otp)
}

func MapToJsonString(data map[string]string) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}
