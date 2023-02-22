package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client %s\n", err)
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting the response %s\n", err)
	}
	defer res.Body.Close()
}
