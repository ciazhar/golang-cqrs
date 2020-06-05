package amqp

import (
	"github.com/ciazhar/golang-cqrs/common/logger"
	"github.com/streadway/amqp"
)

type amqpClient struct {
	amqpPublisherManager
	amqpSubscriptionManager
	connection *amqp.Connection
}

func (cli *amqpClient) Init(connection *amqp.Connection) error {
	cli.connection = connection

	logger.Info("initiate publisher manager")
	if err := cli.amqpPublisherManager.Init(connection); err != nil {
		logger.Warn("Fail initiate publisher manager %v", err)
		return err
	}

	logger.Info("initiate subscription manager")
	if err := cli.amqpSubscriptionManager.Init(connection); err != nil {
		logger.Warn("Fail initiate subscription manager %v", err)
		return err
	}
	return nil
}

func (cli *amqpClient) Close() error {
	logger.Info("try close subscription manager")
	if err := cli.amqpSubscriptionManager.Close(); err != nil {
		logger.Error("Failed to close subscription manager: %v\n", err)
	}

	logger.Info("try close publisher manager")
	if err := cli.amqpPublisherManager.Close(); err != nil {
		logger.Error("Failed to close publisher manager: %v\n", err)
	}

	logger.Info("try close connection")
	return cli.connection.Close()
}
