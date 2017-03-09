package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the Length of the Array")
	var arrLength int
	fmt.Scanf("%d", &arrLength)
	fmt.Println("Enter the Array Elements")
	//Arrays cannot be initialized with dynamic length  . Hence  declared a  slice
	var arr = make([]int, arrLength)
	for i := 0; i < arrLength; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Enter the Number to be searched")
	var x int
	fmt.Scanf("%d", &x)
	binarySearch(arr, x)
}

func binarySearch(arr []int, x int) {
	//get the midlength of the slice
	mid := len(arr) / 2
	if arr[mid] == x {
		fmt.Println("The Index of the number ", arr[mid], " is ", mid)
		return
	} else if x < arr[mid] {
		//the value of x is less than arr[mid] .Hence it should be in the left half of the array
		binarySearch(arr[0:mid], x)
	} else {
		fmt.Println(arr)
		//the value x is greater  than the arr[mid]. hence it is the right half of the array
		binarySearch(arr[mid:], x)
	}

}
