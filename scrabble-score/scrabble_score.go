package scrabble

import "strings"

func Score(word string) int {
	// alphabet is an array where alphabet[i] contains the value of the i-th letter of the alphabet
	alphabet := []int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}
	sum := 0
	// convert to lower case 'cause i will need to compare it with the ASCII value of 'a'
	newWord := strings.ToLower(word)
	for i := 0; i < len(word); i++ {
		// getting the ASCII value of the i-th char in the string
		asciiValue := int((newWord[i]))
		// asciiValue - int('a') --> 'b' - 'a' == 1, 'd' - 'a' == 4 where 1 and 4 are the index of the array
		sum += alphabet[asciiValue-int('a')]
	}

	return sum
}
