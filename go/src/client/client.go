package client

import (
	// "log"
	// "context"
	//"net/http"
	"time"
	"log"
	"fmt"

	"github.com/olivere/elastic"
	//"github.com/spf13/cobra"
)

func GetNewClient(){
	//Creating a new client for Elasticsearch instance
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
	)
	if err != nil {
	// Handle error
	panic(err)
	}
	fmt.Println("Success!")
	if elastic.IsConnErr(err) {
		log.Fatalf("Elasticsearch connection problem: %v", err)
	}
	defer client.Stop()
	return 
}

func AddIndex(){

}

func DeleteIndex(){

}

func AddMapping(){

}

func AddDocument(){

}

func DeleteDocument(){

}

func ReadDocument(){

}

func UpdateDocument(){

}

func UpdateMapping(){

}

