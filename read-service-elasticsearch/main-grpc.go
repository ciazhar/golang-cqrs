package main

import (
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/app"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/pkg/social"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

func main() {
	//setup app
	application, err := app.SetupApp()
	if err != nil {
		panic(err)
	}

	//setup grpc
	if err := InitGRPC(application); err != nil {
		panic(err)
	}
}

func InitGRPC(application *app.Application) error {

	address := application.Env.Get("grpc.address")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	//init grpc server
	s := grpc.NewServer()

	//init client
	social.InitGRPC(s)

	//serve grpc server
	log.Info().Caller().Msg("Running GRPC in port : " + address)
	return s.Serve(lis)
}
