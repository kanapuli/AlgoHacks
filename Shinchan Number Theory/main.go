package main

import "fmt"

func main() {
	var gates, loops, totalSectors int
	fmt.Scan(&loops)
	for i := 0; i < loops; i++ {
		fmt.Scan(&gates)
		//get the number of unique diagonals
		n := (gates * (gates - 3)) / 2
		if n == 0 {
			totalSectors = gates + 1
			fmt.Println(totalSectors)
			continue
		}
		totalSectors = (n * 2) + gates
		fmt.Println(totalSectors)
	}
}
