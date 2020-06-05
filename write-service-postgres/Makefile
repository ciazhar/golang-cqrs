BINARY=engine
MOCK=./domain

run:
	go run main.go

test:
	go test ./...

coverage:
	go test -coverprofile cp.out ./...

coverage-html:
	go tool cover -html=cp.out

mock:
	mockery -dir=pkg -all -output=./internal/mocks

lint:
	golangci-lint run -D=typecheck
	gosec -exclude=G304 ./...
	revive ./...
	gocritic check ./...
	go-consistent -v ./...

readme:
	npx readme-md-generator

docker:
	docker build -f build.dockerfile -t user-dev . && docker run -d user-dev -p 8080:8080 -v $PWD:/go/src/github.com/ciazhar/digisar-service

migrate:
	goose -dir=migrations postgres "user=postgres dbname=orm_test sslmode=disable" up

migrate-dev:
	goose -dir=migrations postgres "user=developer password=developer+1DB dbname=testingdb port=5555 sslmode=disable" up

migrate-prod:
	goose -dir=migrations postgres "user=developer password=developer+1DB dbname=laporgubreport port=5555 sslmode=disable" up

sonar-start:
	/Users/default/Apps/sonarqube-8.2.0.32929/bin/macosx-universal-64/sonar.sh start

sonar-log:
	/Users/default/Apps/sonarqube-8.2.0.32929/bin/macosx-universal-64/sonar.sh start

sonar:
	  /Users/default/Apps/sonar-scanner-4.2.0.1873-macosx/bin/sonar-scanner   -Dsonar.projectKey=organization \
                                                                                  -Dsonar.sources=. \
                                                                                  -Dsonar.host.url=http://localhost:9000 \
                                                                                  -Dsonar.login=e8ceb53356694dcf67b0af096be298497358c05d

.PHONY: clean install unittest build docker run stop vendor