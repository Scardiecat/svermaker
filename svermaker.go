package main

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/blang/semver"
	"github.com/urfave/cli"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the debug severity or above.
	log.SetLevel(log.DebugLevel)
}
func main() {
	app := cli.NewApp()
	app.Name = "svermaker"
	app.Usage = "Help with semver versioning for git projects"
	app.Commands = []cli.Command{
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "just some sample code",
			Action: func(c *cli.Context) error {
				v, _ := semver.Make("0.0.1-alpha.preview+123.github")
				fmt.Printf("Major: %d\n", v.Major)
				fmt.Printf("Minor: %d\n", v.Minor)
				fmt.Printf("Patch: %d\n", v.Patch)
				fmt.Printf("Pre: %s\n", v.Pre)
				fmt.Printf("Build: %s\n", v.Build)

				return nil
			},
		},
		{
			Name:    "current",
			Aliases: []string{"c"},
			Usage:   "show the current semver",
			Action: func(c *cli.Context) error {
				// open a file
				if file, err := os.Open("version.txt"); err == nil {

					// make sure it gets closed
					defer file.Close()

					// create a new scanner and read the file line by line
					scanner := bufio.NewScanner(file)
					for scanner.Scan() {
						log.Println(scanner.Text())
					}

					// check for errors
					if err = scanner.Err(); err != nil {
						log.Fatal(err)
					}

				} else {
					log.Fatal(err)
				}
				return nil
			},
		},
		{
			Name:    "next",
			Aliases: []string{"n"},
			Usage:   "next semvar",
			Subcommands: []cli.Command{
				{
					Name:  "major",
					Usage: "next is a major release (breaking api)",
					Action: func(c *cli.Context) error {
						log.Info("new major release: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "minor",
					Usage: "next is a minor release (non api break adding feature)",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "patch",
					Usage: "next is a patch release (non api break bug fix)",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
