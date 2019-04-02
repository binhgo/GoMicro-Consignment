package main

import (
	"fmt"
	pb "github.com/binhgo/GoMicro-Consignment/proto/consignment"
	vesselPB "github.com/binhgo/GoMicro-Vessel/proto/vessel"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
	"log"
)

const port = ":50051"

func main() {

	repo := &Repository{}

	service := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselPB.NewVesselServiceClient("go.micro.srv.vessel", service.Client())

	service.Init()

	pb.RegisterShippingServiceHandler(service.Server(), &Service{repo:repo, vesselClient:vesselClient})

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
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
	vesselClient vesselPB.VesselServiceClient
}

func (s *Service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselPB.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = consignment
	return nil
}

//func (s *Service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
//
//	consignment, err := s.repo.Create(req)
//	if err != nil {
//		return nil, err
//	}
//
//	return &pb.Response{Created: true, Consignment: consignment}, nil
//}

func (s *Service) GetConsignment(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

//func (s *Service) GetConsignment(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
//	return &pb.Response{Created:true, Consignments:s.repo.GetAll()}, nil
//}