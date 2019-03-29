build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/GoMicro-Consignment/GoMicro-Consignment \
	proto/consignment/consignment.proto
	go get golang.org/x/sys/unix
	GOOS=linux GOARCH=amd64 go build
	docker login --username huynhbinh -p Cicevn2007
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns consignment-service