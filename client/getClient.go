package client

import (
	"fmt"
	"github.com/olivere/elastic"
	"log"
)
/*
GetNewClient will throw error message if there is any, during the
call of NewClient
*/
func GetNewClient() {
	//Creating a new client for Elasticsearch instance
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("Success!")
	if elastic.IsConnErr(err) {
		log.Fatalf("Elasticsearch connection problem: %v", err)
	}
}