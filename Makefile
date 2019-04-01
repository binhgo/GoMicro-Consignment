build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/GoMicro-Consignment/GoMicro-Consignment \
	proto/consignment/consignment.proto
	docker login --username huynhbinh -p Cicevn2007
	docker build --build-arg https://binhgo:32771fab1118c28299299f9f765fb3db0d0a691f@github.com -t consignment-service .

run:
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns consignment-service