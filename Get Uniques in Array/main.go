package main

import "fmt"

func main(){
    //Initialize an array 
    inputArray := []int{10,20,30,56,67,90,10,20}
    printUniqueValue(inputArray)
}

func printUniqueValue( arr []int){
    //Create a   dictionary of values for each element
    var  dict map[int]int 
  
    dict =  make(map[int]int)
    for _ , num :=  range arr {
        dict[num] =dict[num] + 1
    }
    //Print The Number Of Occurances 
    fmt.Println(dict)
    //Loop through dict and  get the unique value
    for key   :=  range dict{
        if dict[key] == 1{
            fmt.Println(key)
        }
    }
}