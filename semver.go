package svermaker

import "strconv"

// Version represents a semver compatible version
type Version struct {
	Major uint64
	Minor uint64
	Patch uint64
	Pre   []PRVersion
	Build []string //No Precedence
}

// Version to string
func (v Version) String() string {
	b := make([]byte, 0, 5)
	b = strconv.AppendUint(b, v.Major, 10)
	b = append(b, '.')
	b = strconv.AppendUint(b, v.Minor, 10)
	b = append(b, '.')
	b = strconv.AppendUint(b, v.Patch, 10)

	if len(v.Pre) > 0 {
		b = append(b, '-')
		b = append(b, v.Pre[0].String()...)

		for _, pre := range v.Pre[1:] {
			b = append(b, '.')
			b = append(b, pre.String()...)
		}
	}

	if len(v.Build) > 0 {
		b = append(b, '+')
		b = append(b, v.Build[0]...)

		for _, build := range v.Build[1:] {
			b = append(b, '.')
			b = append(b, build...)
		}
	}

	return string(b)
}

// PreRelease version to string
func (v PRVersion) String() string {
	if v.IsNum {
		return strconv.FormatUint(v.VersionNum, 10)
	}
	return v.VersionStr
}

// PRVersion represents a PreRelease Version
type PRVersion struct {
	VersionStr string
	VersionNum uint64
	IsNum      bool
}

// SemverComponent is the base of the enum
type SemverComponent int8

//Semvar component enum
const (
	PATCH SemverComponent = iota
	MINOR
	MAJOR
)

// ProjectVersion holds the current and next version identifier
type ProjectVersion struct {
	Current Version
	Next    Version
}

// DefaultProjectVersion is the init value for a new repo
var DefaultProjectVersion = ProjectVersion{Current: Version{0, 1, 0, nil, nil}, Next: Version{0, 1, 0, nil, nil}}

// Serializer is an interface to write and read ProjectVersions
type Serializer interface {
	Serialize(p ProjectVersion) error
	Deserialize() (*ProjectVersion, error)
	Exists() bool
}

//ProjectVersionService is the api for the service
type ProjectVersionService interface {
	Init() (*ProjectVersion, error)
	GetCurrent() (*Version, error)
	Get() (*ProjectVersion, error)
	Bump(component SemverComponent, prerelease []PRVersion) (*ProjectVersion, error)
	Release() (*ProjectVersion, error)
}

//Manipulator is the api to change versions
type Manipulator interface {
	Bump(semver Version, component SemverComponent) (Version, error)
	SetPrerelease(semver Version, prerelease []PRVersion) (Version, error)
	SetMetadata(semver Version, metadata []string) (Version, error)
	MakePrerelease(s ...string) ([]PRVersion, error)
	Create(s string) (*Version, error)
	Compare(v1 Version, v2 Version) int
}

//BuildHelper is the api to write a helper file for build tools
type BuildHelper interface {
	MakeTags(p ProjectVersion, buildMetadata []string) error
}
