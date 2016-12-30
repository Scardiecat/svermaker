package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/blang/semver"
	"github.com/urfave/cli"
)

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
	}

	app.Run(os.Args)
}
