package user

import (
	"encoding/base64"
	"fmt"

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
