package elastic

import (
	"github.com/olivere/elastic/v7"
)

func InitElastic() (*elastic.Client, error) {
	client, err := elastic.NewClient()
	if err != nil {
		return client, err
	}
	return client, nil
}
