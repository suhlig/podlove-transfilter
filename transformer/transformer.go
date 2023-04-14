package transformer

import (
	"fmt"
	"podlove-transfilter/podlove"
	"podlove-transfilter/whisper"
	"strings"
	"time"
)

func Transform(whisperTranscript *whisper.Transcript) (podloveTranscript *podlove.Transcript, err error) {
	podloveTranscript = &podlove.Transcript{}

	for _, whisperSegment := range whisperTranscript.Segments {
		podloveSegment, err := transformSegment(whisperSegment)

		if err != nil {
			return nil, err
		}

		podloveTranscript.Segments = append(podloveTranscript.Segments, &podloveSegment)
	}

	return
}

func transformSegment(wrs *whisper.Segment) (pls podlove.Segment, err error) {
	pls.Start = fmtDuration(time.Duration(wrs.Start * float64(time.Second)))
	pls.StartMs = uint64(wrs.Start * 1000)
	pls.End = fmtDuration(time.Duration(wrs.End * float64(time.Second)))
	pls.EndMs = uint64(wrs.End * 1000)
	pls.Text = strings.Trim(wrs.Text, " ")

	return pls, nil
}

func fmtDuration(d time.Duration) string {
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	d -= s * time.Second
	ms := d / time.Millisecond

	return fmt.Sprintf("%02d:%02d:%02d.%03d", h, m, s, ms)
}
