package social

import (
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/app"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/pkg/social/controller/rest"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/pkg/social/repository/elastic"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/pkg/social/usecase"
	"github.com/gin-gonic/gin"
)

func InitHTTP(engine *gin.Engine, routes string, application *app.Application) {
	repo := elastic.NewSocialElasticRepository(application)
	uc := usecase.NewSocialUseCase(repo)
	controller := rest.NewSocialController(uc)

	r := engine.Group(routes)
	{
		r.GET("/", controller.Fetch)
		r.GET("/:id", controller.GetByID)
	}
}
