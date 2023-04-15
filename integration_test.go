package main_test

import (
	"io"
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("podlove-transfilter", func() {
	var (
		err     error
		session *gexec.Session
		args    []string
		stdin   io.Reader = nil
		stdout            = GinkgoWriter
		stderr            = GinkgoWriter
	)

	BeforeEach(func() {
		args = []string{}
	})

	JustBeforeEach(func() {
		cmd := exec.Command(binary, args...)
		cmd.Stdin = stdin
		session, err = gexec.Start(cmd, stdout, stderr)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("asking for help", func() {
		BeforeEach(func() {
			args = []string{"--help"}
		})

		It("exits with code zero", func() {
			Eventually(session).Should(gexec.Exit(0))
		})

		It("prints something helpful", func() {
			Eventually(session.Err).Should(gbytes.Say("prints it in Podlove format"))
		})
	})

	Context("asking for the version", func() {
		BeforeEach(func() {
			args = []string{"--version"}
		})

		It("exits with code zero", func() {
			Eventually(session).Should(gexec.Exit(0))
		})

		It("prints the dev version", func() {
			Eventually(session.Out).Should(gbytes.Say("vDEV"))
		})

		It("prints the unknown SHA", func() {
			Eventually(session.Out).Should(gbytes.Say("(NONE)"))
		})

		It("prints the unknown build date", func() {
			Eventually(session.Out).Should(gbytes.Say("UNKNOWN"))
		})
	})

	Context("transforming invalid JSON", func() {
		BeforeEach(func() {
			stdin = strings.NewReader(`DEADBEEF`)
		})

		It("exits non-zero", func() {
			Eventually(session).Should(gexec.Exit(1))
		})
	})

	Context("transforming minimal input", func() {
		BeforeEach(func() {
			stdin = strings.NewReader(`{"segments": [{}]}`)
		})

		It("exits with code zero", func() {
			Eventually(session).Should(gexec.Exit(0))
		})

		It("contains the start_ms field", func() {
			Eventually(session.Out).Should(gbytes.Say("start_ms"))
		})
	})
})
