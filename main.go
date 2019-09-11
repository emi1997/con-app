package main

import (
	"log"
	"github.com/elastic/go-elasticsearch/v7"
)

func main(){
	cfg := elasticsearch.Config{
		Addresses: []string{
		  "http://localhost:9200",
		  "http://localhost:9201",
		},
	  }
	elastic, _ := elasticsearch.NewClient(cfg)
	log.Println(elasticsearch.Version)
	log.Println(elastic.Info())
}
