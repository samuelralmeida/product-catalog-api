package hash

import (
	"crypto/sha256"
	"encoding/base64"
)

func Sha256(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}
