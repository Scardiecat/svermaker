package mock

import "github.com/Scardiecat/svermaker"

// Serializer is a mock of the Serializer interface
type Serializer struct {
	SerializerFn      func(p svermaker.ProjectVersion) error
	SerializerInvoked bool

	DeserializeFn       func() (*svermaker.ProjectVersion, error)
	DeserializerInvoked bool

	ExistsFn      func() bool
	ExistsInvoked bool
}

// Serialize is a mock
func (s *Serializer) Serialize(p svermaker.ProjectVersion) error {
	s.SerializerInvoked = true
	return s.SerializerFn(p)
}

// Deserialize is a mock
func (s *Serializer) Deserialize() (*svermaker.ProjectVersion, error) {
	s.DeserializerInvoked = true
	return s.DeserializeFn()
}

// Exists is a mock
func (s *Serializer) Exists() bool {
	s.ExistsInvoked = true
	return s.ExistsFn()
}
