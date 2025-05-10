package common

import (
	"github.com/nocturna-ta/ums/pkg/constants"
	"regexp"
	"strings"
)

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
