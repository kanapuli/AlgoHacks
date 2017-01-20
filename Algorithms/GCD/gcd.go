package gcd

import "fmt"

//FindGCD - finds the greatest common divisor between two numbers
func FindGCD(a, b int) (gcd int) {
	//create a list
	var list []int
	for i := 1; i <= smallestNumber(a, b); i++ {
		//check if i divides both a and b . if yes then add the number to the list
		if a%i == 0 && b%i == 0 {
			list = append(list, i)
		}
	}
	fmt.Println(list)
	//return the last element of the list
	return list[len(list)-1]

}

func smallestNumber(a, b int) (c int) {
	if a < b {
		return a
	}
	return b
}
