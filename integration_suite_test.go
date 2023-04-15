package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestPodloveTransfilter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PodloveTransfilter Integration Test Suite")
}

var binary string

var _ = BeforeSuite(func() {
	var compileError error
	binary, compileError = gexec.Build("main.go")
	Expect(compileError).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
