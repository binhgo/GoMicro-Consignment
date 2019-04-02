package main

import pb "github.com/binhgo/GoMicro-Consignment/proto/consignment"
import "gopkg.in/mgo.v2"

const dbName  = "Shippy"
const consignmentCollection = "Consignments"

type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
 }

type ConsignmentRepository struct {
	session *mgo.Session
}

func (repo *ConsignmentRepository) Create(con *pb.Consignment) error {
	err := repo.Collection().Insert(con)
	return err
}

func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	var cons []*pb.Consignment
	err := repo.Collection().Find(nil).All(&cons)
	return cons, err
}

func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}

func (repo *ConsignmentRepository) Collection() *mgo.Collection  {
	return repo.session.DB(dbName).C(consignmentCollection)
}