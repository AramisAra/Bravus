package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// This generate a random slice of bytes for encoding use.
// The numbers of bytes is define by
// params: n - For the numbers of bytes
func generateRandomBytes(numbersOfbytes int) ([]byte, error) {
	byteSlice := make([]byte, numbersOfbytes)
	_, err := rand.Read(byteSlice)
	if err != nil {
		return nil, err
	}
	return byteSlice, nil
}

// This function hashes the password of the user
func HashPassword(password string) (string, error) {
	salt, err := generateRandomBytes(32)
	if err != nil {
		return "", nil
	}

	hash := argon2.IDKey([]byte(password), salt, 4, 32*1024, 1, 16)
	encodedHash := fmt.Sprintf("%s$%s", base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash))

	return encodedHash, nil
}

// This Compares the passwords of the user.
// It use the hash stored in the database and the give password.
func ComparePasswords(hashedPassword, plainPassword string) error {
	parts := strings.Split(hashedPassword, "$")
	if len(parts) != 2 {
		return errors.New("invalid hash format")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return err
	}

	hash := argon2.IDKey([]byte(plainPassword), salt, 4, 32*1024, 1, 16)
	expectedHash := parts[1]
	actualHash := base64.RawStdEncoding.EncodeToString(hash)

	if actualHash != expectedHash {
		return errors.New("not password")
	}

	return nil
}
