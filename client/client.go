package client

import (
	// "log"
	"context"
	"github.com/emi1997/con-app/scanner"

	//"net/http"
	"fmt"
	"log"
	"time"

	"github.com/olivere/elastic"
	//"github.com/spf13/cobra"
)

//NewClient is being called by the main.go every time the app is being started
var NewClient, err = elastic.NewClient(
	elastic.SetURL("http://localhost:9200"),
	elastic.SetHealthcheckInterval(10*time.Second),
)
var Ctx = context.Background()

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

//AddIndex lets you add a new index
func AddIndex() {

	indexName := scanner.Scanner()
	createIndex, err := NewClient.CreateIndex(indexName).BodyString(NewMapping).Do(Ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !createIndex.Acknowledged {
		fmt.Println(err)
	} else {
		fmt.Printf("Index created!")
	}
}

//IndexExist checks if an index is already there
func IndexExist() {
	indexName := scanner.Scanner()
	indexExists, err := NewClient.IndexExists(indexName).Do(Ctx)
	if indexExists {
		fmt.Println("That one exists already!")
	}
	if err != nil {
		fmt.Println(err)
	}
}

//DeleteIndex lets you delete an Index
func DeleteIndex() {
	indexName := scanner.Scanner()
	deleteIndex, err := NewClient.DeleteIndex(indexName).Do(Ctx)
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

var NewMapping string

//AddMapping lets you add mapping
func AddMapping() *elastic.IndicesPutTemplateResponse {
	var mappingTemplate = `{
		"settings":{
			"number_of_shards":2,
			"number_of_replicas":1
		},
		"dynamic_template":[
			"integers": {
				"match_mapping_type": "long",
				"mapping": {
				  "type": "integer"
				}
			  }
			},
			{
			"strings": {
				"match_mapping_type": "string",
				"mapping": {
				  "type": "text"
				}
			}
		]
	}`
	NewMapping, err := NewClient.IndexPutTemplate(mappingTemplate).Do(Ctx)
	if err != nil {
		fmt.Println(err)
	}
	return NewMapping
}

//NewMapping lets you add mapping

//AddDocument lets you add a document to a given index
func AddDocument() {

}

//DeleteDocument lets you delete a document from a given index
func DeleteDocument() {

}

//ReadDocument lets you read documents froman index
func ReadDocument() {

}

//UpdateDocument lets you update a specific document in a given index
func UpdateDocument() {

}

//UpdateMapping lets you update the mapping of a given index
func UpdateMapping() {

}

//BulkIndex let´s you index multiple documents at once
func BulkIndex() {

}
