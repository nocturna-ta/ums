package utils

import (
	"crypto/sha512"
	"fmt"
	"github.com/nocturna-ta/golib/utils/encryption"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/pkg/constants"
	"log"
	"regexp"
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

func Encryption(text string) string {
	cfg := &config.MainConfig{}
	config.ReadConfig(cfg, "")

	enc, err := encryption.NewEncryption(cfg.Encryption.Key)
	if err != nil {
		log.Fatal(err)
	}

	res, err := enc.Encrypt(text)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func Decryption(text string) string {
	cfg := &config.MainConfig{}
	config.ReadConfig(cfg, "")

	enc, err := encryption.NewEncryption(cfg.Encryption.Key)
	if err != nil {
		log.Fatal(err)
	}

	res, err := enc.Decrypt(text)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func IsValidNIK(nik string) bool {
	nik = strings.TrimSpace(nik)
	if nik == constants.EmptyString {
		return false
	}

	if len(nik) != 16 {
		return false
	}

	matched, err := regexp.MatchString(`^\d{16}$`, nik)
	if err != nil {
		return false
	}

	return matched
}
