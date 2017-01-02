package semver

import (
	"github.com/Scardiecat/svermaker"
	blangs "github.com/blang/semver"
)

// Ensure Manipulator implements svermaker.Manipulator.
var _ svermaker.Manipulator = &Manipulator{}

// Manipulator changes semvar values
type Manipulator struct {
}

func (m *Manipulator) Bump(semver svermaker.Version, component svermaker.SemverComponent) (svermaker.Version, error) {
	switch component {
	case svermaker.PATCH:
		semver.Patch += 1
	case svermaker.MINOR:
		semver.Minor += 1
	case svermaker.MAJOR:
		semver.Major += 1
	}
	return semver, nil
}
func (m *Manipulator) SetPrerelease(semver svermaker.Version, annotation string) (svermaker.Version, error) {
	return semver, nil
}
func (m *Manipulator) SetMetadata(semver svermaker.Version, metadata string) (svermaker.Version, error) {
	return semver, nil
}

func setFrom(bv blangs.Version) svermaker.Version {
	pre := make([]svermaker.PRVersion, 1)
	for _, bpre := range bv.Pre {
		pre = append(pre, svermaker.PRVersion(bpre))
	}

	build := make([]string, 1)
	for _, bbuild := range bv.Build {
		build = append(build, bbuild)
	}
	return svermaker.Version{bv.Major, bv.Minor, bv.Patch, pre, build}
}

func exportTo(v svermaker.Version) blangs.Version {
	bpre := make([]blangs.PRVersion, 0)
	for _, pre := range v.Pre {
		bpre = append(bpre, blangs.PRVersion(pre))
	}

	bbuild := make([]string, 0)
	for _, build := range v.Build {
		bbuild = append(bbuild, build)
	}
	return blangs.Version{v.Major, v.Minor, v.Patch, bpre, bbuild}
}
