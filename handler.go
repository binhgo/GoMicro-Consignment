package main

import (
	vesselPB "github.com/binhgo/GoMicro-Vessel/proto/vessel"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"log"
	pb "github.com/binhgo/GoMicro-Consignment/proto/consignment"
)

type Handler struct {
	session *mgo.Session
	vesselClient vesselPB.VesselServiceClient
}

func (han *Handler) GetRepo() Repository {
	return &ConsignmentRepository{han.session.Clone()}
}

func (han *Handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	repo := han.GetRepo()
	defer repo.Close()

	vesselResponse, err := han.vesselClient.FindAvailable(context.Background(), &vesselPB.Specification{
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
	err = repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = req
	return nil
}

func (han *Handler) GetConsignment(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := han.GetRepo()
	defer repo.Close()

	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}

	res.Consignments = consignments
	return nil
}


