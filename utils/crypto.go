package utils

import (
	"crypto"
	"encoding/hex"
	"hash"
)

func getHasher() hash.Hash {
	encryptMethod := GetConfig().Get("security.encryption")
	switch encryptMethod {
	case "md5":
		return crypto.MD5.New()
	case "sha256":
		return crypto.SHA256.New()
	}
	return nil
}

// EncryptString encrypts strings such as password
func EncryptString(toEncrypt string) string {
	hasher := getHasher()
	hasher.Write([]byte(toEncrypt))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}
