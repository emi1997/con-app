package main

import (
	"fmt"
)

func main(){

	//Declaration:
	letters := []string{"a", "b", "c", "d"}
	//Type: []string
	//No specified length
	//Can be made with the make function:
	var s[]byte = make([]byte, 5, 5)
	//make allocates an array, and returns a slice referring to that array
	//Arguments: 1. type, 2. length, 3. capacity (optional)
	length := len(s) == 5
	//inspecting the length of a slice
	capacity := cap(s) == 5
	//inspecting the capacity of a slice
	b := []string{"g", "o", "l", "a", "n", "g"}
	var slice []string = b[:2]
	//b[:2] == []string{"g", "o"}
	fmt.Println(letters, length, capacity, b[1:4], slice)
}
