package whisper

type Transcript struct {
	Segments []*Segment
}

type Segment struct {
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
	Text  string `json:"text"`
}
