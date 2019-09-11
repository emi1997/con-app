package main

import (
	// "log"
	// "context"
	"net/http"
	"github.com/olivere/elastic"
)

func main(){
	//Creating a new client for Elasticsearch instance
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
	// Handle error
	panic(err)
	}
	defer client.Stop()
}
