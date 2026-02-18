package main

import "strings"

func FirstNonRepeating(str string) string {
	charMap := make(map[rune]int)
	lstr := strings.ToLower(str)

	for _, char := range lstr {
		charMap[char] += 1
	}

	for i, char := range lstr {
		if charMap[char] == 1 {
			return string(str[i])
		}
	}
	return ""
}
