package main

import (
	"log"
	"os"

	"route-beans/repo"
	"route-beans/service"

	"github.com/urfave/cli/v2"
)

func main() {
	routeRepo := repo.NewRouteRepo()
	service := service.NewRouteService(routeRepo)

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "reset",
				Usage: "Reset/Flush routing table to default",
				Action: func(c *cli.Context) error {
					return service.FlushRoutingTable()
				},
			},
			{
				Name:  "profile",
				Usage: "Load routing table setting from profile",
				Action: func(c *cli.Context) error {
					return service.LoadFromProfile("")
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
