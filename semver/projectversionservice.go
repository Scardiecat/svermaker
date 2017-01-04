package semver

import "github.com/Scardiecat/svermaker"
import "errors"

// Ensure ProjectVersionService implements svermaker.ProjectVersionService.
var _ svermaker.ProjectVersionService = &ProjectVersionService{}

// ProjectVersionService
type ProjectVersionService struct {
	Serializer svermaker.Serializer
}

func (p *ProjectVersionService) Init() (*svermaker.ProjectVersion, error) {
	if p.Serializer.Exists() {
		v, err := p.Serializer.Deserialize()
		if err == nil {
			return v, nil
		} else {
			return nil, err
		}
	}
	v := svermaker.DefaultProjectVersion
	p.Serializer.Serialize(v)
	return &v, nil
}

func (p *ProjectVersionService) GetCurrent() (*svermaker.Version, error) {
	if p.Serializer.Exists() {
		v, err := p.Serializer.Deserialize()
		if err != nil {
			return nil, err
		}
		return &v.Current, nil
	}
	return nil, errors.New("version not found")
}

func (p *ProjectVersionService) Bump(component svermaker.SemverComponent, prerelease []svermaker.PRVersion) (*svermaker.ProjectVersion, error) {
	m := Manipulator{}
	if !p.Serializer.Exists() {
		return nil, errors.New("version not found")
	}
	v, err := p.Serializer.Deserialize()
	if err != nil {
		return nil, err
	}
	b, err := m.Bump(v.Next, component)
	if err != nil {
		return nil, err
	}
	v.Next = b
	v.Current = v.Next
	if prerelease == nil {
		switch component {
		case svermaker.PATCH:
			prerelease, _ = m.MakePrerelease("rc")
		case svermaker.MINOR:
			prerelease, _ = m.MakePrerelease("beta")
		case svermaker.MAJOR:
			prerelease, _ = m.MakePrerelease("alpha")
		}
	}
	v.Current.Pre = prerelease
	// write it out
	err = p.Serializer.Serialize(*v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
