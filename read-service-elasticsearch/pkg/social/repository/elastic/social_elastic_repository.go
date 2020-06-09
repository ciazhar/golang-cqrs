package elastic

import (
	"context"
	"fmt"
	"github.com/ciazhar/golang-cqrs/common/rest"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/app"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/pkg/social/model"
	"github.com/olivere/elastic/v7"
	uuid "github.com/satori/go.uuid"
)

type SocialPostgresRepository interface {
	Fetch(param rest.Param) ([]model.Social, error)
	GetByID(id string) (model.Social, error)
	GetByName(name string) (model.Social, error)
	Store(req *model.Social) error
	Update(req *model.Social) error
}

type repository struct {
	IndexName   string
	Application *app.Application
}

func (r repository) Fetch(param rest.Param) ([]model.Social, error) {
	socials := make([]model.Social, 0)

	return socials, nil
}

func (r repository) GetByID(id string) (model.Social, error) {
	social := model.Social{Id: id}
	response, err := r.Application.Elastic.Get().Index(r.IndexName).Id(id).Do(context.Background())
	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			panic(fmt.Sprintf("Document not found: %v", err))
		case elastic.IsTimeout(err):
			panic(fmt.Sprintf("Timeout retrieving document: %v", err))
		case elastic.IsConnErr(err):
			panic(fmt.Sprintf("Connection problem: %v", err))
		default:
			// Some other kind of error
			panic(err)
		}
	}
	fmt.Print(response)
	return social, nil
}

func (r repository) GetByName(name string) (model.Social, error) {
	social := model.Social{Name: name}

	return social, nil
}

func (r repository) Store(req *model.Social) error {
	id := uuid.Must(uuid.NewV4(), nil)
	req.Id = id.String()
	response, err := r.Application.Elastic.Index().
		Index(r.IndexName).
		Id(id.String()).
		BodyJson(req).
		Do(context.Background())

	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}

func (r repository) Update(req *model.Social) error {
	return nil
}

func NewSocialElasticRepository(application *app.Application) SocialPostgresRepository {
	indexName := "social"
	exists, err := application.Elastic.IndexExists(indexName).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !exists {
		// Create a new index.
		mapping := `
		{
			"settings":{
				"number_of_shards":1,
				"number_of_replicas":1
			},
			"mappings":{
				"properties":{
					"name":{
						"type":"keyword"
					},
					"detail":{
						"type":"text"
					},
					"created_at":{
						"type":"date"
					},
					"updated_at":{
						"type":"date"
					},
					"deleted_at":{
						"type":"date"
					}
				}
			}
		}
		`
		createIndex, err := application.Elastic.CreateIndex(indexName).Body(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	return repository{
		IndexName: indexName,
	}
}
