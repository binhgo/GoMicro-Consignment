package main

import (
	"fmt"
	pb "github.com/binhgo/GoMicro-Consignment/proto/consignment"
	vesselPB "github.com/binhgo/GoMicro-Vessel/proto/vessel"
	micro "github.com/micro/go-micro"
	"log"
	"os"
)

const port = ":50051"
const defaultHost  = "localhost:27017"

func main() {

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	service := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselPB.NewVesselServiceClient("go.micro.srv.vessel", service.Client())

	service.Init()

	pb.RegisterShippingServiceHandler(service.Server(), &Handler{session:session, vesselClient:vesselClient})

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}