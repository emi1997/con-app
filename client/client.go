package client

import (
	// "log"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	//"github.com/emi1997/con-app/scanner"

	//"net/http"
	"fmt"
	"log"
	"time"

	"github.com/olivere/elastic"
	//"github.com/spf13/cobra"
)

//Client is being called by the main.go every time the app is being started
var Client, err = elastic.NewClient(
	elastic.SetURL("http://localhost:9200"),
	elastic.SetHealthcheckInterval(10*time.Second),
)

var ctx = context.Background()


type TestEntry struct {
	Fach string `json:"Fach"`
	Stoff string `json:"Thema"`
	Wann string `json:"Wann"`
}
//var indexName = scanner.Scanner()

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

//NewMapping lets you add mapping
var NewMapping string

//AddMapping lets you add mapping
func AddMapping() {
	var mappingTemplate = `
	{
		"dynamic_templates": [
		  {
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
				"type": "text",
				"fields": {
				  "raw": {
					"type":  "keyword",
					"ignore_above": 256
				  }
				}
			  }
			}
		  }
		]
	}
	`

	NewMapping, err := Client.PutMapping().Index("school").BodyString(mappingTemplate).Do(context.TODO())
	if err != nil {
		panic(err)
		// fmt.Printf("expected put mapping to succeed; got: %v", err)
	}

	//fmt.Println(mappingTemplate, NewMapping)
	if !NewMapping.Acknowledged {
		fmt.Print("something went wrong")
	}
}

//AddDocument lets you add a document to a given index
func AddDocument() {
	//use os.Open to open a jsonfile and defer the closure of the file until the first function is finished.
	//Serves to open the json file
	jsonFile, err := os.Open("./test_data.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	fmt.Println("Opend JsonFile successfully: ", jsonFile)

	//read the json file
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []TestEntry
	
	json.Unmarshal([]byte(byteValue), &result)
	// fmt.Printf("%#v\n", result)

	bulkRequest := Client.Bulk()

	for idx, entry := range result{
		indexReq := elastic.NewBulkIndexRequest().Index("school").Type("_doc").Id(string(idx)).Doc(entry)
		bulkRequest = bulkRequest.Add(indexReq)
	}

	bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		panic(err)
	}
	indexed := bulkResponse.Indexed()
	if len(indexed) == 0{
		fmt.Println("Something went wrong!")
	}

}

//DeleteDocument lets you delete a document from a given index
func DeleteDocument() {
	res, err := Client.Delete().
		Index("school").
		Type("_doc").
		Id("1").
		Do(ctx)

	if err != nil {
		panic(err)
	}
	if res != nil {
		fmt.Print("Document deleted from from index\n")
	}
}

//GetDocument lets you read documents from an index
func GetDocument() {
	get1, err := Client.Get().
		Index("school").
		Id("3").
		Do(ctx)
	if err != nil {
		panic(err)
	}
	if get1 != nil {
		fmt.Println(get1)
	}
}

//ReadDocument lets you read documents froman index
func ReadDocument(){
	jsonFile, err := os.Open("./readQuery.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	fmt.Println("Opend JsonFile successfully: ", jsonFile)

	//read the json file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	fmt.Println(result["readQuery"])

	//get request body and post it to url
	requestBody, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("GET", "http://localhost:9200/school/_search", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	fmt.Println(req.Response)
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
