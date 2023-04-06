package podlove

type Transcript struct {
	Segments []*Segment
}

type Segment struct {
	Start   string `json:"start"`
	StartMs uint64 `json:"start_ms"`
	End     string `json:"end"`
	EndMs   uint64 `json:"end_ms"`
	Speaker string `json:"speaker"`
	Voice   string `json:"voice"`
	Text    string `json:"text"`
}
