package jwt

type JWT struct {
	SecretKey []byte
}

func NewJWT(secretKey string) *JWT {
	return &JWT{[]byte(secretKey)}
}

func (j *JWT) CreateToken(claims map[string]interface{}) (string, error) {
	return "", nil
}

func (j *JWT) ParseToken(tokenString string) (map[string]interface{}, error) {
	return nil, nil
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	return "", nil
}
