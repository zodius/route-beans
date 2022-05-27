package service

import (
	"fmt"
	"os"
	"route-beans/model"
)

type routeService struct {
	routeTableRepo model.RouteTableRepo
}

func isRoot() bool {
	return os.Geteuid() == 0
}

func NewRouteService(routeTableRepo model.RouteTableRepo) model.RouteService {
	return &routeService{
		routeTableRepo: routeTableRepo,
	}
}

func (s *routeService) FlushRoutingTable() (err error) {
	if !isRoot() {
		return fmt.Errorf("operation required root permission")
	}
	return s.routeTableRepo.FlushRoutingTable()
}

func (s *routeService) LoadFromProfile(profilePath string) (err error) {
	panic("Not Implemented Error")
}
