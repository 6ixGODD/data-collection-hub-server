package jwt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// TODO: Add logger

type Jwt interface {
	GenerateAccessToken(subject string) (string, error)
	GenerateRefreshToken(subject string) (string, error)
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
	privateKey *ecdsa.PrivateKey, tokenDuration, refreshDuration time.Duration,
	refreshBuffer time.Duration,
) Jwt {
	j := &jwtImpl{
		privateKey:      privateKey,
		publicKey:       &privateKey.PublicKey,
		tokenDuration:   tokenDuration,
		refreshDuration: refreshDuration,
		refreshBuffer:   refreshBuffer,
	}
	if err := j.checkJWT(); err != nil {
		return nil
	}
	return j
}

func (j *jwtImpl) checkJWT() error {
	if j.privateKey == nil {
		_privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			return err
		}
		j.privateKey = _privateKey
		j.publicKey = &j.privateKey.PublicKey
	}
	if j.tokenDuration == 0 {
		j.tokenDuration = time.Hour
	}
	if j.refreshDuration == 0 {
		j.refreshDuration = 24 * time.Hour
	}
	if j.refreshBuffer == 0 {
		j.refreshBuffer = time.Minute
	}
	if j.tokenDuration > j.refreshDuration || j.refreshDuration < j.refreshBuffer || j.tokenDuration < j.refreshBuffer {
		j.tokenDuration = time.Hour
		j.refreshDuration = 24 * time.Hour
		j.refreshBuffer = time.Minute
	}

	return nil
}

func (j *jwtImpl) GenerateAccessToken(subject string) (string, error) {
	if subject == "" {
		return "", fmt.Errorf("subject is empty") // TODO: CHANGE ERROR TYPE
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodES256, &jwt.StandardClaims{
			Subject:   subject,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(j.tokenDuration).Unix(),
			NotBefore: time.Now().Unix(),
		},
	)

	tokenString, err := token.SignedString(j.privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwtImpl) GenerateRefreshToken(subject string) (string, error) {
	if subject == "" {
		return "", fmt.Errorf("subject is empty") // TODO: CHANGE ERROR TYPE
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodES256, &jwt.StandardClaims{
			Subject:   subject,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(j.tokenDuration).Unix(),
			NotBefore: time.Now().Unix(),
		},
	)

	tokenString, err := token.SignedString(j.privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwtImpl) RefreshToken(token string) (string, error) {
	claims, err := j.ExtractClaims(token)
	if err != nil {
		return "", err
	}

	// Check if token is expired
	if time.Unix(int64(claims["exp"].(float64)), 0).Sub(time.Now()) > j.refreshBuffer {
		return "", fmt.Errorf(
			"token is not expired yet: %v", time.Unix(int64(claims["exp"].(float64)), 0).Sub(time.Now()),
		)
	}

	// Generate new token
	newToken, err := j.GenerateAccessToken(claims["sub"].(string))
	if err != nil {
		return "", err
	}

	return newToken, nil
}

func (j *jwtImpl) ExtractClaims(token string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}

	t, err := jwt.ParseWithClaims(
		token, claims, func(token *jwt.Token) (interface{}, error) {
			return j.publicKey, nil
		},
	)
	if err != nil || !t.Valid {
		return nil, err
	}

	return claims, nil
}

func (j *jwtImpl) VerifyToken(token string) (string, error) {
	claims, err := j.ExtractClaims(token)
	if err != nil {
		return "", err
	}

	return claims["sub"].(string), nil
}
