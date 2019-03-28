build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/GoMicro-Consignment/GoMicro-Consignment \
	  proto/consignment/consignment.proto
	  go get golang.org/x/sys/unix
	  GOOS=linux GOARCH=amd64 go build
	  docker login --username huynhbinh -p Cicevn2007
	  docker build -t consignment-service .

run:
	docker run -p 8081:50051 consignment-service
