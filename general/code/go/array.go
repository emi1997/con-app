package main

import (
	"fmt"
)

func main(){

// Declaring an Array:
// The length of an array is static and given in side the parantheses. Otherwise similar to the declaration of variables.
// They are indexed in the usual way, starting with 0.
var a [4]int
a[0] = 1
i := a[0]

// there is another way to initialize arrays:
// let the compiler do the counting of the elements
b := [...]string{"Penn", "Teller"}
fmt.Println(i, b)
}