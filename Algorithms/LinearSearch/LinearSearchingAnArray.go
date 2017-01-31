package main

import "fmt"

func searchAnElement(array [5]int, term int) {
	//Linear Searching is starting from the first element and comparing each element to the  right. If the  number equals then return
	//the index of the array .
	//If nothing matches return -1
	for index, num := range array {
		if num == term {
			fmt.Println("The number ", term, " is in the index ", index)
		}
	}
}

func main() {

	arr := [5]int{1, 3, 4, 5, 6}
	searchAnElement(arr, 6)
	//The time complexity for linear search is O(N)  since it iterates N times
}
