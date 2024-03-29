package semver

import (
	"reflect"

	"github.com/Scardiecat/svermaker"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("bump", func() {

		m := Manipulator{}
		for _, test := range bumpTests {
			if res, err := m.Bump(test.v, test.c); !reflect.DeepEqual(res, test.result) || err != nil {
				GinkgoT().Errorf("result %s does not equal expected %s", exportTo(res).String(), exportTo(test.result).String())
			}
		}
	})
	It("make prerelease", func() {

		m := Manipulator{}
		for _, test := range makePreReleaseTests {
			if res, err := m.MakePrerelease(test.v...); !reflect.DeepEqual(res, test.result) || err != nil {
				GinkgoT().Errorf("result %#v does not equal expected %#v", res, test.result)
			}
		}
	})
	It("set pre and metadata", func() {

		m := Manipulator{}
		for _, test := range setPreAndMetadataTests {
			pre, _ := m.MakePrerelease(test.p...)
			res1, err := m.SetMetadata(test.v, test.m)
			if err != nil {
				GinkgoT().Errorf("error occurred %#v", err)
			}
			if res, err := m.SetPrerelease(res1, pre); !reflect.DeepEqual(res, test.result) || err != nil {
				GinkgoT().Errorf("result %#v does not equal expected %#v", res, test.result)
			}
		}
	})
})

func prstr(s string) svermaker.PRVersion {
	return svermaker.PRVersion{s, 0, false}
}

func prnum(i uint64) svermaker.PRVersion {
	return svermaker.PRVersion{"", i, true}
}

type bumpTest struct {
	v      svermaker.Version
	c      svermaker.SemverComponent
	result svermaker.Version
}

var bumpTests = []bumpTest{
	{svermaker.Version{0, 0, 0, nil, nil}, svermaker.PATCH, svermaker.Version{0, 0, 1, nil, nil}},
	{svermaker.Version{0, 0, 0, nil, nil}, svermaker.MINOR, svermaker.Version{0, 1, 0, nil, nil}},
	{svermaker.Version{0, 0, 0, nil, nil}, svermaker.MAJOR, svermaker.Version{1, 0, 0, nil, nil}},
}

type makePreReleaseTest struct {
	v      []string
	result []svermaker.PRVersion
}

var makePreReleaseTests = []makePreReleaseTest{
	{[]string{"alpha"}, []svermaker.PRVersion{prstr("alpha")}},
	{[]string{"alpha", "beta"}, []svermaker.PRVersion{prstr("alpha"), prstr("beta")}},
	{[]string{"1", "beta"}, []svermaker.PRVersion{prnum(1), prstr("beta")}},
}

type setPreAndMetadataTest struct {
	v      svermaker.Version
	p      []string
	m      []string
	result svermaker.Version
}

var setPreAndMetadataTests = []setPreAndMetadataTest{
	{svermaker.Version{0, 0, 0, nil, nil}, []string{"alpha"}, []string{"build123"}, svermaker.Version{0, 0, 0, []svermaker.PRVersion{prstr("alpha")}, []string{"build123"}}},
}
