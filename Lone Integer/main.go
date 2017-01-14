package main

import "fmt"

func main() {

	var test int
	fmt.Scan(&test)
	loop(test)

}

func loop(times int) {
	count := 0
	mySlice := make([]int, 0, 1)
	newSlice := make([]int, 0, 1)
	for i := 0; i < times; i++ {
		var arrayLength int
		fmt.Scan(&arrayLength)
		array := make([]int, arrayLength)
		for z := 0; z < arrayLength; z++ {
			fmt.Scan(&array[z])
		}
		duplicateArray := make([]int, arrayLength)
		copy(duplicateArray, array)

		for i := 0; i < len(array); i++ {
			temp := array[i]
			for a, b := range array {
				if i != a && a >= i && temp == b {
					count++
					newSlice = append(mySlice, temp)
					fmt.Println(count)
					fmt.Println(newSlice)
					pop
				}
			}
		}

	}
}
