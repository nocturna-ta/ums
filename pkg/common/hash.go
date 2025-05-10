package common

import (
	"crypto/sha512"
	"fmt"
	"strings"
)

func CustomHash(str ...string) string {
	hash := sha512.New()
	hash.Write([]byte(strings.Join(str, "#")))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func PasswordHash(password, salt string) string {
	return CustomHash(password, salt)
}
