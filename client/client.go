package client

import (
	// "log"

	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	//"github.com/emi1997/con-app/scanner"

	//"net/http"
	"fmt"
	"log"
	//"time"

	"github.com/olivere/elastic"
	//"github.com/spf13/cobra"
)
//var indexName = scanner.Scanner()






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

	for idx, entry := range result {
		indexReq := elastic.NewBulkIndexRequest().Index("school").Type("_doc").Id(string(idx)).Doc(entry)
		bulkRequest = bulkRequest.Add(indexReq)
	}

	bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		panic(err)
	}
	indexed := bulkResponse.Indexed()
	if len(indexed) == 0 {
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

//UpdateDocument lets you update a specific document in a given index
func UpdateDocument() {
	jsonFile, err := os.Open("./update_data.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	fmt.Println("Opend JsonFile successfully: ", jsonFile)

	//read the json file
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var updateResult []TestUpdate

	json.Unmarshal([]byte(byteValue), &updateResult)
	//fmt.Printf("%#v\n", updateResult)

	for idx, entry := range updateResult {
		updateReq := elastic.NewBulkUpdateRequest().Index("school").Type("_doc").Id(string(idx)).Doc(entry)
		if err != nil {
			panic(err)
		}
		bulkRequest = bulkRequest.Add(updateReq)
	}
	bulkResponse, err := bulkRequest.Do(ctx)
	if err != nil {
		panic(err)
	}
	updated := bulkResponse.Updated()
	if len(updated) == 0 {
		fmt.Println("Something went wrong!")
	}
}

//UpdateMapping lets you update the mapping of a given index
func UpdateMapping() {
	var mappingTemplate = `
	{
		"dynamic_date_formats": ["MM/dd/yyyy", "E, d M y H:m:s"],
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
