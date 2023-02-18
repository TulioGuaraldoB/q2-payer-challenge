package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashToSHA256(toBeHashed string) string {
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(toBeHashed))
	hashed256 := hex.EncodeToString(sha256Hash.Sum(nil))

	return hashed256
}
