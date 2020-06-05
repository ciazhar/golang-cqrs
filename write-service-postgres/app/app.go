package app

import (
	"github.com/ciazhar/golang-cqrs/common/amqp"
	pg2 "github.com/ciazhar/golang-cqrs/common/db/pg"
	"github.com/ciazhar/golang-cqrs/common/env"
	"github.com/ciazhar/golang-cqrs/common/logger"
	"github.com/ciazhar/golang-cqrs/common/rabbitmq"
	"github.com/ciazhar/golang-cqrs/common/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"os"
)

type Application struct {
	Env            *env.Environtment
	Postgres       *pg.DB
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

	//postgres
	pgConn := pg2.InitPG(environment)

	//rabbitmq
	rabbitMQBroker, err := rabbitmq.InitRabbitMQBroker(environment)
	if err != nil {
		return &Application{}, err
	}

	//validator
	validator.Init()

	return &Application{
		Env:            environment,
		Postgres:       pgConn,
		RabbitMQBroker: rabbitMQBroker,
	}, nil
}

func SetupAppWithPath(path string) (*Application, error) {
	environment := env.InitPath(path)
	logger.InitLogger()
	pgConn := pg2.InitPG(environment)
	rabbitMQBroker, err := rabbitmq.InitRabbitMQBroker(environment)
	if err != nil {
		return &Application{}, err
	}

	validator.Init()

	return &Application{
		Env:            environment,
		Postgres:       pgConn,
		RabbitMQBroker: rabbitMQBroker,
	}, nil
}
