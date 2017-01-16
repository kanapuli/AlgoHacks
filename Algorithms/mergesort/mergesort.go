package mergesort

import "fmt"

//Split : Splits an array into two equal halves
func Split(input []int) []int {
	if len(input) == 1 {
		return input
	}
	fmt.Println("Input Array : ", input)
	//The second index length is not inclusive , means input[0:4] gives array elements 0,1,2,3
	firstHalf := input[0 : len(input)/2]
	secondHalf := input[len(input)/2:]
	firstHalf = Split(firstHalf)
	secondHalf = Split(secondHalf)
	mergedArray := Merge(firstHalf, secondHalf)
	fmt.Println("Merged Array : ", mergedArray)
	return mergedArray
}

//Merge : Merges two array . The order sorted in descending order by default
func Merge(left, right []int) []int {
	var mergedArray []int
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				mergedArray = append(mergedArray, left[0])
				left = left[1:]
			} else {
				mergedArray = append(mergedArray, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			mergedArray = append(mergedArray, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			mergedArray = append(mergedArray, right[0])
			right = right[1:]
		}
	}
	return mergedArray

}
