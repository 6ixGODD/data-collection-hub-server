package jwt__test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"
	"time"

	"data-collection-hub-server/pkg/jwt"
	"github.com/stretchr/testify/assert"
)

const (
	sub             = "test"
	tokenDuration   = 3 * time.Second
	refreshDuration = 10 * time.Second
	refreshBuffer   = 1 * time.Second
)

func TestJwt_GenerateAccessToken(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	jwtInstance := jwt.New(privateKey, tokenDuration, refreshDuration, refreshBuffer)

	accessToken, err := jwtInstance.GenerateAccessToken(sub)
	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
	t.Logf("access token: %s", accessToken)

	invalidSubject := ""
	accessToken, err = jwtInstance.GenerateAccessToken(invalidSubject)
	assert.Error(t, err)
	assert.Empty(t, accessToken)
}

func TestJwt_GenerateRefreshToken(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	jwtInstance := jwt.New(privateKey, tokenDuration, refreshDuration, refreshBuffer)

	refreshToken, err := jwtInstance.GenerateRefreshToken(sub)
	assert.NoError(t, err)
	assert.NotEmpty(t, refreshToken)
	t.Logf("refresh token: %s", refreshToken)

	invalidSubject := ""
	refreshToken, err = jwtInstance.GenerateRefreshToken(invalidSubject)
	assert.Error(t, err)
	assert.Empty(t, refreshToken)
}

func TestJwt_VerifyToken(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	jwtInstance := jwt.New(privateKey, tokenDuration, refreshDuration, refreshBuffer)

	accessToken, err := jwtInstance.GenerateAccessToken(sub)
	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
	t.Logf("access token: %s", accessToken)

	token, err := jwtInstance.VerifyToken(accessToken)
	assert.NoError(t, err)
	assert.Equal(t, sub, token)
	t.Logf("token: %s", token)

	invalidToken := "Invalid token"
	token, err = jwtInstance.VerifyToken(invalidToken)
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestJwt_RefreshToken(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	jwtInstance := jwt.New(privateKey, tokenDuration, refreshDuration, refreshBuffer)

	refreshToken, err := jwtInstance.GenerateRefreshToken(sub)
	assert.NoError(t, err)
	assert.NotEmpty(t, refreshToken)
	t.Logf("refresh token: %s", refreshToken)

	time.Sleep(tokenDuration - refreshBuffer + 1*time.Second) // wait for token to expire
	newAccessToken, err := jwtInstance.RefreshToken(refreshToken)
	assert.NoError(t, err)
	assert.NotEmpty(t, newAccessToken)
	t.Logf("new access token: %s", newAccessToken)

	invalidToken := "Invalid token"
	newAccessToken, err = jwtInstance.RefreshToken(invalidToken)
	assert.Error(t, err)
	assert.Empty(t, newAccessToken)
}

func TestJwt_ExtractClaims(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	jwtInstance := jwt.New(privateKey, tokenDuration, refreshDuration, refreshBuffer)

	subject := "test"
	accessToken, err := jwtInstance.GenerateAccessToken(subject)
	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
	t.Logf("access token: %s", accessToken)

	claims, err := jwtInstance.ExtractClaims(accessToken)
	assert.NoError(t, err)
	assert.NotEmpty(t, claims)
	t.Logf("claims: %v", claims)

	invalidToken := "Invalid token"
	claims, err = jwtInstance.ExtractClaims(invalidToken)
	assert.Error(t, err)
	assert.Empty(t, claims)
}
