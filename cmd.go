package main

import (
	"fmt"
	"log"
	"os"
	"route-beans/repo"
	"route-beans/service"

	"github.com/urfave/cli/v2"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print("[+] " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	routeRepo := repo.NewRouteRepo()
	profileRepo := repo.NewProfileRepo()
	routeBeanService := service.NewRouteBeanService(routeRepo, profileRepo)

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List profiles",
				Action: func(c *cli.Context) error {
					return routeBeanService.ListProfiles()
				},
			},
			{
				Name:  "apply",
				Usage: "Apply profile",
				Action: func(c *cli.Context) error {
					profileName := c.Args().First()
					return routeBeanService.ApplyProfile(profileName)
				},
			},
			{
				Name:  "reset",
				Usage: "Reset routing table to default",
				Action: func(c *cli.Context) error {
					return routeBeanService.Reset()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
