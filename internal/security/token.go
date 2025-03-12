package security

import (
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Access Token 설정: 15분 유효
const AccessTokenExpiry = 15 * time.Minute

// Refresh Token 설정: 7일 유효
const RefreshTokenExpiry = 7 * 24 * time.Hour

// MakeAccessToken 생성 (짧은 유효기간)
func MakeAccessToken(user *ent.User) (string, error) {
	return CreateToken(user.ID.String(), user.Role.String(), AccessTokenExpiry)
}

// MakeRefreshToken 생성 (긴 유효기간)
func MakeRefreshToken(user *ent.User) (string, error) {
	return CreateToken(user.ID.String(), user.Role.String(), RefreshTokenExpiry)
}

// CreateToken 생성 공통 로직
func CreateToken(userID string, role string, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"role":   role,
		"iat":    time.Now().Unix(),             // 발급 시간
		"exp":    time.Now().Add(expiry).Unix(), // 만료 시간
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 비밀 키 확인
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET is not set")
	}

	// 서명 후 토큰 반환
	return token.SignedString([]byte(secret))
}

// VerifyToken: 토큰 검증 로직 추가
func VerifyToken(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET is not set")
	}

	// 파싱 및 검증 로직
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// ExtractClaims: 토큰에서 사용자 정보 추출
func ExtractClaims(tokenString string) (map[string]interface{}, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
