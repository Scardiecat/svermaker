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
