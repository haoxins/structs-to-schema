package structstoschema

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStructsToSchema(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Structs to Schema Suite")
}
