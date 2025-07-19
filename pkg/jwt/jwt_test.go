package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	// テストケースの準備
	secret := "test-secret"
	claims := map[string]any{
		"user_id": 123,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}

	// トークンの生成
	token, err := GenerateToken(claims, secret)

	// 検証
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestVerifyToken(t *testing.T) {
	// テストケースの準備
	secret := "test-secret"
	claims := map[string]any{
		"user_id": 123,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}

	// トークンの生成
	token, err := GenerateToken(claims, secret)
	assert.NoError(t, err)

	// トークンの検証
	verifiedClaims, err := VerifyToken(token, secret)

	// 検証
	assert.NoError(t, err)
	assert.Equal(t, float64(claims["user_id"].(int)), verifiedClaims["user_id"])
}

func TestVerifyTokenWithInvalidToken(t *testing.T) {
	// 無効なトークンでの検証
	invalidToken := "invalid.token.string"
	secret := "test-secret"

	// トークンの検証
	_, err := VerifyToken(invalidToken, secret)

	// 検証
	assert.Error(t, err)
}

func TestVerifyTokenWithInvalidSecret(t *testing.T) {
	// テストケースの準備
	secret := "test-secret"
	wrongSecret := "wrong-secret"
	claims := map[string]any{
		"user_id": 123,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}

	// トークンの生成
	token, err := GenerateToken(claims, secret)
	assert.NoError(t, err)

	// 間違ったシークレットでトークンの検証
	_, err = VerifyToken(token, wrongSecret)

	// 検証
	assert.Error(t, err)
}
