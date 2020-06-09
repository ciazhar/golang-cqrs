package social

import (
	"github.com/ciazhar/golang-cqrs/write-service-postgres/app"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/controller/rest"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/repository/postgres"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/usecase"
	postgres2 "github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/validator/postgres"
	"github.com/gin-gonic/gin"
)

func InitHTTP(engine *gin.Engine, routes string, app *app.Application) {
	repo := postgres.NewSocialPostgresRepository(app)
	uc := usecase.NewSocialUseCase(app, repo)
	controller := rest.NewSocialController(uc)
	postgres2.NewSocialPostgresValidator(repo)

	r := engine.Group(routes)
	{
		r.POST("/", controller.Store)
		r.PUT("/", controller.Update)
		r.DELETE("/:id", controller.Delete)
	}
}
