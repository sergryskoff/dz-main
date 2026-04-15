// package  реализует функцию генерации hash и salt
package util

import (
	"crypto/sha256"
	"math/rand"
)

// Генерируем hash по значению и salt
func MakeHash(value, salt string) [32]byte {
	s := append([]byte(value), []byte(salt)...)
	return sha256.Sum256(s)
}

func MakeSalt(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
