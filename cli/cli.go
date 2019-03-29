package main

import (
	"context"
	"encoding/json"
	pb "github.com/GoMicro-Consignment/GoMicro-Consignment/proto/consignment"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"io/ioutil"
	"log"
	"os"
)

const address  = "localhost:8081"
const defaultFilename = "consignment.json"


func parseFile(file string) (*pb.Consignment, error) {
	var con *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &con)
	return con, err
}

func main() {
	cmd.Init()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("Did not connect: %v", err)
	//}
	//defer conn.Close()
	//client := pb.NewShippingServiceClient(conn)


	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Created: %t", r.Created)


	r, err = client.GetConsignment(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	for _, v := range r.Consignments {
		log.Println(v)
	}
}


