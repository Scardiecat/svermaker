package mock

import "github.com/Scardiecat/svermaker"

type Serializer struct {
	SerializerFn      func(p svermaker.ProjectVersion) error
	SerializerInvoked bool

	DeserializeFn       func() (*svermaker.ProjectVersion, error)
	DeserializerInvoked bool

	ExistsFn      func() bool
	ExistsInvoked bool
}

func (s *Serializer) Serialize(p svermaker.ProjectVersion) error {
	s.SerializerInvoked = true
	return s.SerializerFn(p)
}

func (s *Serializer) Deserialize() (*svermaker.ProjectVersion, error) {
	s.DeserializerInvoked = true
	return s.DeserializeFn()
}

func (s *Serializer) Exists() bool {
	s.ExistsInvoked = true
	return s.ExistsFn()
}
