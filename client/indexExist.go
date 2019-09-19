package client

import (
	"fmt"
)

//IndexExist checks if an index is already there
func IndexExist() {
	//indexName := scanner.Scanner()
	indexExists, err := Client.IndexExists("school").Do(ctx)
	if indexExists {
		fmt.Println("That one exists already!")
	}
	if err != nil {
		fmt.Println(err)
	}
}
