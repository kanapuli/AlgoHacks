package main

import "fmt"

func main() {
	//Jump Search

	var length int
	fmt.Scanf("%d", &length)
	var a = make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &a[i])
	}
	var numberToSearch int
	fmt.Println("Enter the number to search")

	fmt.Scanf("%d", &numberToSearch)
	index := jumpSearch(a, numberToSearch)
	fmt.Println(index)

}

func jumpSearch(arr []int, num int) int {
	step := 3
	for i := 0; i < len(arr); i += step {
		fmt.Println(arr[i])
	}
	return step
}
