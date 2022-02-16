package common

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func ranSequence(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(9999)%len(letters)]
	}
	return string(b)
}

func GenSalt(length int) string {
	if length < 0 {
		length = 50
	}
	return ranSequence(length)
}
