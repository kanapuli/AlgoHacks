package main

import "fmt"

//read me file is the same as the Lone Integer
func main() {
	var loop, arrayLength, exoredValue int
	fmt.Scan(&loop)
	for i := 0; i < loop; i++ {
		fmt.Scan(&arrayLength)
		inputArray := make([]int, arrayLength)
		for j := 0; j < arrayLength; j++ {
			fmt.Scan(&inputArray[j])
		}
		//Perform a  bitwise operation in each array element.
		//so a ^ a = 0  ; a^0 = a
		//There will be only one lone integer .So this formula works
		exoredValue = inputArray[0]
		for index, value := range inputArray {
			if index > 0 {
				exoredValue = exoredValue ^ value
			}
		}
		if exoredValue == 0 {
			fmt.Println(-1)
			continue
		}
		fmt.Println(exoredValue)
	}
}
