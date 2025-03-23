package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHashPasswordAndVerify 테스트는 비밀번호 해싱과 검증 과정을 테스트합니다.
func TestHashPasswordAndVerify(t *testing.T) {
	password := "securepassword"

	// 비밀번호 해싱
	hashedPassword, salt, err := HashPassword(password)
	assert.NoError(t, err, "HashPassword should not return an error")
	assert.NotEmpty(t, hashedPassword, "Hashed password should not be empty")
	assert.NotEmpty(t, salt, "Salt should not be empty")

	// 올바른 비밀번호 검증
	isValid, err := VerifyPassword(password, hashedPassword, salt)
	assert.NoError(t, err, "VerifyPassword should not return an error for a valid password")
	assert.True(t, isValid, "Password verification should be successful for a valid password")

	// 잘못된 비밀번호 검증
	isValid, err = VerifyPassword("wrongpassword", hashedPassword, salt)
	assert.NoError(t, err, "VerifyPassword should not return an error for an invalid password")
	assert.False(t, isValid, "Password verification should fail for an invalid password")
}

// TestHashPasswordUniqueness 테스트는 서로 다른 salt로 해싱된 비밀번호가 다른 해시를 생성하는지 확인합니다.
func TestHashPasswordUniqueness(t *testing.T) {
	password := "securepassword"

	// 첫 번째 해싱
	hashedPassword1, salt1, err := HashPassword(password)
	assert.NoError(t, err, "First HashPassword should not return an error")

	// 두 번째 해싱
	hashedPassword2, salt2, err := HashPassword(password)
	assert.NoError(t, err, "Second HashPassword should not return an error")

	// 서로 다른 Salt와 해시 값 확인
	assert.NotEqual(t, salt1, salt2, "Each generated salt should be unique")
	assert.NotEqual(t, hashedPassword1, hashedPassword2, "Each hashed password should be unique")
}
