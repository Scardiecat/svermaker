package main

import (
	"os"

	"github.com/Scardiecat/svermaker"
	"github.com/Scardiecat/svermaker/semver"
	"github.com/Scardiecat/svermaker/yaml"
	log "github.com/Sirupsen/logrus"
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
			Name:    "current",
			Aliases: []string{"c"},
			Usage:   "show the current semver",
			Action: func(c *cli.Context) error {
				var serializer = yaml.NewSerializer(c.Args().First())
				var pvs = semver.ProjectVersionService{Serializer: serializer}
				if v, err := pvs.GetCurrent(); err == nil {
					log.Infof("Version is %s", v.String())
					return nil
				} else {
					return err
				}
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
						var serializer = yaml.NewSerializer("")
						var m = semver.Manipulator{}
						var pre []svermaker.PRVersion
						if c.Args().Present() {
							first, _ := m.MakePrerelease(c.Args().First())
							pre = append(pre, first...)

							tail, _ := m.MakePrerelease(c.Args().Tail()...)

							pre = append(pre, tail...)
						}

						var pvs = semver.ProjectVersionService{Serializer: serializer}
						if v, err := pvs.Bump(svermaker.MAJOR, pre); err == nil {
							log.Infof("Current version is %s", v.Current.String())
							log.Infof("Next version is %s", v.Next.String())
							return nil
						} else {
							return err
						}
						return nil
					},
				},
				{
					Name:  "minor",
					Usage: "next is a minor release (non api break adding feature)",
					Action: func(c *cli.Context) error {
						var serializer = yaml.NewSerializer("")
						var m = semver.Manipulator{}
						var pre []svermaker.PRVersion
						if c.Args().Present() {
							first, _ := m.MakePrerelease(c.Args().First())
							pre = append(pre, first...)

							tail, _ := m.MakePrerelease(c.Args().Tail()...)

							pre = append(pre, tail...)
						}

						var pvs = semver.ProjectVersionService{Serializer: serializer}
						if v, err := pvs.Bump(svermaker.MINOR, pre); err == nil {
							log.Infof("Current version is %s", v.Current.String())
							log.Infof("Next version is %s", v.Next.String())
							return nil
						} else {
							return err
						}
						return nil
					},
				},
				{
					Name:  "patch",
					Usage: "next is a patch release (non api break bug fix)",
					Action: func(c *cli.Context) error {
						var serializer = yaml.NewSerializer("")
						var m = semver.Manipulator{}
						var pre []svermaker.PRVersion
						if c.Args().Present() {
							first, _ := m.MakePrerelease(c.Args().First())
							pre = append(pre, first...)

							tail, _ := m.MakePrerelease(c.Args().Tail()...)

							pre = append(pre, tail...)
						}

						var pvs = semver.ProjectVersionService{Serializer: serializer}
						if v, err := pvs.Bump(svermaker.PATCH, pre); err == nil {
							log.Infof("Current version is %s", v.Current.String())
							log.Infof("Next version is %s", v.Next.String())
							return nil
						} else {
							return err
						}
						return nil
					},
				},
			},
		},
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init a version",
			Action: func(c *cli.Context) error {
				var serializer = yaml.NewSerializer(c.Args().First())
				var pvs = semver.ProjectVersionService{Serializer: serializer}
				if v, err := pvs.Init(); err == nil {
					log.Infof("Version is %s", v.Current.String())
					return nil
				} else {
					return err
				}
			},
		},
	}

	app.Run(os.Args)
}
