package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"podlove-transfilter/transformer"
	"podlove-transfilter/whisper"
	"time"

	flag "github.com/spf13/pflag"
)

// ldflags will be set by goreleaser
var version = "vDEV"
var commit = "NONE"
var date = "UNKNOWN"

func main() {
	helpWanted := flag.BoolP("help", "h", false, "provides help")
	versionWanted := flag.BoolP("version", "V", false, "prints the version and exits")
	offset := flag.DurationP("offset", "o", time.Duration(0), "offset all podlove timestamps against the whisper timestamps by this amount")

	flag.Parse()

	if helpWanted != nil && *helpWanted {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]...\n", os.Args[0])
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Reads a Whispher transcript from STDIN and prints it in Podlove format to STDOUT.")
		fmt.Fprintln(os.Stderr)
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Example:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintf(os.Stderr, "  %s < whisper.json | tee podlove-transcript.json\n", os.Args[0])
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "will read whisper.json and print the podlove transcript as well as write it to podlove-transcript.json.")
		os.Exit(0)
	}

	if versionWanted != nil && *versionWanted {
		fmt.Printf("%s %s (%s), built on %s\n", filepath.Base(os.Args[0]), version, commit, date)
		os.Exit(0)
	}

	var transcript whisper.Transcript

	err := json.NewDecoder(os.Stdin).Decode(&transcript)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not parse input: %s", err)
		os.Exit(1)
	}

	podloveTranscript, err := transformer.Transform(&transcript, transformer.Options{
		Offset: *offset,
	})

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
