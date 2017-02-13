package main

import "fmt"

func main() {
	//Prepare a slice here
	inputArray := []int{2, 90, 87, 56, 4, 3, 5, 6, 7}
	sortedArray := insertionSort(inputArray)
	fmt.Println(sortedArray)
}

func insertionSort(arr []int) []int  {
    for index , value := range arr{
        if index >= 1 {
            j :=  value 
            k := index
            // fmt.Println("Index values:", k,j)
            for k > 0 {
                if j <  arr[k-1]{
                    arr[k] , arr[k-1] = arr[k-1] ,arr[k] 
                }else{
                    break
                }
                k--
            }
        }
    }
    return arr
}
