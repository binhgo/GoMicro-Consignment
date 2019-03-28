package main

import (
	pb "github.com/GoMicro-Consignment/GoMicro-Consignment/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)
import "golang.org/x/net/context"

const port = ":50051"

func main() {

	repo := &Repository{}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterShippingServiceServer(server, &Service{repo:repo})
	reflection.Register(server)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}


type Repository struct {
	consignments []*pb.Consignment
}


func(repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated  := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll()  []*pb.Consignment {
	return repo.consignments
}

type Service struct {
	repo IRepository
}

func (s *Service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *Service) GetConsignment(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	return &pb.Response{Created:true, Consignments:s.repo.GetAll()}, nil
}