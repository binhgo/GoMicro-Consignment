build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/binhgo/GoMicro-Consignment \
	proto/consignment/consignment.proto
	docker login --username huynhbinh -p Cicevn2007
	SSH_PRIVATE_KEY='$(cat ~/.ssh/id_rsa)'
	docker build . --build-arg SSH_PRIVATE_KEY -t consignment-service

run:
	docker run --link=GoMicroMongoDB:mongodb --name GoMicroConsignmentService -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		-e DB_HOST=mongodb:27017 consignment-service