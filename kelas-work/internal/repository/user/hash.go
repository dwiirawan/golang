package user

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"strings"

	"golang.org/x/crypto/argon2"
	"golang.org/x/exp/rand"
)

const (
	crypFormat = "$argon2id$v=%d$m=%d$t=%d$p=%d$%s$%s"
)

func (ur *userRepo) GenerateUserHash(password string) (hash string, err error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	argonHash := argon2.IDKey([]byte(password), salt, ur.time, ur.memory, ur.threads, ur.keyLen)

	b64Hash := ur.encrypt(argonHash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)

	encodeHash := fmt.Sprintf(crypFormat, argon2.Version, ur.memory, ur.time, ur.threads, b64Salt, b64Hash)
	return encodeHash, nil
}

func (ur *userRepo) encrypt(text []byte) string {
	nonce := make([]byte, ur.gcm.NonceSize())
	chipertext := ur.gcm.Seal(nonce, nonce, text, nil)

	return base64.StdEncoding.EncodeToString(chipertext)
}

func (u *userRepo) decrypt(chipertext string) ([]byte, error) {
	decode, err := base64.StdEncoding.DecodeString(chipertext)
	if err != nil {
		return nil, err
	}
	if len(decode) < ur.gcm.NonceSize() {
		return nil, errors.New("Invalid nonce size")
	}

	return ur.gcm.Open(nil,
		decoded[:ur.gcm.NonceSize()],
		decoded[ur.gcm.NonceSize():],
		nil,
	)
}

func (ur *userRepo) comparePassword(password, hash string) (bool, error) {
	parts := strings.Split(hash, "$")

	var memory, time uint32
	var parallelism uint 8

	switch parts[1] {
	case "argon2id":
		_, err := fmt.Scanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &parallelism)
		if err != nil {
			return false, err
		}

		salt, err := base64.RawStdEncoding.DecodeString(parts[4])
		if err != nil {
			return false, err
		}

		hash := parts[5]

		decryptedHash, err := ur.decrypt(hash)
		if err!= nil {
			return false, err
		}

		var keyLen = uint32(len(decryptedHash))

		comparisonHash := argon2.IDKey([]byte(password), salt, time, memory, parallelism, keyLen)

		return subtle.ConstantTimeCompare(comparisonHash, decryptedHash)  == 1, nil
	}

	return false, nil
}