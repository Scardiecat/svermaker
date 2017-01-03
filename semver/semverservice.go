package semver

import "github.com/Scardiecat/svermaker"

// Ensure SemverService implements svermaker.SemverService.
var _ svermaker.SemverService = &SemverService{}

// SemverService handles Semvars
type SemverService struct {
}

func (ss *SemverService) Semver(key string) (svermaker.Version, error) {
	v := svermaker.Version{}
	return v, nil
}
func (ss *SemverService) CreateSemver(s *svermaker.Version, key string) error {
	return nil
}
func (ss *SemverService) DeleteSemver(key string) error {
	return nil
}
func (ss *SemverService) WriteSemver(s svermaker.Version, key string) error {
	return nil
}
