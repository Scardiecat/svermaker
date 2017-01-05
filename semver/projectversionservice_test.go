package semver_test

import (
	"github.com/Scardiecat/svermaker"
	mock "github.com/Scardiecat/svermaker/mock"

	"github.com/Scardiecat/svermaker/semver"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Projectversionservice", func() {
	var serializer = mock.Serializer{}
	var pvs = semver.ProjectVersionService{Serializer: &serializer}

	BeforeEach(func() {
		serializer = mock.Serializer{}
		pvs = semver.ProjectVersionService{Serializer: &serializer}
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
	Describe("Get()", func() {
		Context("If a ProjectVersion does not exist", func() {
			It("it should raise an error", func() {
				serializer.ExistsFn = func() bool {
					return false
				}
				_, err := pvs.Get()

				Expect(err).To(MatchError("version not found"))
			})
		})
		Context("If a ProjectVersion does  exist", func() {
			It("it should return it", func() {
				current := svermaker.Version{1, 1, 1, nil, nil}
				saved := &svermaker.ProjectVersion{Current: current, Next: current}

				serializer.ExistsFn = func() bool {
					return true
				}
				serializer.DeserializeFn = func() (*svermaker.ProjectVersion, error) {
					return saved, nil
				}
				c, err := pvs.Get()

				Expect(err).To(BeNil())
				Expect(c).To(Equal(saved))
			})
		})
	})
	Describe("GetCurrent()", func() {
		Context("If a ProjectVersion does not exist", func() {
			It("it should raise an error", func() {
				serializer.ExistsFn = func() bool {
					return false
				}
				_, err := pvs.GetCurrent()

				Expect(err).To(MatchError("version not found"))
			})
		})
		Context("If a ProjectVersion does  exist", func() {
			It("it should return the current version", func() {
				current := svermaker.Version{1, 1, 1, nil, nil}
				saved := &svermaker.ProjectVersion{Current: current, Next: svermaker.Version{1, 1, 2, nil, nil}}

				serializer.ExistsFn = func() bool {
					return true
				}
				serializer.DeserializeFn = func() (*svermaker.ProjectVersion, error) {
					return saved, nil
				}
				c, err := pvs.GetCurrent()

				Expect(err).To(BeNil())
				Expect(c).To(Equal(&current))
			})
		})
	})
	Describe("Bump()", func() {
		It("should fail if the repo is not initialized", func() {
			serializer.ExistsFn = func() bool {
				return false
			}

			_, err := pvs.Bump(svermaker.PATCH, nil)
			Expect(err).To(MatchError("version not found"))
		})
		Context("When bumping for a patch ", func() {
			It("should increase the patch version on the next version", func() {

				current := svermaker.Version{1, 1, 1, nil, nil}
				next := svermaker.Version{1, 1, 2, nil, nil}
				saved := &svermaker.ProjectVersion{Current: current, Next: current}
				serializer.ExistsFn = func() bool {
					return true
				}
				serializer.DeserializeFn = func() (*svermaker.ProjectVersion, error) {
					return saved, nil
				}
				serializer.SerializerFn = func(p svermaker.ProjectVersion) error {
					return nil
				}

				p, err := pvs.Bump(svermaker.PATCH, nil)

				Expect(err).To(BeNil())
				Expect(serializer.ExistsInvoked).To(BeTrue())
				Expect(serializer.SerializerInvoked).To(BeTrue())
				Expect(serializer.DeserializerInvoked).To(BeTrue())
				Expect(p.Next).To(Equal(next))
			})
			Context("When a prerelease version is set", func() {
				It("should use the supplied prerelease version for current version", func() {
					current := svermaker.Version{1, 1, 1, nil, nil}
					next := svermaker.Version{1, 1, 2, nil, nil}
					expected := &svermaker.ProjectVersion{Current: current, Next: next}
					saved := &svermaker.ProjectVersion{Current: current, Next: current}
					serializer.ExistsFn = func() bool {
						return true
					}
					serializer.DeserializeFn = func() (*svermaker.ProjectVersion, error) {
						return saved, nil
					}
					serializer.SerializerFn = func(p svermaker.ProjectVersion) error {
						return nil
					}
					m := semver.Manipulator{}
					pre, _ := m.MakePrerelease("testpre")
					p, err := pvs.Bump(svermaker.PATCH, pre)

					Expect(err).To(BeNil())
					Expect(serializer.ExistsInvoked).To(BeTrue())
					Expect(serializer.SerializerInvoked).To(BeTrue())
					Expect(serializer.DeserializerInvoked).To(BeTrue())
					expected.Current, _ = m.SetPrerelease(expected.Next, pre)
					Expect(p).To(Equal(expected))
				})
			})
			Context("When a prerelease version is not set", func() {
				It("should set the prerelease to rc for current version", func() {
					current := svermaker.Version{1, 1, 1, nil, nil}
					next := svermaker.Version{1, 1, 2, nil, nil}
					expected := &svermaker.ProjectVersion{Current: current, Next: next}
					saved := &svermaker.ProjectVersion{Current: current, Next: current}
					serializer.ExistsFn = func() bool {
						return true
					}
					serializer.DeserializeFn = func() (*svermaker.ProjectVersion, error) {
						return saved, nil
					}
					serializer.SerializerFn = func(p svermaker.ProjectVersion) error {
						return nil
					}

					p, err := pvs.Bump(svermaker.PATCH, nil)
					m := semver.Manipulator{}
					pre, _ := m.MakePrerelease("rc")
					Expect(err).To(BeNil())
					Expect(serializer.ExistsInvoked).To(BeTrue())
					Expect(serializer.SerializerInvoked).To(BeTrue())
					Expect(serializer.DeserializerInvoked).To(BeTrue())
					expected.Current, _ = m.SetPrerelease(expected.Next, pre)
					Expect(p).To(Equal(expected))
				})
			})
		})
		Context("When bumping for a minor", func() {
			It("should increase the minor version on the next version", func() {
				current := svermaker.Version{1, 1, 1, nil, nil}
				next := svermaker.Version{1, 2, 0, nil, nil}
				saved := &svermaker.ProjectVersion{Current: current, Next: current}
				serializer.ExistsFn = func() bool {
					return true
				}
				serializer.DeserializeFn = func() (*svermaker.ProjectVersion, error) {
					return saved, nil
				}
				serializer.SerializerFn = func(p svermaker.ProjectVersion) error {
					return nil
				}

				p, err := pvs.Bump(svermaker.MINOR, nil)

				Expect(err).To(BeNil())
				Expect(serializer.ExistsInvoked).To(BeTrue())
				Expect(serializer.SerializerInvoked).To(BeTrue())
				Expect(serializer.DeserializerInvoked).To(BeTrue())
				Expect(p.Next).To(Equal(next))
			})
			Context("When a prerelease version is set", func() {
				It("should use the supplied prerelease version for current version", func() {
					current := svermaker.Version{1, 1, 1, nil, nil}
					next := svermaker.Version{1, 2, 0, nil, nil}
					expected := &svermaker.ProjectVersion{Current: current, Next: next}
					saved := &svermaker.ProjectVersion{Current: current, Next: current}
					serializer.ExistsFn = func() bool {
						return true
					}
					serializer.DeserializeFn = func() (*svermaker.ProjectVersion, error) {
						return saved, nil
					}
					serializer.SerializerFn = func(p svermaker.ProjectVersion) error {
						return nil
					}
					m := semver.Manipulator{}
					pre, _ := m.MakePrerelease("testpre")
					p, err := pvs.Bump(svermaker.MINOR, pre)

					Expect(err).To(BeNil())
					Expect(serializer.ExistsInvoked).To(BeTrue())
					Expect(serializer.SerializerInvoked).To(BeTrue())
					Expect(serializer.DeserializerInvoked).To(BeTrue())
					expected.Current, _ = m.SetPrerelease(expected.Next, pre)
					Expect(p).To(Equal(expected))
				})
			})
			Context("When a prerelease version is not set", func() {
				It("should set the prerelease to beta for current version", func() {
					current := svermaker.Version{1, 1, 1, nil, nil}
					next := svermaker.Version{1, 2, 0, nil, nil}
					expected := &svermaker.ProjectVersion{Current: current, Next: next}
					saved := &svermaker.ProjectVersion{Current: current, Next: current}
					serializer.ExistsFn = func() bool {
						return true
					}
					serializer.DeserializeFn = func() (*svermaker.ProjectVersion, error) {
						return saved, nil
					}
					serializer.SerializerFn = func(p svermaker.ProjectVersion) error {
						return nil
					}

					p, err := pvs.Bump(svermaker.MINOR, nil)
					m := semver.Manipulator{}
					pre, _ := m.MakePrerelease("beta")
					Expect(err).To(BeNil())
					Expect(serializer.ExistsInvoked).To(BeTrue())
					Expect(serializer.SerializerInvoked).To(BeTrue())
					Expect(serializer.DeserializerInvoked).To(BeTrue())
					expected.Current, _ = m.SetPrerelease(expected.Next, pre)
					Expect(p).To(Equal(expected))
				})
			})
		})
		Context("When bumping for a major", func() {
			It("should increase the major version on the next version", func() {
				current := svermaker.Version{1, 1, 1, nil, nil}
				next := svermaker.Version{2, 0, 0, nil, nil}
				saved := &svermaker.ProjectVersion{Current: current, Next: current}
				serializer.ExistsFn = func() bool {
					return true
				}
				serializer.DeserializeFn = func() (*svermaker.ProjectVersion, error) {
					return saved, nil
				}
				serializer.SerializerFn = func(p svermaker.ProjectVersion) error {
					return nil
				}

				p, err := pvs.Bump(svermaker.MAJOR, nil)

				Expect(err).To(BeNil())
				Expect(serializer.ExistsInvoked).To(BeTrue())
				Expect(serializer.SerializerInvoked).To(BeTrue())
				Expect(serializer.DeserializerInvoked).To(BeTrue())
				Expect(p.Next).To(Equal(next))
			})
			Context("When a prerelease version is set", func() {
				It("should use the supplied prerelease version for current version", func() {
					current := svermaker.Version{1, 1, 1, nil, nil}
					next := svermaker.Version{2, 0, 0, nil, nil}
					expected := &svermaker.ProjectVersion{Current: current, Next: next}
					saved := &svermaker.ProjectVersion{Current: current, Next: current}
					serializer.ExistsFn = func() bool {
						return true
					}
					serializer.DeserializeFn = func() (*svermaker.ProjectVersion, error) {
						return saved, nil
					}
					serializer.SerializerFn = func(p svermaker.ProjectVersion) error {
						return nil
					}
					m := semver.Manipulator{}
					pre, _ := m.MakePrerelease("testpre")
					p, err := pvs.Bump(svermaker.MAJOR, pre)

					Expect(err).To(BeNil())
					Expect(serializer.ExistsInvoked).To(BeTrue())
					Expect(serializer.SerializerInvoked).To(BeTrue())
					Expect(serializer.DeserializerInvoked).To(BeTrue())
					expected.Current, _ = m.SetPrerelease(expected.Next, pre)
					Expect(p).To(Equal(expected))
				})
			})
			Context("When a prerelease version is not set", func() {
				It("should set the prerelease to alpha for current version", func() {
					current := svermaker.Version{1, 1, 1, nil, nil}
					next := svermaker.Version{2, 0, 0, nil, nil}
					expected := &svermaker.ProjectVersion{Current: current, Next: next}
					saved := &svermaker.ProjectVersion{Current: current, Next: current}
					serializer.ExistsFn = func() bool {
						return true
					}
					serializer.DeserializeFn = func() (*svermaker.ProjectVersion, error) {
						return saved, nil
					}
					serializer.SerializerFn = func(p svermaker.ProjectVersion) error {
						return nil
					}

					p, err := pvs.Bump(svermaker.MAJOR, nil)
					m := semver.Manipulator{}
					pre, _ := m.MakePrerelease("alpha")
					Expect(err).To(BeNil())
					Expect(serializer.ExistsInvoked).To(BeTrue())
					Expect(serializer.SerializerInvoked).To(BeTrue())
					Expect(serializer.DeserializerInvoked).To(BeTrue())
					expected.Current, _ = m.SetPrerelease(expected.Next, pre)
					Expect(p).To(Equal(expected))
				})
			})
		})
	})
	Describe("Release()", func() {
		It("should set the current version to the next version", func() {
			current := svermaker.Version{1, 1, 1, nil, nil}
			next := svermaker.Version{2, 0, 0, nil, nil}
			expected := &svermaker.ProjectVersion{Current: next, Next: next}
			saved := &svermaker.ProjectVersion{Current: current, Next: next}
			serializer.ExistsFn = func() bool {
				return true
			}
			serializer.DeserializeFn = func() (*svermaker.ProjectVersion, error) {
				return saved, nil
			}
			serializer.SerializerFn = func(p svermaker.ProjectVersion) error {
				return nil
			}

			p, err := pvs.Release()
			Expect(err).To(BeNil())
			Expect(serializer.ExistsInvoked).To(BeTrue())
			Expect(serializer.SerializerInvoked).To(BeTrue())
			Expect(serializer.DeserializerInvoked).To(BeTrue())
			Expect(p).To(Equal(expected))
		})
	})
})
