// Package strings is a helper library written in Go for work with numbers strings.
package strings

import (
	"errors"
	"strings"
)

// Flip outputs the reverse string of a given string s
func Flip(s string) string {
	strLen := len(s)
	// The reverse of a empty string is a empty string
	if strLen == 0 || strLen == 1 {
		return s
	}
	// Convert s into unicode points
	r := []rune(s)
	// Last index
	rLen := len(r) - 1
	// String new home
	rev := []string{}
	for i := rLen; i >= 0; i-- {
		rev = append(rev, string(r[i]))
	}
	return strings.Join(rev, "")
}

// PigLatin translator
// "egg" -> "eggway"
// "pig" -> "igay"
func PigLatin(s string) (string, error) {
	if s == "" {
		return "", errors.New("cannot translate an empty string")
	}
	word := strings.ToLower(s)
	switch word[0] {
	case 'a', 'e', 'i', 'o', 'u':
		return word + "way", nil
	default:
		return word[1:] + string(word[0]) + "ay", nil
	}
}
