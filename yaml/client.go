package yaml

import "github.com/Scardiecat/svermaker"

// Client represents a client to the yaml serializer
type Client struct {
	// Filename to ProjectConfig.
	Path string

	// Services
	projectVersionService ProjectVersionService
}

func NewYamlClient(path string) svermaker.Client {
	if path == "" {
		path = "./"
	}
	c := &Client{Path: path}
	c.projectVersionService.client = c
	return c
}
func (c *Client) ProjectVersionService() svermaker.ProjectVersionService {
	return &c.projectVersionService
}
