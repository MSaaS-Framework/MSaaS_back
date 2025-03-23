package security

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

// HashPassword는 평문 비밀번호를 해싱하고, 해시와 랜덤 Salt를 반환합니다.
func HashPassword(password string) (hashedPassword string, salt string, err error) {
	// 랜덤 Salt 생성
	saltBytes := make([]byte, 16)
	_, err = rand.Read(saltBytes)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate salt: %w", err)
	}

	// Argon2로 해싱
	hash := argon2.IDKey([]byte(password), saltBytes, 1, 64*1024, 4, 32)

	// Base64로 인코딩
	hashedPassword = base64.RawStdEncoding.EncodeToString(hash)
	salt = base64.RawStdEncoding.EncodeToString(saltBytes)
	return hashedPassword, salt, nil
}

// VerifyPassword는 입력된 비밀번호가 저장된 해시와 일치하는지 검증합니다.
func VerifyPassword(inputPassword, storedHash, storedSalt string) (bool, error) {
	// Base64로 인코딩된 Salt 디코딩
	saltBytes, err := base64.RawStdEncoding.DecodeString(storedSalt)
	if err != nil {
		return false, fmt.Errorf("failed to decode salt: %w", err)
	}

	// Argon2로 해시 재계산
	hash := argon2.IDKey([]byte(inputPassword), saltBytes, 1, 64*1024, 4, 32)
	computedHash := base64.RawStdEncoding.EncodeToString(hash)

	// 해시 비교
	return storedHash == computedHash, nil
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	return salt, err
}
