package main

import (
	"context"

	pb "github.com/jmwinn21/shippy/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRespository{s.session.Clone()}
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	defer s.GetRepo().Close()
	vessel, err := s.GetRepo().FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	defer s.GetRepo().Close()
	if err := s.GetRepo().Create(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}
