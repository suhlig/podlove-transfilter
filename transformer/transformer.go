package transformer

import (
	"fmt"
	"podlove-transfilter/podlove"
	"podlove-transfilter/whisper"
	"strings"
	"time"
)

type Options struct {
	Offset time.Duration
}

func Transform(whisperTranscript *whisper.Transcript, options Options) (podloveTranscript *podlove.Transcript, err error) {
	podloveTranscript = &podlove.Transcript{}

	for _, whisperSegment := range whisperTranscript.Segments {
		if time.Duration(whisperSegment.Start) < options.Offset {
			continue
		}

		podloveSegment, err := transformSegment(whisperSegment, options)

		if err != nil {
			return nil, err
		}

		podloveTranscript.Segments = append(podloveTranscript.Segments, &podloveSegment)
	}

	return
}

func transformSegment(wrs *whisper.Segment, options Options) (pls podlove.Segment, err error) {
	startWithOffset := time.Duration(wrs.Start) - options.Offset
	pls.Start = fmtDuration(startWithOffset)
	pls.StartMs = uint64(startWithOffset / time.Millisecond)

	endWithOffset := time.Duration(wrs.End) - options.Offset
	pls.End = fmtDuration(endWithOffset)
	pls.EndMs = uint64(endWithOffset / time.Millisecond)

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
