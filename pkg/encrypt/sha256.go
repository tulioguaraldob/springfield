package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func ToHash256(toBeHashed string) string {
	hash := sha256.New()
	hash.Write([]byte(toBeHashed))

	hashed256 := hex.EncodeToString(hash.Sum(nil))

	return hashed256
}
