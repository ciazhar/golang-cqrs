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

	//setup
	application, err := app.SetupApp()
	if err != nil {
		panic(err)
	}

	//config router api
	router := gin.New()
	social.New(router, "/social", application)

	//middleware
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.Use(logger.SetLogger())

	//run
	log.Info().Caller().Msg("Running in port : " + application.Env.Get("port"))
	if err := router.Run(":" + application.Env.Get("port")); err != nil {
		panic(err.Error())
	}
}
