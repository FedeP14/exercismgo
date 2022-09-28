package isogram

import "strings"

func IsIsogram(word string) bool {
	var alphabet [27]int
	newWord := strings.ToLower(word)

	for i := 0; i < len(word); i++ {
		if word[i] == '-' || word[i] == '_' || word[i] == ' ' {
			alphabet[26]++
		} else {
			asciiValue := int(newWord[i])
			alphabet[asciiValue-int('a')]++

			if alphabet[asciiValue-int('a')] == 2 {
				return false
			}

		}
	}
	return true
}
