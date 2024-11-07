package utils

import "fmt"

func GetUserKey(hashKey string) string {
	return fmt.Sprintf("u:%s:otp", hashKey)
}
