package yaml

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/Scardiecat/svermaker"
	"github.com/Scardiecat/svermaker/semver"
	"gopkg.in/yaml.v2"
)

type Serializer struct {
	// path to version.txt
	Path     string
	Filename string

	// Services
	projectVersionService ProjectVersionService
}

type projectVersion struct {
	Current string
	Next    string
}

func NewSerializer(path string) *Serializer {
	if path == "" {
		path = "."
	}
	s := &Serializer{Path: path, Filename: "/version.yml"}
	s.projectVersionService.Serializer = s
	return s
}

func (s *Serializer) Exists() bool {
	if _, err := os.Stat(path.Join(s.Path, s.Filename)); os.IsNotExist(err) {
		return false
	}
	return true
}

func (s *Serializer) Serialize(p svermaker.ProjectVersion) error {
	v := projectVersion{Current: p.Current.String(), Next: p.Next.String()}
	b, err := yaml.Marshal(&v)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path.Join(s.Path, s.Filename), b, 0755)
	if err != nil {
		return err
	}
	return nil
}

func (s *Serializer) Deserialize() (*svermaker.ProjectVersion, error) {
	v := projectVersion{}
	m := semver.Manipulator{}
	projectVersion := svermaker.ProjectVersion{}
	if file, err := os.Open(path.Join(s.Path, s.Filename)); err == nil {

		// make sure it gets closed
		defer file.Close()
		if d, err := ioutil.ReadAll(file); err == nil {

			if err := yaml.Unmarshal([]byte(d), &v); err == nil {

			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
	if current, err := m.Create(v.Current); err == nil {
		projectVersion.Current = *current
	} else {
		return nil, err
	}

	if next, err := m.Create(v.Next); err == nil {
		projectVersion.Next = *next
	} else {
		return nil, err
	}

	return &projectVersion, nil
}
