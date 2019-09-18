package utils

import "unicode/utf8"

func Runes(s string) (runes []string) {
	for i := 0; i < utf8.RuneCountInString(s); i++ {
		runes = append(runes, s[i:i+1])
	}
	return runes
}
