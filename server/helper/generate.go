package helper

import "github.com/google/uuid"

func RandomString(n int) string {
	id := uuid.NewString()
	return id[:n]
}
