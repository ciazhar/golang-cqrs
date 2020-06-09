package main

import (
	"github.com/ciazhar/golang-cqrs/write-service-postgres/app"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {

	//setup app
	application, err := app.SetupApp()
	if err != nil {
		panic(err)
	}

	//setup http
	if err := InitHTTP(application); err != nil {
		panic(err)
	}
}

func InitHTTP(application *app.Application) error {
	//config router api
	router := gin.New()
	social.InitHTTP(router, "/social", application)

	//middleware
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.Use(logger.SetLogger())

	//run
	log.Info().Caller().Msg("Running in port : " + application.Env.Get("port"))
	return router.Run(":" + application.Env.Get("port"))
}
