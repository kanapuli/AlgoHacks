package main

import (
	"fmt"

	ms "github.com/Athavankanapuli/algorithms/Algorithms/mergesort"
)

func main() {
	//Merge sort
	arr := []int{2, 56, 4, 103, 78, 34, 84, 5, 2}
	mergedArray := ms.Split(arr)
	fmt.Println("In The main function. The Merged Array is : ", mergedArray)
}
