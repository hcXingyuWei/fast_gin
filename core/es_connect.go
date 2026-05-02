package core

import (
	"fmt"

	"github.com/olivere/elastic/v7"
)

func EsConnect() (client *elastic.Client) {

	client, err := elastic.NewClient(elastic.SetURL(
		"http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	return client
}
