package yaml

import "github.com/Scardiecat/svermaker"

// Ensure ProjectVersionService implements svermaker.ProjectVersionService.
var _ svermaker.ProjectVersionService = &ProjectVersionService{}

// ProjectVersionService
type ProjectVersionService struct {
	client *Client
}

func (p *ProjectVersionService) Init() (*svermaker.ProjectVersion, error) {
	return &svermaker.ProjectVersion{}, nil
}
