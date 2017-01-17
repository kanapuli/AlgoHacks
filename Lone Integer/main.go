package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var test int
	fmt.Scan(&test)
	loop(test)
	fmt.Println("Press any key  to  quit")
	end := make([]byte, 10)
	if _, err := os.Stdin.Read(end); err != nil {
		log.Fatal(err)
	}
}

func loop(times int) {

	for i := 0; i < times; i++ {
		arrayLength := 0
		positiveCount := 0
		fmt.Scan(&arrayLength)
		arrayMap := make(map[int]int)
		array := make([]int, arrayLength)
		var output string
		for j := 0; j < arrayLength; j++ {
			fmt.Scan(&array[j])
			arrayMap[array[j]]++
		}
		for key, value := range arrayMap {
			if value == 1 {
				output = output + " " + strconv.Itoa(key)
				positiveCount++
				continue
			}

		}
		if positiveCount == 0 {
			fmt.Println(-1)
		} else {
			fmt.Println(output)
		}

		end := make([]byte, 10)
		if _, err := os.Stdin.Read(end); err != nil {
			log.Fatal(err)
		}
	}

}
