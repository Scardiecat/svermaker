package semver

import (
	"reflect"
	"testing"

	"github.com/Scardiecat/svermaker"
)

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

func TestBump(t *testing.T) {
	m := Manipulator{}
	for _, test := range bumpTests {
		if res, err := m.Bump(test.v, test.c); !reflect.DeepEqual(res, test.result) || err != nil {
			t.Errorf("result %s does not equal expected %s", exportTo(res).String(), exportTo(test.result).String())
		}
	}
}
