build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/binhgo/GoMicro-Consignment \
	proto/consignment/consignment.proto
	docker login --username huynhbinh -p Cicevn2007
	export SSH_PRIVATE_KEY='$(cat ~/.ssh/id_rsa)'
	docker build . --build-arg SSH_PRIVATE_KEY -t consignment-service

run:
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns consignment-service