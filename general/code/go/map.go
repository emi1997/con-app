package main

import (
	"fmt"
)

func main(){

//implement hash tables
//map[KeyType]ValueType
//Initialization
m := make(map[string]int)
m["route"] = 66
//retrieve the value stored in a key:
//i := m["route"]
i := m["root"] //=> i == 0 because key doesn´t exist
n := len(m) //=> returns number of elements in a map
delete(m, "route") // deletes an entry from a map
_, ok := m["route"]// How to test for a key, without retrieving the value
m["route"] = 66
commits := map [string]int{ //=> initializing map with some data using literal
	"rsc": 3711,
	"r": 2138,
	"gri": 1908,
	"adg": 912,
}
for key, value := range commits{ //=> iterates over contents of map with keyword range
	fmt.Println("Key: ", key, " Value: ", value)
}
fmt.Println(m)
fmt.Println(i)
fmt.Println(n)
fmt.Println(ok)
}