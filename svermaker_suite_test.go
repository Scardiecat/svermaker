package svermaker_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSvermaker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Svermaker Suite")
}
