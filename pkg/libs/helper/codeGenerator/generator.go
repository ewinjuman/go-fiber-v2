package codeGenerator

import (
	"crypto/rand"
	"math/big"
)

var (
	LowerCaseLettersCharset      = []rune("abcdefghijklmnopqrstuvwxyz")
	UpperCaseLettersCharset      = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LettersCharset               = append(LowerCaseLettersCharset, UpperCaseLettersCharset...)
	NumbersCharset               = []rune("0123456789")
	AlphanumericCharset          = append(LettersCharset, NumbersCharset...)
	SpecialCharset               = []rune("!@#$%^&*()_+-=[]{}|;':\",./<>?")
	AllCharset                   = append(AlphanumericCharset, SpecialCharset...)
	AlphanumericLowerCaseCharset = append(LowerCaseLettersCharset, NumbersCharset...)
	AlphanumericUpperCaseCharset = append(UpperCaseLettersCharset, NumbersCharset...)
)

// RandomString return a random string.
func RandomString(size int, charset []rune) string {
	if size <= 0 {
		println("Size parameter must be greater than 0")
		return ""
	}
	if len(charset) <= 0 {
		println("Charset parameter must not be empty")
		return ""
	}

	b := make([]rune, size)
	possibleCharactersCount := big.NewInt(int64(len(charset)))
	for i := range b {
		randomNumber, err := rand.Int(rand.Reader, possibleCharactersCount)
		if err != nil {
			println("charset error")
			return ""
		}
		b[i] = charset[randomNumber.Int64()]
	}
	return string(b)
}
