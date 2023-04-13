package whisper

type Transcript struct {
	Segments []*Segment
}

type Segment struct {
	Start float64 `json:"start"`
	End   float64 `json:"end"`
	Text  string  `json:"text"`
}
