package elastic

import (
	"github.com/ciazhar/golang-cqrs/common"
	"github.com/ciazhar/golang-cqrs/common/env"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/app"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/pkg/social/model"
	"testing"
	"time"
)

var Application *app.Application

func init() {
	application, err := app.SetupAppWithPath(env.GetEnvPath() + "/write-service-postgres/config.json")
	if err != nil {
		panic(err)
	}
	Application = application
}

var ID string

func NewActual() model.Social {
	var social model.Social
	common.ToStruct("social/actual.1.golden", &social)
	social.CreatedAt = time.Now()
	social.UpdatedAt = time.Now()
	return social
}

func NewActual2() model.Social {
	var social model.Social
	common.ToStruct("social/actual.2.golden", &social)
	social.CreatedAt = time.Now()
	social.UpdatedAt = time.Now()
	return social
}

func TestRepositoryFetch(t *testing.T) {

}

func TestRepositoryGetByID(t *testing.T) {

}

func TestRepositoryGetByName(t *testing.T) {

}
