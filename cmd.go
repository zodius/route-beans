package main

import (
	"fmt"
	"log"
	"os"

	"route-beans/repo"
	"route-beans/service"

	"github.com/urfave/cli/v2"
)

func main() {
	routeRepo := repo.NewRouteRepo()
	profileRepo := repo.NewProfileRepo()
	service := service.NewRouteService(routeRepo, profileRepo)

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
					profile_path := c.String("p")
					if profile_path == "" {
						return fmt.Errorf("profile yaml is required")
					}
					return service.ApplyProfile(profile_path)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "p",
						Aliases: []string{"profile"},
						Usage:   "path of profile yaml",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
