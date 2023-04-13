package main

import (
	"encoding/json"
	"fmt"
	"main/transformer"
	"main/whisper"
	"os"
)

func main() {
	var transcript whisper.Transcript

	err := json.NewDecoder(os.Stdin).Decode(&transcript)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not parse input: %s", err)
		os.Exit(1)
	}

	podloveTranscript, err := transformer.Transform(&transcript)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:  Could not transform transcript from whisper to podlove: %s", err)
		os.Exit(1)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(podloveTranscript.Segments)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning:  Could not encode podlove transcript to JSON: %s", err)
		os.Exit(1)
	}
}
