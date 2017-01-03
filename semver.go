package svermaker

// Version represents a semver compatible version
type Version struct {
	Major uint64
	Minor uint64
	Patch uint64
	Pre   []PRVersion
	Build []string //No Precendence
}

// PRVersion represents a PreRelease Version
type PRVersion struct {
	VersionStr string
	VersionNum uint64
	IsNum      bool
}

type SemverComponent int8

const (
	PATCH SemverComponent = iota
	MINOR
	MAJOR
)

// Serializer allows reading and writing of a Semver
type Serializer interface {
	Write(semver Version) error
	Read() (Version, error)
}

type SemverService interface {
	Semver(key string) (Version, error)
	CreateSemver(s *Version, key string) error
	DeleteSemver(key string) error
	WriteSemver(s Version, key string) error
}

type Manipulator interface {
	Bump(semver Version, component SemverComponent) (Version, error)
	SetPrerelease(semver Version, prerelease []PRVersion) (Version, error)
	SetMetadata(semver Version, metadata []string) (Version, error)
	MakePrerelease(s ...string) ([]PRVersion, error)
}
