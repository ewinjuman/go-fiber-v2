package codeGenerator

import (
	"math/rand"
)

var (
	LowerCaseLettersCharset = []rune("abcdefghijklmnopqrstuvwxyz")
	UpperCaseLettersCharset = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LettersCharset          = append(LowerCaseLettersCharset, UpperCaseLettersCharset...)
	NumbersCharset          = []rune("0123456789")
	AlphanumericCharset     = append(LettersCharset, NumbersCharset...)
	SpecialCharset          = []rune("!@#$%^&*()_+-=[]{}|;':\",./<>?")
	AllCharset              = append(AlphanumericCharset, SpecialCharset...)
)

// RandomString return a random string.
func RandomString(size int, charset []rune) string {
	if size <= 0 {
		println("lo.RandomString: Size parameter must be greater than 0")
		return ""
	}
	if len(charset) <= 0 {
		println("lo.RandomString: Charset parameter must not be empty")
		return ""
	}

	b := make([]rune, size)
	possibleCharactersCount := len(charset)
	for i := range b {
		b[i] = charset[rand.Intn(possibleCharactersCount)]
	}
	return string(b)
}
