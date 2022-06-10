package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"route-beans/model"
)

func root_required() error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("operation required root privileges")
	}
	return nil
}

type routeBeanService struct {
	routeRepo   model.RouteRepo
	profileRepo model.ProfileRepo
}

func NewRouteBeanService(routeRepo model.RouteRepo, profileRepo model.ProfileRepo) model.RouteBeanService {
	return &routeBeanService{
		routeRepo,
		profileRepo,
	}
}

func (s *routeBeanService) ListProfiles() (err error) {
	ymls, err := filepath.Glob("profiles/*.yml")
	if err != nil {
		return err
	}

	for _, yml := range ymls {
		// Ignore format error when listing
		profile, err := s.profileRepo.LoadProfileFromFile(yml)
		if err == nil {
			fmt.Println(profile.Name)
		}
	}

	return
}

func (s *routeBeanService) findProfile(profileName string) (profile model.Profile, err error) {
	ymls, err := filepath.Glob("profiles/*.yml")
	if err != nil {
		return
	}

	for _, yml := range ymls {
		// Ignore format error when listing
		profile, err := s.profileRepo.LoadProfileFromFile(yml)
		if err == nil && profile.Name == profileName {
			return profile, nil
		}
	}

	return model.Profile{}, fmt.Errorf("profile not found")
}

func (s *routeBeanService) ApplyProfile(profileName string) (err error) {
	if err = root_required(); err != nil {
		return err
	}

	// Flush before adding
	if err = s.routeRepo.FlushRoutingTable(); err != nil {
		return err
	}

	profile, err := s.findProfile(profileName)
	if err != nil {
		return err
	}

	for _, route := range profile.Routes {
		log.Printf("Add %s -> %s\n", route.Dst, route.Gateway.IP)
		err = s.routeRepo.AddRouting(route.Dst, route.Gateway.IP)
		if err != nil {
			return err
		}
	}
	log.Printf("Add Default gateway %s\n", profile.DefaultGateway.IP)
	err = s.routeRepo.AddRouting("0.0.0.0", profile.DefaultGateway.IP)
	return
}

func (s *routeBeanService) Reset() (err error) {
	if err = root_required(); err != nil {
		return err
	}

	return s.routeRepo.FlushRoutingTable()
}
