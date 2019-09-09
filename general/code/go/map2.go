package main

import (
	"fmt"
)

func main(){
type Node struct {
	Next *Node
	Value interface{}
}
var first *Node

visited := make(map[*Node]bool)
for n := first; n != nil; n = n.Next {
	if visited[n] {
		fmt.Println("cycle detected")
		break
	}
	visited[n] = true
	fmt.Println(n.Value)
}

	fmt.Println()
}