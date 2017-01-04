package yaml

import "github.com/Scardiecat/svermaker"

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
