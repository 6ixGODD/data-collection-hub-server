package jwt

type JWT struct {
	GenerateToken func(subject string) (string, error)
}
