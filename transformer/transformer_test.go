package transformer_test

import (
	"main/podlove"
	"main/transformer"
	"main/whisper"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transformer", func() {
	var out *podlove.Transcript
	var in *whisper.Transcript
	var err error

	JustBeforeEach(func() {
		out, err = transformer.Transform(in)
	})

	Context("No segments", func() {
		BeforeEach(func() {
			in = &whisper.Transcript{}
		})

		It("succeeds", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("creates the same amount of segments", func() {
			Expect(out.Segments).To(HaveLen(len(in.Segments)))
		})
	})
})
