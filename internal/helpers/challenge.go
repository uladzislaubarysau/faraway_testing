package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

const zeroCount = 5

var zeros string

func init() {
	zeros = strings.Repeat("0", zeroCount)
}

// To check pow of the clients, the hash calculation method was chosen by the sha256 algorithm,
// with the condition that the hash must start with a certain number of zeros

func Solve(challenge []byte) []byte {
	for i := 0; ; i++ {
		prefix := []byte(strconv.Itoa(i))
		if Verify(challenge, prefix) {
			return prefix
		}
	}
}

func Verify(challenge, prefix []byte) bool {
	hash := sha256.New()
	hash.Write(prefix)
	hash.Write(challenge)

	encodedHash := hex.EncodeToString(hash.Sum(nil))

	return strings.HasPrefix(encodedHash, zeros)
}
