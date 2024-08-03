package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

func generateRandomBytes(n int) ([]byte, error) {
	slice := make([]byte, n)
	_, err := rand.Read(slice)
	if err != nil {
		return nil, err
	}
	return slice, nil
}

func HashPassword(password string) (string, error) {
	salt, err := generateRandomBytes(32)
	if err != nil {
		return "", nil
	}

	hash := argon2.IDKey([]byte(password), salt, 2, 64*1024, 2, 64)
	encodedHash := fmt.Sprintf("%s$%s", base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash))

	fmt.Print(encodedHash)
	return encodedHash, nil
}

// This Compares the passwords
func ComparePasswords(hashedPassword, plainPassword string) error {
	parts := strings.Split(hashedPassword, "$")
	if len(parts) != 2 {
		return errors.New("invalid hash format")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return err
	}

	hash := argon2.IDKey([]byte(plainPassword), salt, 2, 64*1024, 2, 64)
	expectedHash := parts[1]
	actualHash := base64.RawStdEncoding.EncodeToString(hash)

	if actualHash != expectedHash {
		return errors.New("not password")
	}

	return nil
}
