package client


import (
	"fmt"
)
//AddIndex lets you add a new index
func AddIndex() {

	//indexName := scanner.Scanner()
	createIndex, err := Client.CreateIndex("school").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !createIndex.Acknowledged {
		fmt.Println(err)
	} else {
		fmt.Printf("Index created!")
	}
	//return indexName
}