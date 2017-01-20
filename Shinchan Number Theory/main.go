package main

import "fmt"

func main() {
	var gates, loops, totalSectors int
	fmt.Scan(&loops)
	for i := 0; i < loops; i++ {
		fmt.Scan(&gates)
		//get the number of unique diagonals
		//Use Euler Formula
		//Faces =  Edges - Vertices + 2
		//Vertices can be calculated by the number of gates plus the combination of intersection points
		vertices := gates +
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


func factorial (n uint64) (result uint64){
	
}
