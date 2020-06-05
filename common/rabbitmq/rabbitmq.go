package rabbitmq

import (
	"github.com/ciazhar/golang-cqrs/common/amqp"
	"github.com/ciazhar/golang-cqrs/common/env"
)

func InitRabbitMQBroker(environment *env.Environtment) (*amqp.AmqpBroker, error) {
	amqpConf := amqp.RabbitConfig{
		Host:     environment.Get("rabbitmq.host"),
		User:     environment.Get("rabbitmq.username"),
		Password: environment.Get("rabbitmq.passoword"),
	}
	broker := amqp.NewAmqpBroker(&amqpConf)
	err := broker.Start()
	return broker, err
}
