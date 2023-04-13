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

	Context("no segments", func() {
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

	Context("some segments", func() {
		BeforeEach(func() {
			in = &whisper.Transcript{Segments: []*whisper.Segment{
				{Start: 0, End: 1.5, Text: "Hello"},
				{Start: 1.5, End: 3.5, Text: "Welcome to the show"},
				{Start: 3.5, End: 4.5, Text: "Good bye"},
			}}
		})

		It("succeeds", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("creates the same amount of segments", func() {
			Expect(out.Segments).To(HaveLen(len(in.Segments)))
		})

		Describe("first segment", func() {
			var segment *podlove.Segment

			JustBeforeEach(func() {
				segment = out.Segments[0]
			})

			It("has the expected start timestamp", func() {
				Expect(segment.Start).To(Equal("00:00:00.000"))
			})

			It("has the expected start millicesond timestamp", func() {
				Expect(segment.StartMs).To(BeNumerically("==", 0))
			})

			It("has the expected end timestamp", func() {
				Expect(segment.End).To(Equal("00:00:01.500"))
			})

			It("has the expected end millicesond timestamp", func() {
				Expect(segment.EndMs).To(BeNumerically("==", 1500))
			})

			It("has the expected text", func() {
				Expect(segment.Text).To(Equal("Hello"))
			})
		})

		Describe("middle segment", func() {
			var segment *podlove.Segment

			JustBeforeEach(func() {
				segment = out.Segments[1]
			})

			It("has the expected start timestamp", func() {
				Expect(segment.Start).To(Equal("00:00:01.500"))
			})

			It("has the expected start millicesond timestamp", func() {
				Expect(segment.StartMs).To(BeNumerically("==", 1500))
			})

			It("has the expected end timestamp", func() {
				Expect(segment.End).To(Equal("00:00:03.500"))
			})

			It("has the expected end millicesond timestamp", func() {
				Expect(segment.EndMs).To(BeNumerically("==", 3500))
			})

			It("has the expected text", func() {
				Expect(segment.Text).To(Equal("Welcome to the show"))
			})
		})

		Describe("last segment", func() {
			var segment *podlove.Segment

			JustBeforeEach(func() {
				segment = out.Segments[2]
			})

			It("has the expected start timestamp", func() {
				Expect(segment.Start).To(Equal("00:00:03.500"))
			})

			It("has the expected start millicesond timestamp", func() {
				Expect(segment.StartMs).To(BeNumerically("==", 3500))
			})

			It("has the expected end timestamp", func() {
				Expect(segment.End).To(Equal("00:00:04.500"))
			})

			It("has the expected end millicesond timestamp", func() {
				Expect(segment.EndMs).To(BeNumerically("==", 4500))
			})

			It("has the expected text", func() {
				Expect(segment.Text).To(Equal("Good bye"))
			})
		})
	})
})
