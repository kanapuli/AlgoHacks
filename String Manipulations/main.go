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
}

func reverseWord(input string) (output string) {
	output = ""
	splittedChars := strings.Split(input, "")
	for i := len(splittedChars) - 1; i >= 0; i-- {
		output += splittedChars[i]
	}

	return output
}
