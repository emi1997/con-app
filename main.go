package main

import (
	// "log"
	// "context"
	"github.com/olivere/elastic"
)

func main(){
	client, err := elastic.NewClient()
	if err != nil {
	// Handle error
	panic(err)
	}
	defer client.Stop()
}
