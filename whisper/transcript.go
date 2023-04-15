package whisper

import (
	"strconv"
	"time"
)

type Transcript struct {
	Segments []*Segment
}

type Timestamp time.Duration

type Segment struct {
	Start Timestamp `json:"start"`
	End   Timestamp `json:"end"`
	Text  string    `json:"text"`
}

func NewTimestampFromSeconds(seconds float64) Timestamp {
	return Timestamp(fromSeconds(seconds))
}

func (t *Timestamp) UnmarshalJSON(s []byte) error {
	seconds, err := strconv.ParseFloat(string(s), 64)

	if err != nil {
		return err
	}

	*(*time.Duration)(t) = fromSeconds((seconds))
	return nil
}

func fromSeconds(seconds float64) time.Duration {
	return time.Duration(seconds * float64(time.Second))
}
