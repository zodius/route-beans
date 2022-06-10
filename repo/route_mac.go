//build: +darwin
package repo

import (
	"fmt"
	"log"
	"os/exec"
	"route-beans/model"
	"strings"
)

type routeRepo struct{}

func NewRouteRepo() model.RouteRepo {
	return &routeRepo{}
}

func (r *routeRepo) FlushRoutingTable() (err error) {
	// Run several times until all cleanned.
	try := 10
	for i := try; i > 0; i-- {
		cmd := exec.Command("route", "-n", "flush")
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("Failed to run cmd: %s\n", err)
		}
		if len(out) == 0 {
			break
		}
	}
	log.Print("Flush successed.")
	return
}

func (r *routeRepo) AddRouting(dst string, gateway string) (err error) {
	// Run several times until all cleanned.
	cmd := exec.Command("route", "-n", "add", dst, gateway)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Failed to run cmd: %s\n", err)
	}

	if strings.Contains(string(out), "File exists") {
		return nil
	} else if strings.Contains(string(out), "add net") {
		return nil
	}

	return fmt.Errorf(string(out))
}
