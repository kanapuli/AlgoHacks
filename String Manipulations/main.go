package main

import (
	"fmt"
	"strings"
)

//Reverse a C style string.  C style string means 'fegh' has five character spaces including the null character

func main() {
	var word string
	fmt.Println("Enter the word to reverse")
	fmt.Scanf("%s", &word)
	reversed := reverseWord(word)
	fmt.Println(reversed)
	uniqueString := removeDuplicateLiterals(word)
	fmt.Println(uniqueString)
}

func reverseWord(input string) (output string) {
	output = ""
	splittedChars := strings.Split(input, "")
	for i := len(splittedChars) - 1; i >= 0; i-- {
		output += splittedChars[i]
	}

	return output
}
func removeDuplicateLiterals(input string) (output string) {
	dicts := make(map[string]int)
	for _, character := range strings.Split(input, "") {
		dicts[character]++
		if dicts[character] > 1 {
			continue
		}
		output += character
	}
	return output
}
