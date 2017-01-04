package yaml_test

import (
	"github.com/Scardiecat/svermaker"
	mock "github.com/Scardiecat/svermaker/mock"

	yaml "github.com/Scardiecat/svermaker/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Projectversionservice", func() {
	var serializer = mock.Serializer{}
	var pvs = yaml.ProjectVersionService{Serializer: &serializer}

	BeforeEach(func() {
		serializer = mock.Serializer{}
		pvs = yaml.ProjectVersionService{Serializer: &serializer}
	})
	Describe("Init()", func() {
		Context("If a ProjectVersion does not exist", func() {
			It("should create a new saved version", func() {
				serializer.ExistsFn = func() bool {
					return false
				}
				serializer.SerializerFn = func(p svermaker.ProjectVersion) error {
					return nil
				}

				p, err := pvs.Init()

				Expect(err).To(BeNil())
				Expect(serializer.ExistsInvoked).To(BeTrue())
				Expect(serializer.SerializerInvoked).To(BeTrue())
				Expect(serializer.DeserializerInvoked).To(BeFalse())
				Expect(p).To(Equal(&svermaker.DefaultProjectVersion))
			})
			It("should return the default version", func() {
				serializer.ExistsFn = func() bool {
					return false
				}
				serializer.SerializerFn = func(p svermaker.ProjectVersion) error {
					return nil
				}

				p, err := pvs.Init()
				Expect(err).To(BeNil())
				Expect(p).To(Equal(&svermaker.DefaultProjectVersion))
			})
		})
		Context("If a ProjectVersion does exist", func() {
			It("should use the existing version and return it", func() {
				saved := &svermaker.ProjectVersion{Current: svermaker.Version{1, 1, 1, nil, nil}, Next: svermaker.Version{1, 1, 1, nil, nil}}

				serializer.ExistsFn = func() bool {
					return true
				}
				serializer.DeserializeFn = func() (*svermaker.ProjectVersion, error) {
					return saved, nil
				}

				p, err := pvs.Init()

				Expect(err).To(BeNil())
				Expect(serializer.ExistsInvoked).To(BeTrue())
				Expect(serializer.SerializerInvoked).To(BeFalse())
				Expect(serializer.DeserializerInvoked).To(BeTrue())
				Expect(p).To(Equal(saved))
			})
		})
	})
})
