package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("test-results/junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Launch client", []Reporter{junitReporter})
}
