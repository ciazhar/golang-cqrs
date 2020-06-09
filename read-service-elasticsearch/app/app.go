package app

import (
	"github.com/ciazhar/golang-cqrs/common/amqp"
	"github.com/ciazhar/golang-cqrs/common/elastic"
	"github.com/ciazhar/golang-cqrs/common/env"
	"github.com/ciazhar/golang-cqrs/common/logger"
	"github.com/ciazhar/golang-cqrs/common/rabbitmq"
	"github.com/ciazhar/golang-cqrs/common/validator"
	"github.com/gin-gonic/gin"
	elastic2 "github.com/olivere/elastic/v7"
	"os"
)

type Application struct {
	Env            *env.Environtment
	Elastic        *elastic2.Client
	RabbitMQBroker *amqp.AmqpBroker
}

func SetupApp() (*Application, error) {

	//env
	environment := env.InitEnv()

	//set default timezone
	if err := os.Setenv("TZ", "Asia/Jakarta"); err != nil {
		panic(err.Error())
	}

	//profile
	gin.SetMode(environment.Get("profile"))

	//logger
	logger.InitLogger()

	//elastic
	elasticClient, err := elastic.InitElastic()
	if err != nil {
		return &Application{}, err
	}

	//rabbitmq
	rabbitMQBroker, err := rabbitmq.InitRabbitMQBroker(environment)
	if err != nil {
		return &Application{}, err
	}

	//validator
	validator.Init()

	return &Application{
		Env:            environment,
		Elastic:        elasticClient,
		RabbitMQBroker: rabbitMQBroker,
	}, nil
}

func SetupAppWithPath(path string) (*Application, error) {
	environment := env.InitPath(path)
	logger.InitLogger()
	//elastic
	elasticClient, err := elastic.InitElastic()
	if err != nil {
		return &Application{}, err
	}
	rabbitMQBroker, err := rabbitmq.InitRabbitMQBroker(environment)
	if err != nil {
		return &Application{}, err
	}

	validator.Init()

	return &Application{
		Env:            environment,
		Elastic:        elasticClient,
		RabbitMQBroker: rabbitMQBroker,
	}, nil
}
