//build: +darwin
package repo

import (
	"fmt"
	"log"
	"os/exec"
	"route-beans/model"
)

type routeTableRepo struct{}

func NewRouteRepo() model.RouteTableRepo {
	return &routeTableRepo{}
}

func (r *routeTableRepo) FlushRoutingTable() (err error) {
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
