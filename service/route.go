package service

import (
	"fmt"
	"log"
	"os"
	"route-beans/model"
)

type routeService struct {
	routeTableRepo model.RouteTableRepo
	profileRepo    model.ProfileRepo
}

func isRoot() bool {
	return os.Geteuid() == 0
}

func NewRouteService(routeTableRepo model.RouteTableRepo, profileRepo model.ProfileRepo) model.RouteService {
	return &routeService{
		routeTableRepo: routeTableRepo,
		profileRepo:    profileRepo,
	}
}

func (s *routeService) FlushRoutingTable() (err error) {
	if !isRoot() {
		return fmt.Errorf("operation required root permission")
	}
	return s.routeTableRepo.FlushRoutingTable()
}

func (s *routeService) ApplyProfile(profilePath string) (err error) {
	if !isRoot() {
		return fmt.Errorf("operation required root permission")
	}

	profile, err := s.profileRepo.LoadProfileFile(profilePath)
	if err != nil {
		return err
	}

	log.Printf("Load profile %s\n", profile.Name)

	log.Println("Flush table")
	if err := s.routeTableRepo.FlushRoutingTable(); err != nil {
		return err
	}

	for _, rule := range profile.Routes {
		log.Printf("Add route %s -> %s\n", rule.Dst, rule.Gateway.IP)
		if err := s.routeTableRepo.AddRouting(rule.Dst, rule.Gateway.IP); err != nil {
			return err
		}
	}

	log.Printf("Add default gateway -> %s\n", profile.DefaultGateway.IP)
	if err := s.routeTableRepo.AddRouting("0.0.0.0", profile.DefaultGateway.IP); err != nil {
		return err
	}

	return
}
