package mock

import (
	"fmt"
	"math/rand"
	"time"
)

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func randomEnum(enum []string) string {
	return enum[rand.Intn(len(enum))]
}

func randomTimeBeforeNow() time.Time {
	return time.Now().Add(-time.Duration(rand.Intn(1000)) * time.Hour)
}

func randomIp() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}
