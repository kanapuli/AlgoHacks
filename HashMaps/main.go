package main

import (
	"fmt"
	"strings"
)

/* Hash Maps can solve a large amount of programming puzzles
In Golang Hashmaps are called maps
The syntax to implement a map is map[KeyType]ValueType
*/

func main() {
	//dict := make(map[string]int) //This is the syntax to initialize a map .
	/* var dict map[string]int  is the syntax .  Map is a reference type , hence the variable dict is a empty map . Writing to a
		empty map gives runtime panic , hence to initialize a  map 'make' is used .
	  Maps can also be initialized during declaration  like dict := make(map[string]int){
	  "Apples": 5,
	  "oranges": 90
	}
	*/
	// dict["Apples"] = 5
	// dict["Oranges"] = 19
	// dict["Grapes"] = 89
	// for k, v := range dict {
	// 	fmt.Printf("%s : %d \n", k, v)
	// }
	var word string
	fmt.Println("Enter the Word")
	fmt.Scanf("%s", &word)
	isUnique := checkUniqueChars(word)
	fmt.Println(isUnique)

}

//Algorithm to Implement a String has only unique characters
func checkUniqueChars(input string) (isUnique bool) {
	values := strings.Split(input, "")
	uniqueKey := make(map[string]int)
	for _, i := range values {
		uniqueKey[i]++
		if uniqueKey[i] > 1 {
			//fmt.Println(uniqueKey)
			return false
		}
	}

	return true
}
