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

// Bump will increase a version
func (m *Manipulator) Bump(semver svermaker.Version, component svermaker.SemverComponent) (svermaker.Version, error) {
	switch component {
	case svermaker.PATCH:
		semver.Patch++
	case svermaker.MINOR:
		semver.Minor++
		semver.Patch = 0
	case svermaker.MAJOR:
		semver.Major++
		semver.Minor = 0
		semver.Patch = 0
	}
	return semver, nil
}

// SetPrerelease sets a prerelease version
func (m *Manipulator) SetPrerelease(semver svermaker.Version, prerelease []svermaker.PRVersion) (svermaker.Version, error) {
	semver.Pre = prerelease
	return semver, nil
}

// SetMetadata sets build metadata
func (m *Manipulator) SetMetadata(semver svermaker.Version, metadata []string) (svermaker.Version, error) {
	semver.Build = metadata
	return semver, nil
}

// MakePrerelease makes a prerelease
func (m *Manipulator) MakePrerelease(s ...string) ([]svermaker.PRVersion, error) {
	bpres := make([]blangs.PRVersion, 0)
	for _, p := range s {
		v, err := blangs.NewPRVersion(p)
		if err == nil {
			bpres = append(bpres, v)
		}
	}
	return setPreFrom(bpres), nil
}

// Create a version
func (m *Manipulator) Create(s string) (*svermaker.Version, error) {
	bv, err := blangs.Make(s)
	if err != nil {
		return nil, err
	}
	return setFrom(bv), nil
}

func setPreFrom(bv []blangs.PRVersion) []svermaker.PRVersion {
	pre := make([]svermaker.PRVersion, 0)
	for _, bpre := range bv {
		pre = append(pre, svermaker.PRVersion(bpre))
	}
	return pre
}

func exportPreTo(v []svermaker.PRVersion) []blangs.PRVersion {
	bpre := make([]blangs.PRVersion, 0)
	for _, pre := range v {
		bpre = append(bpre, blangs.PRVersion(pre))
	}
	return bpre
}
func setFrom(bv blangs.Version) *svermaker.Version {
	pre := setPreFrom(bv.Pre)

	build := make([]string, 0)
	for _, bbuild := range bv.Build {
		build = append(build, bbuild)
	}
	return &svermaker.Version{bv.Major, bv.Minor, bv.Patch, pre, build}
}

func exportTo(v svermaker.Version) *blangs.Version {
	bpre := exportPreTo(v.Pre)

	bbuild := make([]string, 0)
	for _, build := range v.Build {
		bbuild = append(bbuild, build)
	}
	return &blangs.Version{v.Major, v.Minor, v.Patch, bpre, bbuild}
}

// Compare compares 2 versions
func (m *Manipulator) Compare(v1 svermaker.Version, v2 svermaker.Version) int {
	b1 := exportTo(v1)
	b2 := exportTo(v2)

	return b1.Compare(*b2)
}
