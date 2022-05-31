package repo

import (
	"fmt"
	"io/ioutil"
	"route-beans/model"

	yaml "gopkg.in/yaml.v3"
)

type ProfileImplement model.Profile

func (p *ProfileImplement) UnmarshalYAML(data *yaml.Node) error {
	type TempProfile struct {
		Name           string      `yaml:"name"`
		Gateways       []yaml.Node `yaml:"gateways"`
		Routes         []yaml.Node `yaml:"routes"`
		DefaultGateway string      `yaml:"default_gateway"`
	}
	var tempProfile TempProfile
	data.Decode(&tempProfile)

	p.Name = tempProfile.Name

	// Parse Gateway
	for _, gateway_node := range tempProfile.Gateways {
		type TempGateway struct {
			Name string `yaml:"name"`
			IP   string `yaml:"ip"`
		}
		// Try decode well format Gateway
		var gateway TempGateway
		gateway_node.Decode(&gateway)
		if gateway.Name != "" && gateway.IP != "" {
			p.Gateways = append(p.Gateways, model.Gateway(gateway))
			continue
		}

		// Try decode string format Gateway
		var ip string
		gateway_node.Decode(&ip)
		if ip != "" {
			p.Gateways = append(p.Gateways, model.Gateway{
				Name: ip,
				IP:   ip,
			})
			continue
		}

		// Else error
		return fmt.Errorf("gateway format error")
	}

	// Parse Routes
	for _, routes_node := range tempProfile.Routes {
		type TempRoute struct {
			Dst     string `yaml:"dst"`
			Gateway string `yaml:"gateway"`
		}
		var route TempRoute
		routes_node.Decode(&route)
		if route.Dst != "" && route.Gateway != "" {
			if gateway, found := findGateway(route.Gateway, p.Gateways); found {
				p.Routes = append(p.Routes, model.Route{
					Dst:     route.Dst,
					Gateway: gateway,
				})
				continue
			} else {
				return fmt.Errorf("gateway is not defined in definition")
			}
		}
		return fmt.Errorf("routes format error")
	}

	// Parse DefaultGateway
	gateway, found := findGateway(tempProfile.DefaultGateway, p.Gateways)
	if !found {
		return fmt.Errorf("default gateway is not defined in definition")
	}

	p.DefaultGateway = gateway

	return nil
}

type profileRepo struct{}

func NewProfileRepo() model.ProfileRepo {
	return &profileRepo{}
}

func (r *profileRepo) LoadProfileFile(filename string) (profile model.Profile, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	var parsedProfile ProfileImplement

	err = yaml.Unmarshal(data, &parsedProfile)
	return model.Profile(parsedProfile), err
}

func findGateway(gateway_name string, gateways []model.Gateway) (result model.Gateway, found bool) {
	for _, gateway := range gateways {
		if gateway_name == gateway.Name {
			return gateway, true
		}
	}
	return model.Gateway{}, false
}
