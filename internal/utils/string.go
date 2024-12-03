package utils

import (
	"time"

	"golang.org/x/exp/rand"
)

var r = rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}