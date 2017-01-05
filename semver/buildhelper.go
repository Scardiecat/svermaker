package semver

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Scardiecat/svermaker"

	log "github.com/Sirupsen/logrus"
)

// Ensure BuildHelper implements svermaker.BuildHelper.
var _ svermaker.BuildHelper = &BuildHelper{}

// BuildHelper
type BuildHelper struct {
}

func (b *BuildHelper) MakeTags(p svermaker.ProjectVersion, buildMetadata []string) error {
	m := Manipulator{}

	isRelease := m.Compare(p.Current, p.Next) == 0
	c := p.Current
	if !isRelease {
		md, err := m.SetMetadata(c, buildMetadata)
		if err != nil {
			return err
		}
		c = md
	}

	bc := os.Getenv("buildcounter")
	if !isRelease && bc != "" {
		i, err := strconv.ParseUint(bc, 10, 64)
		if err != nil {
			log.Errorf("buildcounter is not not a number its: %s", bc)
		} else {
			pre := c.Pre
			if len(pre) >= 1 && pre[len(pre)-1].IsNum {
				pre = pre[:len(pre)-1]
			}
			pre = append(pre, svermaker.PRVersion{"", i, true})
			c.Pre = pre
		}
	} else {
		if bc == "" {
			log.Warningln("Add a buildcounter environment variable to have it added to the prereleaseversion")
		}
	}
	f, err := os.Create("./buildhelper.tmp")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("export svermakerBuildVersion=%s\nexport svermakerRelease=%t", c.String(), isRelease))
	if err != nil {
		return err
	}
	f.Sync()
	return nil
}