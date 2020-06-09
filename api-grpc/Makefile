BINARY=engine

run:
	protoc --go_out=plugins=grpc:. ./api/social.proto

.PHONY: clean install unittest build docker run stop vendor