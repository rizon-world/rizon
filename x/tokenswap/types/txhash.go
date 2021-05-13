package types

import (
	"regexp"
)

const (
	// 32 byte tx hash's hexadecimal string length
	TxHashStringLength = 64
	// lower case hexadecimal typed string format
	TxHashRegExp = "[a-f0-9]{64}$"
)

// ValidTxHash returns whether tx hash is valid or not
// tx hash should be 64-length, lower case, hexadecimal type string
func ValidTxHash(hash string) bool {
	// check simply string length first
	if len(hash) != TxHashStringLength {
		return false
	}

	// check the format of tx hash string
	matched, _ := regexp.MatchString(TxHashRegExp, hash)
	if !matched {
		return false
	}

	return true
}
