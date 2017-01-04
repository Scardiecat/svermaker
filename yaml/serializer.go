package yaml

import (
	"os"

	"io/ioutil"

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
	current string
	next    string
}

func NewSerializer(path string) svermaker.Serializer {
	if path == "" {
		path = "./"
	}
	s := &Serializer{Path: path, Filename: "version.yml"}
	s.projectVersionService.Serializer = s
	return s
}

func (s *Serializer) Exists() bool {
	return false
}

func (s *Serializer) Serialize(p svermaker.ProjectVersion) error {
	v := projectVersion{current: p.Current.String(), next: p.Next.String()}
	b, err := yaml.Marshal(&v)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.Path+s.Filename, b, 0755)
	if err != nil {
		return err
	}
	return nil
}

func (s *Serializer) Deserialize() (*svermaker.ProjectVersion, error) {
	v := projectVersion{}
	m := semver.Manipulator{}
	projectVersion := svermaker.ProjectVersion{}
	if file, err := os.Open(s.Path + s.Filename); err == nil {

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
	if current, err := m.Create(v.current); err == nil {
		projectVersion.Current = *current
	} else {
		return nil, err
	}

	if next, err := m.Create(v.next); err == nil {
		projectVersion.Next = *next
	} else {
		return nil, err
	}

	return &projectVersion, nil
}
