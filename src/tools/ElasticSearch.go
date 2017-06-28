package tools

import (
	"gopkg.in/olivere/elastic.v5"
	"os"
)

func CreateElasticsearchClient () *elastic.Client {
	//"http://somar:logitech03@localhost:9200"
	//"http://37.187.23.89:9200"
	env := os.Getenv("ENV")
	var url string

	if env == "dev" {
		url = "http://localhost:9200"
	} else {
		url = "http://37.187.23.89:9200"
	}

	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		return nil
	}

	return client
}
