package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var totalNumbers int
	var temp string
	fmt.Scanf("%d", &totalNumbers)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		temp = scanner.Text()
		break
	}
	array := strings.Split(temp, " ")
	array = array[0:totalNumbers]
	for i := totalNumbers - 1; i >= 0; i-- {
		fmt.Printf("%v ", array[i])
	}
}
