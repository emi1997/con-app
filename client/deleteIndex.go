package client

import (
	"fmt"
	"log"
)

//DeleteIndex lets you delete an Index
func DeleteIndex() {
	//indexName := scanner.Scanner()
	deleteIndex, err := Client.DeleteIndex("school").Do(ctx)
	if err != nil {
		// Handle error
		log.Fatalf("An Error has happend. Most likely your index wasn´t found. This is the error message: %v", err)
	}
	if !deleteIndex.Acknowledged {
		panic(err)
	} else {
		fmt.Println("Success!")
	}
}
