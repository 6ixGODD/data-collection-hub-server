package jwt

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtInstance Jwt

type Jwt interface {
	GenerateToken(subject string) (string, error)
	VerifyToken(token string) (string, error)
	RefreshToken(token string) (string, error)
	ExtractClaims(token string) (map[string]interface{}, error)
}

type jwtImpl struct {
	privateKey      *ecdsa.PrivateKey
	publicKey       *ecdsa.PublicKey
	tokenDuration   time.Duration
	refreshDuration time.Duration
	refreshBuffer   time.Duration
}

func New(
	privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, tokenDuration, refreshDuration time.Duration,
	refreshBuffer time.Duration,
) Jwt {
	if jwtInstance != nil {
		return jwtInstance
	}
	var _ Jwt = (*jwtImpl)(nil) // Ensure jwtImpl implements Jwt
	jwtInstance = &jwtImpl{
		privateKey:      privateKey,
		publicKey:       publicKey,
		tokenDuration:   tokenDuration,
		refreshDuration: refreshDuration,
		refreshBuffer:   refreshBuffer,
	}
	return jwtInstance
}

func (a *jwtImpl) GenerateToken(subject string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodES256, &jwt.StandardClaims{
			Subject:   subject,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(a.tokenDuration).Unix(),
			NotBefore: time.Now().Unix(),
		},
	)

	tokenString, err := token.SignedString(a.privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *jwtImpl) RefreshToken(token string) (string, error) {
	claims, err := a.ExtractClaims(token)
	if err != nil {
		return "", err
	}

	// Check if token is expired
	if time.Unix(int64(claims["exp"].(float64)), 0).Sub(time.Now()) > a.refreshBuffer {
		return "", fmt.Errorf("token is not expired yet")
	}

	// Generate new token
	newToken, err := a.GenerateToken(claims["sub"].(string))
	if err != nil {
		return "", err
	}

	return newToken, nil
}

func (a *jwtImpl) ExtractClaims(token string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}

	t, err := jwt.ParseWithClaims(
		token, claims, func(token *jwt.Token) (interface{}, error) {
			return a.publicKey, nil
		},
	)
	if err != nil || !t.Valid {
		return nil, err
	}

	return claims, nil
}

func (a *jwtImpl) VerifyToken(token string) (string, error) {
	claims, err := a.ExtractClaims(token)
	if err != nil {
		return "", err
	}

	return claims["sub"].(string), nil
}
