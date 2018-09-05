package experiments

import (
	"unicode"
)

const (
	alphaLen = 26
)

// mod returns the modulo of a and b
// and is augumented to account for a negative value of a
func mod(a, b int) int {
	i := a % b
	if i < 0 {
		return b + i
	}
	return i
}

// numToAlpha maps a given number to the corresponding
// alphabet in the universe of alphabet (a-z)
func numToAlpha(i int) (string, bool) {
	dict := map[int]string{
		1: "A", 2: "B", 3: "C", 4: "D", 5: "E", 6: "F", 7: "G",
		8: "H", 9: "I", 10: "J", 11: "K", 12: "L", 13: "M", 14: "N",
		15: "O", 16: "P", 17: "Q", 18: "R", 19: "S", 20: "T",
		21: "U", 22: "V", 23: "W", 24: "X", 25: "Y", 26: "Z",
	}
	l, ok := dict[i]
	return l, ok
}

// alphaToNum maps a given string to the corresponding number
// from 1 - 26 of the alphabets
func alphaToNum(a string) (int, bool) {
	dict := map[string]int{
		"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7,
		"H": 8, "I": 9, "J": 10, "K": 11, "L": 12, "M": 13, "N": 14,
		"O": 15, "P": 16, "Q": 17, "R": 18, "S": 19, "T": 20, "U": 21,
		"V": 22, "W": 23, "X": 24, "Y": 25, "Z": 26,
	}
	k := unicode.ToUpper(rune(a[0]))
	l, ok := dict[string(k)]
	return l, ok
}

// forward returns an alphabet k steps after the alphabet a
// if there is no value for a in the dictionary it returns
// a unmodified
func forward(a string, k int) string {
	num, ok := alphaToNum(a)
	if !ok {
		return a
	}
	m := mod((num + k), alphaLen)
	alpha, ok := numToAlpha(m)
	if !ok {
		return a
	}
	if unicode.IsUpper(int32(a[0])) {
		return string(unicode.ToUpper(rune(alpha[0])))
	}
	return string(unicode.ToLower(rune(alpha[0])))
}

// backward returns an alphabet k steps after the alphabet a
// if there is no value for a in the dictionary it returns a unmodified
func backward(a string, k int) string {
	num, ok := alphaToNum(a)
	if !ok {
		return a
	}
	m := mod((num - k), alphaLen)
	alpha, ok := numToAlpha(m)
	if !ok {
		return a
	}
	if unicode.IsUpper(int32(a[0])) {
		return string(unicode.ToUpper(rune(alpha[0])))
	}
	return string(unicode.ToLower(rune(alpha[0])))
}

// Cipher returns the ciphertext for the given string
// using the ceaser method of substitution
func Cipher(alpha string, step int) string {
	ct := ""
	for _, a := range alpha {
		ct += forward(string(a), step)
	}
	return ct
}

// Decipher returns the plaintext for the given string
// using the ceaser method of substitution
func Decipher(alpha string, step int) string {
	ct := ""
	for _, a := range alpha {
		ct += backward(string(a), step)
	}
	return ct
}
